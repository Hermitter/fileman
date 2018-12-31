# Fileman
[![Build Status](https://travis-ci.com/Hermitter/fileman.svg?branch=master)](https://travis-ci.com/Hermitter/fileman)
[![GoDoc](https://godoc.org/github.com/Hermitter/fileman?status.svg)](https://godoc.org/github.com/Hermitter/fileman)
[![Go Report Card](https://goreportcard.com/badge/github.com/hermitter/fileman)](https://goreportcard.com/report/github.com/hermitter/fileman)

Fileman is a go package that provides handy functions for **files**, **directories**, and **symbolic links**.

## Current Functionality
- [x] Copy
- [x] Paste
- [x] Delete
- [x] Cut
- [x] Move
- [x] Rename
- [x] Search
- [x] Duplicate

## Installation
```bash
go get -u github.com/Hermitter/fileman
```

## Usage
`More examples coming soon`
<details open>
<summary><b>Copy & Paste</b></summary>

```go
// Directory Example
func main() {
  // Copy directory
  newDir, err := fileman.CopyDir("/home/john/documents")
  if err != nil {
    fmt.Println(err)
    os.Exit(3)
  }

  // Rename copied directory
  newDir.Name = "cloned_documents"
  // Remove last file from directory
  newDir.Files = newDir.Files[:len(newDir.Files)-1]

  // Paste directory
  err = newDir.Paste("/home/john", false)
  if err != nil {
    fmt.Println(err)
    os.Exit(3)
  }
}
```
</details>

<details open>
<summary><b>Delete</b></summary>

```go
func main(){
  err := fileman.Delete("/home/jake/documents/myFile.txt")
  if err != nil {
    fmt.Println(err)
  }
}
```
</details>

