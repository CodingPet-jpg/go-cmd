package base

import (
  "context"
  "flag"
)

type Command struct{
  Run func (ctx context.Context,cmd *Command,args []string)
  UsageLine string
  Short string
  Long string
  Flag flag.FlagSet
  CustomFlags bool
  Commands []*Command
}
