package fileman

// File is a structure representing a single file.
type File struct {
	Name     string
	Contents *[]byte
}

// Folder is a structure representing a single folder.
// 2 slices will represent any files & folders inside the folder.
type Folder struct {
	Name    string
	Folders *[]Folder
	Files   *[]File
}
