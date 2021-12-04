package modget

import (
  "strings"
  "fmt"
)


type query struct {
  raw string
  rawVersion string
  pattern string
  patternIsLocal bool
  version string
}

func newQuery(raw string) (*query,error)  {
  pattern:=raw
  rawVers:=""
  if i:=strings.Index(raw,"@");i>=0 {
    pattern,rawVers = raw[:i],raw[i+1:]
    // 不允许@后无内容，也不允许一个pattern中出现多个@
    if strings.Contains(rawVers,"@")||rawVers=="" {
      return nil,fmt.ErrorOf("invalid module version syntax %q",raw)
    }
  }
  version:=rawVers
  return nil,nil
}
