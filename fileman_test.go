package fileman

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestMain creates a file for testing
func TestMain(m *testing.M) {
	// create test directories for File, Dir, Symlink
	os.MkdirAll("./fTest", os.ModePerm)
	os.MkdirAll("./dTest", os.ModePerm)
	os.MkdirAll("./sTest", os.ModePerm)
	os.MkdirAll("./fileman", os.ModePerm)

	// run fileman tests
	code := m.Run()
	// delete test directory & exit
	os.RemoveAll("./fTest")
	os.RemoveAll("./dTest")
	os.RemoveAll("./sTest")
	os.RemoveAll("./fileman")
	os.Exit(code)
}

func TestGetType(t *testing.T) {
	// create test items (file, dir, symLink)
	ioutil.WriteFile("./fileman/file", []byte("Sup brah"), 0644)
	err := os.Mkdir("./fileman/dir", os.ModePerm)
	err = os.Symlink("./fileman/file", "./fileman/symLink")

	if err != nil {
		t.Error(err)
	}

	// valid tests
	if itemType, _ := GetType("./fileman/file", false); itemType != "file" || err != nil {
		t.Error("GetType did not detect file")
	}

	if itemType, _ := GetType("./fileman/dir", false); itemType != "dir" || err != nil {
		t.Error("GetType did not detect directory")
	}

	if itemType, _ := GetType("./fileman/symLink", true); itemType != "symLink" || err != nil {
		t.Error("GetType did not detect symLink")
	}

	// invalid test
	if itemType, err := GetType("./fileman/fake", true); itemType != "" && err == nil {
		t.Error("GetType did not fail correctly for non-Existing file")
	}
}
