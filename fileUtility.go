package fileman

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Copy returns a File struct
// from a specified file path.
func Copy(path string) (File, error) {
	// Initialize empty File
	var contents []byte
	file := File{"", &contents}

	// read file contents
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return file, err
	}

	// save file name
	file.Name = filepath.Base(path)

	return file, nil
}

// Paste creates a file inside a specified path.
// This will overwrite any file with the same name.
func Paste(file File, path string, sync bool) error {
	// create empty file
	newFile, err := os.Create(path)
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
	if sync {
		newFile.Sync()
	}

	return nil
}
