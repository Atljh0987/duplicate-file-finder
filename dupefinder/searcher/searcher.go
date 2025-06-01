package searcher

import df "duplicate-file-finder/dupefinder"

type Searcher interface {
	CollectAll(root string) df.ScanResult
	CollectAllDeep(root string) []df.FileData
}
