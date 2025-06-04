package analyzer

import (
	"duplicate-file-finder/dupefinder/storage"
)

type DefaultFilter struct{}

func (d DefaultFilter) SameSize(files []storage.FileData) []storage.FileData {
	sortedMap := map[int64][]storage.FileData{}
	resStorage := []storage.FileData{}

	for _, file := range files {
		sortedMap[file.Size] = append(sortedMap[file.Size], file)
	}

	for key, value := range sortedMap {
		if len(value) == 1 {
			delete(sortedMap, key)
		} else {
			resStorage = append(resStorage, value...)
		}
	}

	return resStorage
}

func (d DefaultFilter) SameExtension(files []storage.FileData) []storage.FileData {
	sortedMap := map[string][]storage.FileData{}
	resStorage := []storage.FileData{}

	for _, file := range files {
		sortedMap[file.Extension] = append(sortedMap[file.Extension], file)
	}

	for key, value := range sortedMap {
		if len(value) == 1 {
			delete(sortedMap, key)
		} else {
			resStorage = append(resStorage, value...)
		}
	}

	return resStorage
}

func (d DefaultFilter) SameHash(files []storage.FileData) []storage.FileData {
	encFiles := Encryptor.Encrypt(DefaultEncryptor{}, storage.XXHash, files)

	sortedMap := map[uint64][]storage.FileData{}
	resStorage := []storage.FileData{}

	for _, file := range encFiles {
		sortedMap[file.HashData] = append(sortedMap[file.HashData], file)
	}

	for key, value := range sortedMap {
		if len(value) == 1 {
			delete(sortedMap, key)
		} else {
			resStorage = append(resStorage, value...)
		}
	}

	return resStorage
}
