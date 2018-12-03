package fileman

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Copy returns a File struct
// from a specified file path.
func Copy(filePath string) (File, error) {
	// initialize empty File
	var contents []byte
	file := File{"", &contents}

	// read & set file contents
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return file, err
	}

	// set file name from path
	file.Name = filepath.Base(filePath)

	return file, nil
}

// Paste a file inside a specified path.
// This will overwrite any file with the same name.
func Paste(fileContents *[]byte, filePath string, sync bool) error {
	// create empty file
	newFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// paste new file contents
	_, err = newFile.Write(*fileContents)
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

// Delete will remove the specified file
func Delete(filePath string) error {
	return os.Remove(filePath)
}

// Cut will simultaneously Copy() & Delete()
// a specified file
func Cut(filePath string) (File, error) {
	// copy specified file
	file, err := Copy(filePath)
	if err != nil {
		return file, err
	}

	// return copied file & any errors after deletion
	return file, Delete(filePath)
}

// Rename a specified file
func Rename(filePath, newName string) error {
	// extract directory from path
	folderPath := filepath.Dir(filePath) + "/"

	// return any errors from renaming file
	return os.Rename(filePath, folderPath+newName)
}

// Move a file to a specified path
func Move(filePath, folderPath string) error {
	// throw error if newPath doesn't exist || points to file
	if err := validFolder(folderPath); err != nil {
		return err
	}

	// if folderPath does not end with "/"
	if folderPath[len(folderPath)-1:] != "/" {
		// add "/"
		folderPath = folderPath + "/"
	}

	// extract file name from path
	fileName := filepath.Base(filePath) + "/"

	// return any errors from moving file
	return os.Rename(filePath, folderPath+fileName)
}

// Search Inside (x) Number Of Folders For File (WORK IN PROGRESS)
func Search(desiredFile string, searchDir string, searchDepth int) (filePath string) {
	// Stop if search depth reached
	if searchDepth < 0 {
		return
	}

	// Read Current Directory Items
	folders, _ := ioutil.ReadDir(searchDir)

	// For each item in Directory
	for _, item := range folders {
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
		filePath = Search(desiredFile, newSearchDir, searchDepth-1)

		if filePath != "" {
			return
		}
	}

	return
}

// check if a folder path is valid
func validFolder(folderPath string) error {
	// throw error if folder Path doesn't exist || points to file
	if pathType, err := os.Stat(folderPath); os.IsNotExist(err) || pathType.Mode().IsRegular() {
		// throw any os related errors
		if err != nil {
			return errors.New("Folder does not exist: " + folderPath)
		}
		// notify that folder is invalid
		return errors.New("Invalid folder: " + folderPath)
	}

	// folder is valid
	return nil
}
