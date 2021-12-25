package modget

import (
  "strings"
  "fmt"
  "path/filepath"
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
  // 优先使用命令行参数传递的版本,其次选择go get -u指定的版本,默认为upgrade
  if version == "" {
    if getU.version = "" {
      version="upgrade"
    }else{
      version=getU.version
    }
  }
  q:=&query{
    raw:raw,
    rawVersion:rawVers,
    pattern:pattern,
    version:version,
    patternIsLocal:filepath.IsAbs(pattern),
  }
  return q,nil
}
