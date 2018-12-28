package fileman

import (
	"errors"
	"fmt"
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
	// prevent broken path ex. /homeMyFile.txt --> /home/MyFile.txt
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

// Paste will paste a Directory inside a specified path.
// This will not overwrite a Directory with the same name.
func (d Dir) Paste(path string, sync bool) error {
	// create new directory
	dirPath := path + "/" + d.Name
	err := os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Paste each file inside directory
	for i := range d.Files {
		err := d.Files[i].Paste(dirPath, sync)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Paste each symbolic link inside directory
	for i := range d.SymLinks {
		err := d.SymLinks[i].Paste(dirPath)
		if err != nil {
			fmt.Println(err)
		}
	}

	// for each direcotry
	for i := range d.Dirs {
		err := d.Dirs[i].Paste(dirPath, sync)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
