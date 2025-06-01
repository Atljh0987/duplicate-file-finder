package searcher

import (
	df "duplicate-file-finder/dupefinder"
	nf "duplicate-file-finder/dupefinder/notifier"
	"os"
	fp "path/filepath"
	"strings"

	"k8s.io/client-go/util/workqueue"
)

var notifier nf.Notifier = nf.ConsoleNotifier{}

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
			scanResult.Folders = append(scanResult.Folders, dirToFolder(root, entity))
		} else {
			scanResult.Files = append(scanResult.Files, dirToFile(root, entity))
		}
	}

	return scanResult
}

func (d Default) CollectAllDeep(root string) []df.FileData {
	foldersQueue := workqueue.DefaultQueue[df.FolderData]()
	files := []df.FileData{}

	result := d.CollectAll(root)

	for _, folder := range result.Folders {
		foldersQueue.Push(folder)
	}

	files = append(files, result.Files...)

	for foldersQueue.Len() > 0 {
		result := d.CollectAll(foldersQueue.Pop().Path)

		for _, folder := range result.Folders {
			foldersQueue.Push(folder)
		}

		files = append(files, result.Files...)
	}

	return files
}

func dirToFile(path string, direntry os.DirEntry) df.FileData {
	var info, err = direntry.Info()

	if err != nil {
		notifier.NotifyError(err)
		return df.FileData{}
	}

	return df.FileData{
		Data:      df.Data{Path: fp.Join(path, info.Name())},
		Size:      info.Size(),
		Extension: getExtension(info.Name()),
	}
}

func dirToFolder(path string, direntry os.DirEntry) df.FolderData {
	var info, err = direntry.Info()

	if err != nil {
		notifier.NotifyError(err)
		return df.FolderData{}
	}

	return df.FolderData{
		Data: df.Data{Path: fp.Join(path, info.Name())},
	}
}

func getExtension(filename string) string {
	return strings.TrimPrefix(fp.Ext(filename), ".")
}
