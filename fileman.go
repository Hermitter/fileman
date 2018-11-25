package fileman

import "io/ioutil"

// **** INITIAL UPLOAD TO BE EDITED LATER **** \\
// - Search Inside (x) Number Of Folders For File
func fileSearch(desiredFile string, searchDir string, searchDepth int) (filePath string) {
	// Stop if search depth reached
	if searchDepth < 0 {
		return
	}

	// Read Current Directory Items
	folders, _ := ioutil.ReadDir(searchDir)

	// For Each Directory
	for _, item := range folders {
		// Update current directory
		newSearchDir := searchDir + "/" + item.Name()

		// If item is a file & the desired file
		if item.Mode().IsRegular() && item.Name() == desiredFile {
			filePath = newSearchDir
			return
		}

		// Run again
		//fmt.Println(newSearchDir)
		filePath = fileSearch(desiredFile, newSearchDir, searchDepth-1)
	}

	return
}
