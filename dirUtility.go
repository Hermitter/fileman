package fileman

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Dir is a structure representing a single Directory.
// 4 slices will represent any files, directories, & symbolic links inside.
type Dir struct {
	Name     string
	Dirs     []Dir
	Files    []File
	SymLinks []SymLink
}

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

// Paste will do something FILL IN LATER...
func (d Dir) Paste(path string, sync bool) error {
	// create initial directory
	err := os.Mkdir(path+"/"+d.Name, os.ModePerm)
	if err != nil {
		return err
	}

	// MAKE RECURSIVE
	// // for Each directory inside
	// for i := range d.Dirs {
	// 	newDir := d.Dirs[i]
	// 	dirPath := path + "/" + d.Name + "/" + newDir.Name
	// 	// create the directory
	// 	os.Mkdir(dirPath, os.ModePerm)
	// 	// paste each File
	// 	for f := range newDir.Files {
	// 		newDir.Files[f].Paste(dirPath+"/"+newDir.Files[f].Name, sync)
	// 	}
	// 	// paste each SymLink
	// }

	return nil
}
