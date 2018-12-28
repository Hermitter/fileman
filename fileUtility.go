package fileman

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// File is a structure representing a single file.
type File struct {
	Name     string
	Contents []byte
}

// ToString returns the string value of a File's contents.
func (f File) ToString() string {
	return fmt.Sprintf("%s", f.Contents)
}

// Paste will paste a file inside a specified path.
// This will overwrite any file with the same name.
func (f File) Paste(path string, sync bool) error {
	// create empty file
	newFile, err := os.Create(path + "/" + f.Name)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// paste new file contents
	_, err = newFile.Write(f.Contents)
	if err != nil {
		return err
	}

	// if set, commit file contents to stable storage
	// setting this to true will trade performance for reliability
	if sync {
		newFile.Sync()
	}

	return nil
}

// CopyFile returns a File struct
// from a specified file path.
func CopyFile(path string) (File, error) {
	// initialize empty File
	file := File{"", []byte{}}

	// read & set file contents
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return file, err
	}

	// get & set file name from path
	file.Name = filepath.Base(path)
	file.Contents = contents

	return file, nil
}

// CloneFile will Copy & Paste a file into a specified path.
// The cloned file's name will be taken from the path given.
func CloneFile(path string, newPath string, sync bool) error {
	// copy file
	newFile, err := CopyFile(path)
	// set copied file's name from newPath
	newFile.Name = filepath.Base(newPath)
	// paste new file
	err = newFile.Paste(newPath, sync)

	return err
}

// CutFile will simultaneously Copy() & Delete()
// a specified file
func CutFile(path string) (File, error) {
	// copy specified file
	file, err := CopyFile(path)
	if err != nil {
		return file, err
	}

	// return copied file & any errors after deletion
	return file, Delete(path)
}

// SearchFile will search inside (x) number of directories for a file (WORK IN PROGRESS)
func SearchFile(desiredFile string, searchDir string, searchDepth int) (filePath string) {
	// Stop if search depth reached
	if searchDepth < 0 {
		return
	}

	// Read Current Directory Items
	dirs, _ := ioutil.ReadDir(searchDir)

	// For each item in Directory
	for _, item := range dirs {
		// Update current directory
		newSearchDir := searchDir + "/" + item.Name()

		// If item is a file & the desired file
		if item.Mode().IsRegular() && item.Name() == desiredFile {
			//fmt.Println(newSearchDir)
			filePath = newSearchDir
			return
		}

		// Run again
		//fmt.Println(newSearchDir)
		filePath = SearchFile(desiredFile, newSearchDir, searchDepth-1)

		if filePath != "" {
			return
		}
	}

	return
}
