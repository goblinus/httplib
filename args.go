package app

import "os"

const (
	noArgs        = 0
	noSubCommands = 1

	commandIdx    = 0
	subCommandIdx = 1
)

type (
	CLIArgs struct {
		args []string
	}
)

func NewCLIArgs() *CLIArgs {
	return &CLIArgs{
		args: os.Args[1:],
	}
}

func (c CLIArgs) Command() string {
	if !c.HasCommand() {
		return ""
	}

	return c.args[commandIdx]
}

func (c CLIArgs) SubCommand() string {
	if !c.HasSubCommand() {
		return ""
	}

	return c.args[subCommandIdx]
}

func (c CLIArgs) HasCommand() bool {
	return len(c.args) > noArgs
}

func (c CLIArgs) HasSubCommand() bool {
	return len(c.args) > noSubCommands
}
