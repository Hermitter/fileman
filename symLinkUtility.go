package fileman

import (
	"errors"
	"os"
	"path/filepath"
)

// CopySymLink returns a SymLink struct
// from a specified path.
func CopySymLink(symLinkPath string) (SymLink, error) {
	// initialize empty SymLink
	symLink := SymLink{"", ""}

	// if item type is not symlink, return error
	if iType, err := GetType(symLinkPath); err != nil || iType != "symlink" {
		return symLink, errors.New("Item is not a SymLink: " + symLinkPath)
	}

	// get symlink's link
	link, err := os.Readlink(symLinkPath)
	if err != nil {
		return symLink, err
	}

	// if symlink points to file, obtain full path
	if iType, _ := os.Stat(symLinkPath); iType.Mode().IsRegular() {
		link = filepath.Dir(symLinkPath) + "\\" + link
	}

	// read symlink path
	symLink.Name = filepath.Base(symLinkPath)
	symLink.Link = link

	return symLink, nil
}
