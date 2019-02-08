package fileman

import (
	"os"
	"path/filepath"
)

// SymLink is a structure representing a single symbolic link
type SymLink struct {
	Name string
	Link string
	Type string
}

// Paste will paste a symLink inside a specified path.
// This will overwrite any symLink with the same name.
func (s *SymLink) Paste(path string) error {
	// extract full path
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Attempt to create symlink
	return os.Symlink(s.Link, filepath.Join(path, s.Name))
}

// CopySymLink returns a SymLink struct
// from a specified path.
func CopySymLink(path string) (SymLink, error) {
	// initialize empty SymLink
	symLink := SymLink{"", "", ""}

	// extract full path
	path, err := filepath.Abs(path)
	if err != nil {
		return symLink, err
	}

	// get link of symlink
	link, err := os.Readlink(path)
	if err != nil {
		return symLink, err
	}

	// determine symLink type
	if pathType, _ := GetType(path, false); pathType == "file" {

		// create full file path
		link = filepath.Join(filepath.Dir(path), link)

		symLink.Type = "file"

	} else {
		symLink.Type = "dir"
	}

	// read symlink path
	symLink.Name = filepath.Base(path)
	symLink.Link = link

	return symLink, nil
}

// CloneSymLink will Copy & Paste a symlink into a specified path.
// The cloned symLink's name will be taken from the path given.
func cloneSymLink(path string, newPath string) error {
	// copy symLink
	newSymLink, err := CopySymLink(path)
	if err != nil {
		return err
	}

	// set symLink name from newPath
	newSymLink.Name = filepath.Base(newPath)
	// paste new symLink
	err = newSymLink.Paste(filepath.Dir(newPath))

	return err
}

// CutSymLink will simultaneously Copy() & Delete()
// a specified symlink.
func CutSymLink(path string) (SymLink, error) {
	// extract full path
	path, err := filepath.Abs(path)
	if err != nil {
		return SymLink{}, err
	}

	// copy specified symlink
	symLink, err := CopySymLink(path)
	if err != nil {
		return symLink, err
	}

	// return copied file & any errors after deletion
	return symLink, Delete(path)
}
