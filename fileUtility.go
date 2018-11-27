package fileman

import (
	"io/ioutil"
	"path/filepath"
)

// Copy returns a File struct
// from a specified file path.
func Copy(path string) (File, error) {
	// Initialize empty File
	var contents []byte
	file := File{"", &contents}

	// attempt to read file contents
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return file, err
	}

	// save file name
	file.Name = filepath.Base(path)

	return file, nil
}
