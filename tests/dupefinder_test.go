package dupefinder_test

import (
	sh "duplicate-file-finder/dupefinder/searcher"
	h "duplicate-file-finder/tests/testhelpers"
	fp "path/filepath"
	"testing"
)

var resourcePath = h.GetResourcePath()

func TestIsFolder(t *testing.T) {
	var searcher sh.Searcher = sh.Default{}
	var path = fp.Join(resourcePath, "TestIsFolder")

	result := searcher.CollectAll(path)

	if len(result.Folders) != 1 {
		t.Errorf("Have been found %d folders instead of one", len(result.Folders))
	}

	if result.Folders[0].Path != fp.Join(path, "TestFolder") {
		t.Errorf("Path must be %s, but was %s", fp.Join(path, "TestFolder"), result.Folders[0].Path)
	}
}
