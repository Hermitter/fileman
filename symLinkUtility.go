package fileman

import (
	"errors"
	"os"
	"path/filepath"
)

// CopySymLink returns a SymLink struct
// from a specified path.
func CopySymLink(path string) (SymLink, error) {
	// initialize empty SymLink
	symLink := SymLink{"", ""}

	// if path type is not symlink, return error
	if pType, err := GetType(path, true); err != nil || pType != "symlink" {
		return symLink, errors.New("Path is not a SymLink: " + path)
	}

	// get symlink's link
	link, err := os.Readlink(path)
	if err != nil {
		return symLink, err
	}

	// if symlink points to file, obtain full path
	if pType, _ := os.Stat(path); pType.Mode().IsRegular() {
		link = filepath.Dir(path) + "\\" + link
	}

	// read symlink path
	symLink.Name = filepath.Base(path)
	symLink.Link = link

	return symLink, nil
}

// PasteSymLink will paste a symLink inside a specified path.
// This will overwrite any symLink with the same name.
func PasteSymLink(symlink *SymLink, path string) error {
	// Attempt to create symlink
	err := os.Symlink(symlink.Link, path+"/"+symlink.Name)

	return err
}

// CloneSymLink will Copy & Paste a symlink into a specified path.
// The cloned symLink's name will be taken from the path given.
func CloneSymLink(path string, newPath string) error {
	// copy file
	newSymLink, err := CopySymLink(path)
	if err != nil {
		return err
	}

	// set file name from newFilePath
	newSymLink.Name = filepath.Base(newPath)
	// paste new file
	err = PasteSymLink(&newSymLink, filepath.Dir(newPath))

	return err
}

// CutSymLink will simultaneously Copy() & Delete()
// a specified symlink
func CutSymLink(path string) (SymLink, error) {
	// copy specified symlink
	symLink, err := CopySymLink(path)
	if err != nil {
		return symLink, err
	}

	// return copied file & any errors after deletion
	return symLink, Delete(path)
}
