// Package fileman contains file explorer-like functions for
// Directories, Files, and Symbolic Links (Copy, Paste, Delete, etc..).
package fileman

import (
	"errors"
	"os"
	"path/filepath"
)

// GetType returns "dir" or "file" from the path given.
// If set, "symlink" can also be returned.
func GetType(path string, includeSymLinks bool) (string, error) {
	// obtain info from path
	item, err := os.Stat(path)

	// if path does not exist
	if os.IsNotExist(err) {
		return "", errors.New("Path does not exist: " + path)
	}

	// if symbolic links are being checked
	if includeSymLinks {
		// if path is symbolic link
		if itemSL, err := os.Lstat(path); err == nil && itemSL.Mode()&os.ModeSymlink == os.ModeSymlink {
			return "symlink", nil
		}
	}

	// if path is a dir
	if item.Mode().IsDir() {
		return "dir", nil
	}

	// if path is a file
	if item.Mode().IsRegular() {
		return "file", nil
	}

	// else error occurred
	return "", errors.New("Error analysing: " + path)
}

// Rename a specified item.
// This calls os.Rename(), but prevents moving.
func Rename(path, newName string) error {
	// extract directory from item path
	dirPath := filepath.Dir(path) + "/"

	// return any errors from renaming file
	return os.Rename(path, dirPath+newName)
}

// Move an item to a specified path.
// This calls os.Rename(), but prevents renaming.
func Move(path, dirPath string) error {
	// throw error if dirPath doesn't exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return errors.New("Path destination does not exist: " + dirPath)
	}

	// if dirPath does not end with "/"
	if dirPath[len(dirPath)-1:] != "/" {
		// add "/"
		dirPath = dirPath + "/"
	}

	// extract item name from path
	itemName := filepath.Base(path) + "/"

	// return any errors from moving file
	return os.Rename(path, dirPath+itemName)
}

// Delete will remove a specified item.
func Delete(path string) error {
	// if path type is not a dir, remove normally
	if pType, _ := GetType(path, true); pType != "dir" {
		return os.Remove(path)
	}

	// else delete dir
	return os.RemoveAll(path)
}
