package testfileman

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hermitter/fileman"
)

var newFile = fileman.File{}

// TestMain creates a file for testing
func TestMain(m *testing.M) {
	// create test file
	err := ioutil.WriteFile("file.txt", []byte("hello world"), 0644)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	// run fileman tests
	code := m.Run()

	os.Exit(code)
}

func TestCopy(t *testing.T) {
	// copy a nonexistent file
	_, err := fileman.CopyFile("fakeFile.txt")
	if err == nil {
		t.Error("A nonexistent file was copied.")
	}

	// copy an existing file
	newFile, err = fileman.CopyFile("file.txt")
	if err != nil {
		t.Error(err)
	}
	// check if contents were copied
	if newFile.ToString() != "hello world" {
		t.Error("Content from test file was not copied correctly.")
	}
}

func TestPaste(t *testing.T) {
	// paste File with no name \\
	newFile.Name = ""
	err := newFile.Paste("", false)
	if err == nil {
		t.Error("A File with no name was pasted")
	}

	// paste a new file with path already taken \\
	newFile.Name = "file.txt"
	err = newFile.Paste("./", false)
	if err == nil {
		t.Error(err)
	}

	// paste a valid file \\
	newFile.Name = "file2.txt"
	newFile.Contents = []byte("goodbye world")
	err = newFile.Paste("./", false)
	if err != nil {
		t.Error(err)
	}
	// check if content was added
	if newFile, _ = fileman.CopyFile("file2.txt"); newFile.ToString() != "goodbye world" {
		t.Error("Paste did not match contents given")
	}
	// delete file
	fileman.Delete("file2.txt")

}

// func TestPasteOverwrite(t *testing.T) {
// 	if itemType, _ := fileman.GetType("./file.txt", false); itemType != "file" {
// 		t.Error("file was not pasted")
// 		os.Exit(0)
// 	}
// }

func TestCut(t *testing.T) {
	// cut recently pasted file
	_, err := fileman.CutFile("./file.txt")
	if err != nil {
		t.Error(err)
	}

	// verify cut file was deleted
	if _, err := fileman.GetType("./file.txt", false); err == nil {
		t.Error("Cut file was not deleted")
	}
}
