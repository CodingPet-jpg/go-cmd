package modget

import "github.com/CodingPet-jpg/go-cmd/base"

var CmdGet = &base.Command{
  UsageLine: "go get [-d] [-t] [-u] [-v] [build flags] [packages]",
	Short:     "add dependencies to current module and install them",
}

type upgradeFlag struct {
  rawVersion string
  version string
}

var (
  getU upgradeFlag
  getF = CmdGet.Flag.Bool("f",false,"")
)
