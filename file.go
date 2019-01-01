package fileman

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"unicode/utf8"
)

// File is a structure representing a single file.
type File struct {
	Name     string
	Contents []byte
}

// ToString returns the string value of a File's contents.
// Will return an error if string is not utf8 valid
func (f *File) ToString() (string, error) {
	// format file text
	fileText := fmt.Sprintf("%s", f.Contents)

	// if utf8 valid, return string
	if utf8.ValidString(fileText) {
		return fileText, nil
	}

	// return error if not utf8 valid
	return "", errors.New("File " + f.Name + " contents do not contain a valid string")

}

// Paste will paste a file inside a specified path.
// This will not overwrite any existing paths.
func (f File) Paste(path string, sync bool) error {
	// where new file will be pasted
	pastePath := filepath.Join(path, f.Name)

	// if path is taken, return error
	if _, err := GetType(pastePath, false); err == nil {
		return errors.New("Path already exists: " + pastePath)
	}

	// create empty file
	newFile, err := os.Create(pastePath)
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

	// check if File exists
	itemType, err := GetType(path, true)
	if itemType != "file" {
		return file, errors.New("Not a valid file: " + path)
	} else if err != nil {
		return file, errors.New("Path is not valid: " + path)
	}

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
// The new path given should include the name of the file
func cloneFile(path string, newPath string, sync bool) error {
	// copy file
	newFile, err := CopyFile(path)
	if err != nil {
		return err
	}

	// set copied file's name from newPath
	newFile.Name = filepath.Base(newPath)
	// paste new file
	err = newFile.Paste(filepath.Dir(newPath), sync)

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
