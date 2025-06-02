package searcher

import "duplicate-file-finder/dupefinder/storage"

type Searcher interface {
	CollectAll(root string) storage.ScanResult
	CollectAllDeep(root string) []storage.FileData
}
