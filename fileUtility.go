package fileman

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// CopyFile returns a File struct
// from a specified file path.
func CopyFile(path string) (File, error) {
	// initialize empty File
	var contents []byte
	file := File{"", &contents}

	// read & set file contents
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return file, err
	}

	// get & set file name from path
	file.Name = filepath.Base(path)

	return file, nil
}

// PasteFile will paste a file inside a specified path.
// This will overwrite any file with the same name.
func PasteFile(file *File, path string, sync bool) error {
	// create empty file
	newFile, err := os.Create(filepath.Dir(path) + "/" + file.Name)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// paste new file contents
	_, err = newFile.Write(*file.Contents)
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

// CloneFile will Copy & Paste a file into a specified path.
// The cloned file's name will be taken from the path given.
func CloneFile(path string, newPath string, sync bool) error {
	// copy file
	newFile, err := CopyFile(path)
	// set copied file's name from newPath
	newFile.Name = filepath.Base(newPath)
	// paste new file
	err = PasteFile(&newFile, newPath, sync)

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
