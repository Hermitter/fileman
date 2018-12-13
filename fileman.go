package fileman

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

//////////////////////////////////////////////////////////////////
// Structs used to identify each item you'd see in a file explorer

// File is a structure representing a single file.
type File struct {
	Name     string
	Contents *[]byte
}

// ToString returns the string value of a File's contents (presumably text).
func (f File) ToString() string {
	return fmt.Sprintf("%s", *f.Contents)
}

// SymLink is a structure representing a single symbolic link
type SymLink struct {
	Name string
	Link string
}

// Dir is a structure representing a single Directory.
// 3 slices will represent any files, aliases & Directories inside.
type Dir struct {
	Name     string
	Dirs     *[]Dir
	Files    *[]File
	SymLinks *[]SymLink
}

//////////////////////////////////////////////////////////////////
// Functions that work for all file explorer items (dir, file, symLink).

// GetType returns "dir", "file", or "symlink" from on the path given.
func GetType(itemPath string) (string, error) {
	// obtain item info from path
	item, err := os.Stat(itemPath)

	// if item does not exist
	if os.IsNotExist(err) {
		return "", errors.New("Item does not exist: " + itemPath)
	}

	// if symbolic link
	if itemSL, err := os.Lstat(itemPath); err == nil && itemSL.Mode()&os.ModeSymlink == os.ModeSymlink {
		return "symlink", nil // MAY WANT TO DIFFERENTIATE BETWEEN file & dir symlinks IN FUTURE
	}

	// if item is a dir
	if item.Mode().IsDir() {
		return "dir", nil
	}

	// if item is a file
	if item.Mode().IsRegular() {
		return "file", nil
	}

	// else error occurred
	return "", errors.New("Error analysing: " + itemPath)
}

// Rename a specified item.
func Rename(itemPath, newName string) error {
	// extract directory from item path
	dirPath := filepath.Dir(itemPath) + "/"

	// return any errors from renaming file
	return os.Rename(itemPath, dirPath+newName)
}

// Move an item to a specified path.
// This calls os.Rename(), but ensures the item is only moved.
func Move(itemPath, dirPath string) error {
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
	itemName := filepath.Base(itemPath) + "/"

	// return any errors from moving file
	return os.Rename(itemPath, dirPath+itemName)
}

// Delete will delete a specified item.
func Delete(itemPath string) error {
	// if item path is not a dir, remove normally
	if itemType, _ := GetType(itemPath); itemType != "dir" {
		return os.Remove(itemPath)
	}

	// else delete dir
	return os.RemoveAll(itemPath)
}
