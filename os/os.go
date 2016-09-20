package os

import (
  "os"
)

func ExistsPath(path string) bool {
  if _, err := os.Stat(path); err != nil {
    if os.IsNotExist(err) {
      return false
    }
  }
  return true
}

func IsDirectory(path string) bool {
  if ExistsPath(path) {
    fdir, _ := os.Open(path)
    defer fdir.Close()
    finfo, _ := fdir.Stat()
    return finfo.Mode().IsDir()
  } else {
    return false
  }
}

func IsFile(path string) bool {
  if ExistsPath(path) {
    return !IsDirectory(path)
  } else {
    return false
  }
}
