# Fileman

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

## Usage
`More examples coming soon`
<details open>
<summary><b>Copy & Paste</b></summary>

* Directories
```go
  // Copy
  newDir, _ := fileman.CopyDir("/home/jake/documents")
  newDir.Name = "cloned_documents"

  // Paste
  err = dir.Paste("/home/jake", false)
  if err != nil {
    fmt.Println(err)
  }
```

* Files

* Symbolic Links

</details>

