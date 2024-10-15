package system

import "os"

// Exists checks whether the given path exists in the file system.
// Takes a string `path` as input and returns a boolean value.
// Returns true if the path exists, false otherwise.
func Exists(path string) bool {
	_, e := os.Stat(path)
	if e == nil {
		return true
	}

	if os.IsNotExist(e) {
		return false
	}

	return false
}
