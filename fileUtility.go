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
func Paste(file File, filePath string, sync bool) error {
	// create empty file
	newFile, err := os.Create(filePath)
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

func validFolder(folderPath string) error {
	// throw error if folder Path doesn't exist || points to file
	if pathType, err := os.Stat(folderPath); os.IsNotExist(err) || pathType.Mode().IsRegular() {
		// throw any os related errors
		if err != nil {
			return err
		}

		// notify that folder is invalid
		return errors.New("You did not specify a folder")
	}

	// folder is valid
	return nil
}

// Move a file to a specified path
func Move(filePath, folderPath string) error {
	// throw error if newPath doesn't exist || points to file
	if pathType, err := os.Stat(folderPath); os.IsNotExist(err) || pathType.Mode().IsRegular() {
		// throw any path validating errors
		if err != nil {
			return err
		}

		// Notify user of mistake
		return errors.New("You did not specify a folder")
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

// Search Inside (x) Number Of Folders For File
func Search(desiredFile string, folderPath string, searchDepth int) (string, error) {
	// // throw error if folder is invalid
	// if err := validFolder(folderPath); err != nil {
	// 	return "", err
	// }

	// // throw error if search depth reached
	// if searchDepth < 0 {
	// 	return "", nil
	// }

	// // read current folder's contents (files & folders)
	// folders, _ := ioutil.ReadDir(folderPath)

	// // for each folder
	// for _, item := range folders {
	// 	// update current folder
	// 	currentFolder := folderPath + "/" + item.Name()

	// 	// If item is a file & the desired file
	// 	if item.Mode().IsRegular() && item.Name() == desiredFile {
	// 		filePath = newSearchDir
	// 		return
	// 	}

	// 	// Run again
	// 	filePath = fileSearch(desiredFile, newSearchDir, searchDepth-1)
	// }

	return "", nil
}
