package base

import (
  "context"
  "flag"
)

type Command struct{
  // 运行命令
  Run func (ctx context.Context,cmd *Command,args []string)
  UsageLine string
  Short string
  Long string
  Flag flag.FlagSet
  CustomFlags bool
  Commands []*Command
}

var Go = &Command {
  UsageLine:"go",
  Long:`Go is a tool for managing Go source code`,
}
