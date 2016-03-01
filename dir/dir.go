// This package defines some directory operation functions.

package dir

import (
	"os"
	"path/filepath"
)

// ExecutePath returns the absolute path for the current command if success, otherwise returns empty string.
func ExecutePath() string {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		return ""
	}
	return filepath.Dir(pathAbs)
}
