package fileman

import (
	"io/ioutil"
	"os"
	"path/filepath"
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
	ioutil.WriteFile("./fileman/getTypeFile", []byte("Sup brah"), 0644)
	err := os.Mkdir("./fileman/getTypeDir", os.ModePerm)
	linkPath, err := filepath.Abs("./fileman/getTypeDir")
	err = os.Symlink(linkPath, "./fileman/getTypeSymLink")

	if err != nil {
		t.Error(err)
	}

	// valid GetType
	if itemType, _ := GetType("./fileman/getTypeFile", false); itemType != "file" || err != nil {
		t.Error(err, "GetType did not detect file")
	}

	if itemType, _ := GetType("./fileman/getTypeDir", false); itemType != "dir" || err != nil {
		t.Error(err, "GetType did not detect directory")
	}

	if itemType, _ := GetType("./fileman/getTypeSymLink", true); itemType != "symLink" || err != nil {
		t.Error(err, "GetType did not detect symLink")
	}

	// invalid GetType
	if itemType, err := GetType("./fileman/fake", true); itemType != "" && err == nil {
		t.Error("GetType did not fail correctly for non-Existing file")
	}
}

func TestRename(t *testing.T) {
	// create test items (file, dir, symLink)
	ioutil.WriteFile("./fileman/renameFile", []byte("Sup brah"), 0644)
	err := os.Mkdir("./fileman/renameDir", os.ModePerm)
	linkPath, err := filepath.Abs("./fileman/renameDir")
	err = os.Symlink(linkPath, "./fileman/renameSymLink")

	if err != nil {
		t.Error(err)
	}

	// valid rename
	if err := Rename("./fileman/renameFile", "renamedFile"); err != nil {
		t.Error(err)
	}
	if err := Rename("./fileman/renameDir", "renamedDir"); err != nil {
		t.Error(err)
	}
	if err := Rename("./fileman/renameSymLink", "renamedSymLink"); err != nil {
		t.Error(err)
	}

	// invalid rename
	if err := Rename("./fileman/NotAPath", "./fileman/renameShouldFail"); err == nil {
		t.Error("Rename tried to use an invalid path")
	}
}

func TestMove(t *testing.T) {
	// create test items (file, dir, symLink)
	ioutil.WriteFile("./fileman/moveFile", []byte("Sup brah"), 0644)
	err := os.Mkdir("./fileman/moveDir", os.ModePerm)
	linkPath, err := filepath.Abs("./fileman/moveDir")
	err = os.Symlink(linkPath, "./fileman/moveSymLink")

	err = os.Mkdir("./fileman/goal", os.ModePerm) // move destination

	if err != nil {
		t.Error(err)
	}

	// valid move
	if err := Move("./fileman/moveFile", "./fileman/goal"); err != nil {
		t.Error(err)
	}
	if err := Move("./fileman/moveDir", "./fileman/goal"); err != nil {
		t.Error(err)
	}
	if err := Move("./fileman/moveSymLink", "./fileman/goal"); err != nil {
		t.Error(err)
	}

	// invalid move
	if err := Move("./fileman/NotAPath", "./fileman/shouldFail"); err == nil {
		t.Error("Move tried to use an invalid path")
	}
}

func TestDelete(t *testing.T) {
	// create test items (file, dir, symLink)
	ioutil.WriteFile("./fileman/deleteFile", []byte("Sup brah"), 0644)
	err := os.Mkdir("./fileman/deleteDir", os.ModePerm)
	linkPath, err := filepath.Abs("./fileman/deleteDir")
	err = os.Symlink(linkPath, "./fileman/deleteSymLink")

	if err != nil {
		t.Error(err)
	}

	// valid delete
	if err := Delete("./fileman/deleteFile"); err != nil {
		t.Error(err)
	}
	if err := Delete("./fileman/deleteDir"); err != nil {
		t.Error(err)
	}
	if err := Delete("./fileman/deleteSymLink"); err != nil {
		t.Error(err)
	}

	// invalid delete
	if err := Delete("./fileman/NotAPath"); err != nil {
		t.Error("Delete tried to use an invalid path")
	}
}
