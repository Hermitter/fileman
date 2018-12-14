package fileman

import (
	"errors"
	"io/ioutil"
)

// CopyDir returns a Directory struct
// from a specified path.
func CopyDir(path string) (Dir, error) {
	// initialize empty Dir
	dir := Dir{"", &[]Dir{}, &[]File{}, &[]SymLink{}, &[]SymLink{}}

	return dir, nil
}

// SortDirItems dsd s
func SortDirItems(path string) (Dir, error) {
	// initialize empty dir
	dir := Dir{"", &[]Dir{}, &[]File{}, &[]SymLink{}, &[]SymLink{}}
	// get current directory items
	paths, _ := ioutil.ReadDir(path)

	// for each dir item
	for _, item := range paths {
		// get path type
		switch pType, _ := GetType(path+item.Name(), true); pType {

		// if file, copy to dir's file list
		case "file":
			newFile, _ := CopyFile(path + item.Name())
			*dir.Files = append(*dir.Files, newFile)

		// if directory, copy to dir's dir list
		case "dir":
			//fmt.Println("Found A Dir")
			// if  symbolic link

		// if directory, copy to dir's symlink lists
		case "symlink":
			newSymLink, _ := CopySymLink(path + item.Name())
			// get symlink type
			switch pType, err := GetType(path+item.Name(), false); pType {
			case "file":
				*dir.FileSymLinks = append(*dir.FileSymLinks, newSymLink)
			case "dir":
				*dir.DirSymLinks = append(*dir.DirSymLinks, newSymLink)
			default:
				return dir, err
			}
		default:
			return dir, errors.New("Could not determine path type: " + path)
		}
	}
	return dir, nil
}
