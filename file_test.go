package fileman

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	// create test file
	ioutil.WriteFile("./fTest/copy.txt", []byte("Hello "), 0644)

	// copy a nonexistent file
	_, err := CopyFile("badCopy.txt")
	if err == nil {
		t.Error("A nonexistent file was copied.")
	}

	// copy test file & edit copied contents
	newFile, err := CopyFile("./ftest/copy.txt")
	if err != nil {
		t.Error(err)
	}
	newFile.Contents = append(newFile.Contents, []byte("World")...)

	// check if contents were copied
	text, _ := newFile.ToString()
	if text != "Hello World" {
		t.Error("Content from test file was not copied correctly.")
	}

}

func TestPaste(t *testing.T) {
	// create test file
	ioutil.WriteFile("./fTest/paste.txt", []byte("Hello "), 0644)

	// copy & delete test file
	newFile, _ := CopyFile("./fTest/paste.txt")
	os.Remove("./fTest/paste.txt")

	// test valid paste
	err := newFile.Paste("./fTest", false)
	if err != nil {
		t.Error(err)
	}

	// test invalid paste with empty struct
	newFile = File{}
	err = newFile.Paste("./fTest", false)
	if err == nil {
		t.Error("Tried to paste an empty File struct")
	}
}

func TestClone(t *testing.T) {
	// create test file
	ioutil.WriteFile("./fTest/clone.txt", []byte("Hello"), 0644)

	// test valid clone
	err := cloneFile("./fTest/clone.txt", "./fTest/newClone.txt", false)
	if err != nil {
		t.Error(err)
	}
	// check if new file was created
	if _, err := os.Stat("./fTest/newClone.txt"); os.IsNotExist(err) {
		t.Error("newClone.txt was not cloned")
	}

	// test invalid clone by copying a dir
	err = cloneFile("./fTest", "./fTest/newClone.txt", false)
	if err == nil {
		t.Error("Path to clone was not a file")
	}

}

func TestCut(t *testing.T) {
	// create test file
	ioutil.WriteFile("./fTest/cut.txt", []byte("Hello"), 0644)

	// test valid cut
	newFile, err := CutFile("./fTest/cut.txt")
	if err != nil {
		t.Error(err)
	}

	// check if content was copied
	if text, _ := newFile.ToString(); newFile.Name != "cut.txt" && text != "Hello" {
		t.Error("cut.txt was not properly copied")
	}

	// test invalid cut
	newFile, err = CutFile("./fTest/badCut.txt")
	if err == nil {
		t.Error("Cannot cut and invalid path")
	}
}
