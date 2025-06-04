package storage

type HashType int

const (
	XXHash HashType = iota
)

type ScanResult struct {
	Files   []FileData
	Folders []FolderData
}

type Hash struct {
	HashType HashType
	HashData uint64
}

type Data struct {
	Path string
}

type FolderData struct {
	Data
}

type FileData struct {
	Data
	Hash
	Size      int64
	Extension string
}
