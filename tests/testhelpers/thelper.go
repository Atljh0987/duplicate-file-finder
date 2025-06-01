package testhelpers

import (
	"os"
	fp "path/filepath"
)

func GetResourcePath() string {
	dir, err := os.Getwd()

	if err != nil {
		panic(err.Error())
	}

	return fp.Join(dir, "/resources")
}
