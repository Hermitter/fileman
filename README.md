# Fileman

Fileman is a go package that provides handy functions for **files**, **directories**, and **symbolic links**.

GoDocs: https://godoc.org/github.com/Hermitter/fileman

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
  newDir, _ := fileman.CopyDir("/home/jake/documents")
  newDir.Name = "cloned_documents"

  // Remove last file from directory
  newDir.Files = newDir.Files[:len(newDir.Files)-1]

  // Paste directory
  err := newDir.Paste("/home/jake", false)
  if err != nil {
    fmt.Println(err)
  }
}
```
</details>

<details close>
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

