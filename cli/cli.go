package cli

import (
	"os"
)

type CommandHandlerFunc func(flags []Flag)

type CommandInterface interface {
	Flags() []Flag
	Commands() []CommandInterface
	CommandName() string
	CommandDescription() string
	HandlerFunc() CommandHandlerFunc
}

type Command struct {
	Name          string
	FlagSet       []Flag
	Description   string
	SubCommandSet []CommandInterface
	Handler       CommandHandlerFunc
}

func (c *Command) CommandName() string {
	return c.Name
}

func (c *Command) Flags() []Flag {
	return c.FlagSet
}

func (c *Command) CommandDescription() string {
	return c.Description
}

func (c *Command) SubCommands() []CommandInterface {
	return c.SubCommandSet
}

func (c *Command) HandlerFunc() CommandHandlerFunc {
	return c.Handler
}

type JacliInterface interface {
	CommandInterface

	Run() []error
}
type CLI struct {
	GlobalFlags []Flag
	Description string
	CommandSet  []CommandInterface
	Handler     CommandHandlerFunc
}

func (cli *CLI) HandlerFunc() CommandHandlerFunc {
	return cli.Handler
}

func (cli *CLI) CommandName() string {
	return os.Args[0]
}

func (cli *CLI) CommandDescription() string {
	return cli.Description
}

func (cli *CLI) Run() []error {
	return ParseInput(cli, os.Args[1:])
}

func (cli *CLI) Flags() []Flag {
	return cli.GlobalFlags
}

func (cli *CLI) Commands() []CommandInterface {
	return cli.CommandSet
}

func ParseInput(cli JacliInterface, argSet []string) []error {
	iterator := newArgsIterator(argSet)

	return parseSubCommand(cli, iterator, cli.Flags())
}

func parseSubCommand(cmd CommandInterface, iterator *ArgsIterator, extraFlags []Flag) []error {
	if arg, ok := iterator.Peek(); ok && !arg.IsFlag() {
		if cmd := findCommandFromName(arg.String(), cmd.Commands()); cmd != nil {
			iterator.Next()
			parseSubCommand(cmd, iterator, extraFlags)
		}
	}

	flags := append(extraFlags, cmd.Flags()...)
	errs := parseFlags(flags, iterator)
	if errs != nil {
		return errs
	}

	errs = validateFlags(flags)
	if errs != nil {
		return errs
	}

	cmd.HandlerFunc()(flags)
	return nil
}

func findCommandFromName(needle string, haystack []CommandInterface) CommandInterface {
	for _, v := range haystack {
		if v.CommandName() == needle {
			return v
		}
	}
	return nil
}
