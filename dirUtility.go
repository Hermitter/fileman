package fileman

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// CopyDir returns a Directory struct
// from a specified path.
func CopyDir(path string) (Dir, error) {
	// initialize empty Dir
	dir := Dir{"", &[]Dir{}, &[]File{}, &[]SymLink{}, &[]SymLink{}}

	return dir, nil
}

// SortDirItems dsd s
func SortDirItems(path string) (Dir, error) {
	// initialize empty dir
	dir := Dir{filepath.Base(path), &[]Dir{}, &[]File{}, &[]SymLink{}, &[]SymLink{}}
	// get current directory items
	paths, _ := ioutil.ReadDir(path)

	// for each dir item
	for _, item := range paths {
		fmt.Println(item.Name()) // CC TEST
		// get path type
		switch pType, _ := GetType(path+item.Name(), true); pType {

		// if file, copy to dir's file list
		case "file":
			newFile, _ := CopyFile(path + item.Name())
			*dir.Files = append(*dir.Files, newFile)

		// if directory, copy to dir's dir list
		case "dir":
			////////////////////////////////////////
			// FIX LOGIC TO MAKE RECURSIVE
			newDir, _ := SortDirItems(path + item.Name())
			*dir.Dirs = append(*dir.Dirs, newDir)
			//////////////////////////////////////////

		// if directory, copy to dir's symlink lists
		case "symlink":
			newSymLink, _ := CopySymLink(path + item.Name())
			// if linked to file, append to FileSymLinks
			if newSymLink.Type == "file" {
				*dir.FileSymLinks = append(*dir.FileSymLinks, newSymLink)
			} else {
				// else append to FileSymLinks
				*dir.DirSymLinks = append(*dir.DirSymLinks, newSymLink)
			}
		// return error
		default:
			return dir, errors.New("Could not determine path type: " + path)
		}
	}
	return dir, nil
}
