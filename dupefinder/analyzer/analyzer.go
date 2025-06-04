package analyzer

import "duplicate-file-finder/dupefinder/storage"

type Filter interface {
	SameSize(files []storage.FileData) []storage.FileData
	SameExtension(files []storage.FileData) []storage.FileData
	SameHash(files []storage.FileData) []storage.FileData
}

type Encryptor interface {
	Encrypt(hash storage.HashType, files []storage.FileData) []storage.FileData
}
