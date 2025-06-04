package analyzer

import (
	"duplicate-file-finder/dupefinder/notifier"
	"duplicate-file-finder/dupefinder/storage"
	"io"
	"os"
	"github.com/cespare/xxhash"
)

type DefaultEncryptor struct{}

func (e DefaultEncryptor) Encrypt(hash storage.HashType, files []storage.FileData) []storage.FileData {
	return xxhashEncode(files)
}

func xxhashEncode(files []storage.FileData) []storage.FileData {
	hasher := xxhash.New()

	for i := range files {
		fileData := &files[i]
		
		file, err := os.Open(fileData.Path)
		if err != nil {
			notifier.Notifier.NotifyError(notifier.ConsoleNotifier{}, err)
			continue
		}

		defer file.Close()

		if _, err := io.Copy(hasher, file); err != nil {
			notifier.Notifier.NotifyError(notifier.ConsoleNotifier{}, err)
			continue
		}

		fileData.Hash.HashType = storage.XXHash
		fileData.Hash.HashData = hasher.Sum64()

		hasher.Reset()
	}

	return files
}
