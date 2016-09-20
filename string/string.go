package support_string

import (
  // "fmt"
  "unicode"
  "strings"
)

func CharOfString(data string, position int) string {
  //http://stackoverflow.com/questions/15018545/how-to-index-characters-in-a-golang-string
  return string([]rune(data)[position])
}

func IsUpper(data string) bool {
  return unicode.IsUpper([]rune(data)[0])
}

// SnakeToCamel returns a string converted from snake case to uppercase
// Part of code from https://github.com/serenize/snaker/blob/master/snaker.go
func SnakeToCamel(s string) string {
  var result string

  words := strings.Split(s, "_")

  for _, word := range words {
    if len(word) > 0 {
      w := []rune(word)
      w[0] = unicode.ToUpper(w[0])
      result += string(w)
    }
  }

  return result
}

// CamelToSnake converts a given string to snake case
// Part of code from https://github.com/serenize/snaker/blob/master/snaker.go
func CamelToSnake(s string) string {
  var result string
  var words []string
  var lastPos int
  rs := []rune(s)

  for i := 0; i < len(rs); i++ {
    if i > 0 && unicode.IsUpper(rs[i]) {
      words = append(words, s[lastPos:i])
      lastPos = i
    }
  }

  // append the last word
  if s[lastPos:] != "" {
    words = append(words, s[lastPos:])
  }

  for k, word := range words {
    if k > 0 {
      result += "_"
    }

    result += strings.ToLower(word)
  }

  return result
}
