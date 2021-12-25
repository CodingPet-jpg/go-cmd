package search

import (
  "strings"
)

func IsRelativePath(pattern string) bool  {
  return strings.HasPrefix(pattern,"../") || strings.HasPrefix(pattern,"./") || pattern=".." || pattern="."
}
