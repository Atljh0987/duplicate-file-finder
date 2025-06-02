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

	for _, fileData := range files {
		file, err := os.Open(fileData.Path)
		defer file.Close()

		if err != nil {
			notifier.Notifier.NotifyError(notifier.ConsoleNotifier{}, err)
			continue
		}

		if _, err := io.Copy(hasher, file); err != nil {
			notifier.Notifier.NotifyError(notifier.ConsoleNotifier{}, err)
			continue
		}
		fileData.Hash.HashType = storage.XXHash
		fileData.Hash.Payload = hasher.Sum64()
	}

	return files
}
