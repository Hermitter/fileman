package fileman

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hermitter/fileman"
)

// TestMain creates a file for testing
func TestMain(m *testing.M) {
	// create test file
	err := ioutil.WriteFile("file.txt", []byte("hello world"), 0644)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	// run fileman tests
	code := m.Run()

	// delete test file & exit
	os.Remove("file.txt")
	os.Exit(code)
}

// TestCopy does things
func TestCopy(t *testing.T) {
	// copy a nonexistent file
	_, err := fileman.CopyFile("fakeFile.txt")
	if err == nil {
		t.Error("A nonexistent file was copied.")
	}

	// copy an existing file
	newFile, err := fileman.CopyFile("file.txt")
	if err != nil {
		t.Error(err)
	}
	// check if contents were copied
	if newFile.ToString() != "hello world" {
		t.Error("Content from test file was not copied correctly.")
	}
}

func TestPaste(t *testing.T) {

}
