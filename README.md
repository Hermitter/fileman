# Fileman
[![Build Status](https://travis-ci.com/Hermitter/fileman.svg?branch=master)](https://travis-ci.com/Hermitter/fileman)
[![codecov](https://codecov.io/gh/Hermitter/fileman/branch/master/graph/badge.svg)](https://codecov.io/gh/Hermitter/fileman)
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
- [x] Get Type

## Installation
```bash
go get -u github.com/Hermitter/fileman
```

## Usage
`More examples coming soon`
<details open>
<summary><b>Copy & Paste</b></summary>

```go
// Copy: returns a struct that allows you to edit/view the name & contents of an item
copiedDir, err := fileman.CopyDir("/home/john/Desktop/someDirectory")
if err != nil {
  return err
}

copiedFile, err := fileman.CopyFile("/home/john/Desktop/someFile.txt")
if err != nil {
  return err
}

copiedSymLink, err := fileman.CopySymLink("/home/john/Desktop/someSymLink.txt")
if err != nil {
  return err
}

// Paste: will not overwrite existing items
// the names of the pasted items are copiedDir.Name, copiedFile.Name, copiedSymLink.Name
if err := copiedDir.Paste("/home/john/documents", false); err != nil {
  return err
}

if err := copiedFile.Paste("/home/john/documents", false); err != nil {
  return err
}

if err := copiedSymLink.Paste("/home/john/documents"); err != nil {
  return err
}
```
</details>

<details open>
<summary><b>Rename, Move & Delete</b></summary>

```go
// Rename
err := fileman.Rename("./old.txt", "new.txt")
if err != nil {
  fmt.Println(err)
  return
}

// Move
err = fileman.Move("./new.txt", "/home/john/documents")
if err != nil {
fmt.Println(err)
  return
}

// Delete
err = fileman.Delete("/home/john/documents/new.txt")
if err != nil {
  fmt.Println(err)
  return
}
```
</details>

<details open>
<summary><b>Search</b></summary>

```go
found, path := fileman.Search("needle.txt", "/home/john/haystack", 1)
if !found {
  fmt.Println("Search Failed!")
  return
}

fmt.Println("Found it here: " + path)
```
</details>

<details open>
<summary><b>Get Type</b></summary>

```go
itemType, err := fileman.GetType("/home/john/media", false)

if err != nil {
  fmt.Println(err)
  return
}

fmt.Println(itemType)// Can print "dir" or "file"
```
</details>
