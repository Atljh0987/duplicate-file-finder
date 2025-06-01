package searcher

import (
	df "duplicate-file-finder/dupefinder"
	nf "duplicate-file-finder/dupefinder/notifier"
	"os"
	fp "path/filepath"
)

var notifier nf.Notifier = nf.ConsoleNotifier{}

type Searcher interface {
	CollectAll(root string) df.ScanResult
}

type Default struct {
}

func (d Default) CollectAll(root string) df.ScanResult {
	entries, err := os.ReadDir(root)

	var scanResult = df.ScanResult{}

	if err != nil {
		notifier.NotifyError(err)
		return scanResult
	}

	for _, entity := range entries {
		if entity.IsDir() {
			scanResult.Folders = append(scanResult.Folders, DirToFolder(root, entity))
		} else {
			scanResult.Files = append(scanResult.Files, DirToFile(root, entity))
		}
	}

	return scanResult
}

func DirToFile(path string, direntry os.DirEntry) df.FileData {
	var info, err = direntry.Info()

	if err != nil {
		notifier.NotifyError(err)
		return df.FileData{}
	}

	return df.FileData{
		Data: df.Data{Path: fp.Join(path, info.Name())},
		Size: info.Size(),
	}
}

func DirToFolder(path string, direntry os.DirEntry) df.FolderData {
	var info, err = direntry.Info()

	if err != nil {
		notifier.NotifyError(err)
		return df.FolderData{}
	}

	return df.FolderData{
		Data: df.Data{Path: fp.Join(path, info.Name())},
	}
}
