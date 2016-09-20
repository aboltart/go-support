package support_slice

func StringInSlice(value string, list []string) bool {
  for _, list_val := range list {
    if list_val == value {
        return true
    }
  }
  return false
}

func CompactStringSlice(source []string) []string {
  var result []string
  for _, str := range source {
    if str != "" {
      result = append(result, str)
    }
  }

  return result
}

func RemoveFromSlice(value string, source []string) []string {
  deletedElement := false

  for {
    for i, str := range source {
      // If found matching value
      if str == value {
        // Delete element from slice, and break search to start again from beggining
        source = append(source[:i], source[i+1:]...)
        deletedElement = true
        break
      }
    }

    // If searching element was not deleted from slice, then stop infinite loop
    if !deletedElement { break }
    deletedElement = false
  }

  return source
}

func PrependForSlice(value string, source []string) []string {
  source = append(source, "")
  copy(source[1:], source)
  source[0] = value

  return source
}
