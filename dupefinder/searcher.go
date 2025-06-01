package dupefinder

type ScanResult struct {
    Files   []FileData
    Folders []FolderData
}

type Data struct {
	Path string
}

type FolderData struct {
	Data
}

type FileData struct {
	Data
	Size int64
}
