// Package fileman contains file explorer-like functions for
// Directories, Files, and Symbolic Links.
package fileman

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetType returns "dir" or "file" from the path given.
// Including symlinks will allow for "symlink" as a return value.
// If no file is found, an error will be returned
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
func Rename(path, newName string) error {
	// extract full paths
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// extract directory from path
	dirPath := filepath.Dir(path)

	// create path for newName
	newName = filepath.Join(dirPath, newName)

	// throw error if new name is taken
	if _, err := GetType(newName, false); err == nil {
		return errors.New("Already Exists: " + newName)
	}

	// return any errors from renaming file
	return os.Rename(path, newName)
}

// Move an item to a specified directory.
func Move(path, dirPath string) error {
	// extract full paths
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	dirPath, err = filepath.Abs(dirPath)
	if err != nil {
		return err
	}

	// extract item name from path
	itemName := filepath.Base(path)

	// return any errors from moving file
	return os.Rename(path, filepath.Join(dirPath, itemName))
}

// Delete will remove a specified item.
func Delete(path string) error {
	// extract full paths
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// if path type is not a dir, remove normally
	if pathType, _ := GetType(path, true); pathType != "dir" {
		return os.Remove(path)
	}

	// else delete dir
	return os.RemoveAll(path)
}

// Duplicate will clone a specified item and place it in the newPath given.
// newPath should include the new name of what's being duplicated.
// Sync parameter does not affect symlinks.
func Duplicate(path, newPath string, sync bool) error {
	// extract full directory path
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// call correct function for cloning item
	switch itemType, err := GetType(path, true); itemType {
	// File clone
	case "file":
		return cloneFile(path, newPath, sync)
	// Directory clone
	case "dir":
		return cloneDir(path, newPath, sync)
	// Symbolic Link clone
	case "symlink":
		return cloneSymLink(path, newPath)
	// error getting type
	default:
		return err
	}
}

// Search will look inside searchDir for the item you want to find.
// SearchDepth determines how far the search will look inside each directory.
func Search(itemName string, searchDir string, searchDepth int) (itemFound bool, path string) {
	// Stop if search depth is passed
	if searchDepth < 0 {
		return
	}

	// Read Current Directory Items
	searchDir, _ = filepath.Abs(path)
	dirs, _ := ioutil.ReadDir(searchDir)

	// For each item in Directory
	for _, item := range dirs {
		// Update current directory
		newSearchDir := filepath.Join(searchDir, item.Name())

		// If desired item is found, return
		if item.Name() == itemName {
			itemFound = true
			path = newSearchDir
			return
		}
		// Run again since item wasn't found
		itemFound, path = Search(itemName, newSearchDir, searchDepth-1)

		// if path was already found, exit loop
		if path != "" {
			return
		}
	}
	return
}
