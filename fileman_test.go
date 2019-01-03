package fileman

import (
	"os"
	"testing"
)

// TestMain creates a file for testing
func TestMain(m *testing.M) {
	// create test directories for File, Dir, Symlink
	os.MkdirAll("./fTest", os.ModePerm)
	os.MkdirAll("./dTest", os.ModePerm)
	os.MkdirAll("./sTest", os.ModePerm)

	// run fileman tests
	code := m.Run()
	// delete test directory & exit
	os.RemoveAll("./fTest")
	os.RemoveAll("./dTest")
	os.RemoveAll("./sTest")
	os.Exit(code)
}
