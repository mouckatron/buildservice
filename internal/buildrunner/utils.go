package buildrunner

import (
	"path/filepath"
	"strings"
)

func getNameFromURL(url string) (name string) {

	parts := strings.Split(url, "/")

	part := len(parts) - 1

	name = strings.TrimSuffix(parts[part], filepath.Ext(parts[part]))
	return
}
