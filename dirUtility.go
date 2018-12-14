package fileman

// CopyDir returns a Directory struct
// from a specified path.
func CopyDir(path string) (Dir, error) {
	// initialize empty Dir
	dir := Dir{"", &[]Dir{}, &[]File{}, &[]SymLink{}, &[]SymLink{}}

	return dir, nil
}
