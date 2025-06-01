package dupefinder_test

import (
	sh "duplicate-file-finder/dupefinder/searcher"
	h "duplicate-file-finder/tests/testhelpers"
	fp "path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var resourcePath = h.GetResourcePath()

func TestIsFolder(t *testing.T) {
	var searcher sh.Searcher = sh.Default{}
	var path = fp.Join(resourcePath, "TestIsFolder")

	result := searcher.CollectAll(path)

	assert.Equal(t, 0, len(result.Files))
	assert.Equal(t, 1, len(result.Folders))
	assert.Equal(t, fp.Join(path, "TestFolder"), result.Folders[0].Path)
}

func TestIsFile(t *testing.T) {
	var searcher sh.Searcher = sh.Default{}
	var path = fp.Join(resourcePath, "TestIsFile")

	result := searcher.CollectAll(path)

	assert.Equal(t, 1, len(result.Files))
	assert.Equal(t, 0, len(result.Folders))
	assert.Equal(t, int64(0), result.Files[0].Size)
	assert.Equal(t, fp.Join(path, "1.txt"), result.Files[0].Path)
	assert.Equal(t, "txt", result.Files[0].Extension)
}

func TestFileSize(t *testing.T) {
	var searcher sh.Searcher = sh.Default{}
	var path = fp.Join(resourcePath, "TestFileSize")

	result := searcher.CollectAll(path)

	assert.Equal(t, 1, len(result.Files))
	assert.Equal(t, 0, len(result.Folders))
	assert.Equal(t, int64(23), result.Files[0].Size)
	assert.Equal(t, fp.Join(path, "2.txt"), result.Files[0].Path)
	assert.Equal(t, "txt", result.Files[0].Extension)
}

func TestDeepFiles(t *testing.T) {
	var searcher sh.Searcher = sh.Default{}
	var path = fp.Join(resourcePath, "TestDeepFiles")

	result := searcher.CollectAllDeep(path)

	assert.Equal(t, 5, len(result))
}
