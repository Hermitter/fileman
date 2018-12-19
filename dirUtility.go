package fileman

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// CopyDir returns a Directory struct
// from a specified path.
func CopyDir(path string) (Dir, error) {
	// prevent broken path ex. /homeIShouldBeSeperate.txt
	path += "/"
	// initialize empty dir
	dir := Dir{filepath.Base(path), []Dir{}, []File{}, []SymLink{}}
	// get current directory items
	paths, _ := ioutil.ReadDir(path)

	// for each dir item
	for _, item := range paths {
		// get path type
		switch pType, _ := GetType(path+item.Name(), true); pType {

		// if file, copy to dir's file list
		case "file":
			newFile, _ := CopyFile(path + item.Name())
			dir.Files = append(dir.Files, newFile)

		// if directory, copy to dir's dir list
		case "dir":
			newDir, _ := CopyDir(path + item.Name())
			dir.Dirs = append(dir.Dirs, newDir)

		// if directory, copy to dir's symlink lists
		case "symlink":
			newSymLink, _ := CopySymLink(path + item.Name())
			dir.SymLinks = append(dir.SymLinks, newSymLink)

		// return error
		default:
			return dir, errors.New("Could not determine path type: " + path)
		}
	}

	return dir, nil
}

// PasteDir will do...
func PasteDir() error {
	return nil
}
