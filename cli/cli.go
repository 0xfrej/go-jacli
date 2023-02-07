package cli

import (
	"errors"
	"fmt"
	"github.com/lai0n/go-jacli/cli/arg"
	"github.com/lai0n/go-jacli/cli/flag"
)

var (
	FlagNotFound = errors.New("flag not found")
)

type Ctx struct {
	flagMap        map[string]flag.Flag
	rootCommand    CommandInterface
	currentCommand CommandInterface
	values         []string
}

func (c *Ctx) Flags() map[string]flag.Flag {
	return c.flagMap
}

func (c *Ctx) setFlags(flagMap map[string]flag.Flag) {
	c.flagMap = flagMap
}

func (c *Ctx) Flag(name string) (flag.Flag, error) {
	if v, ok := c.flagMap[name]; ok {
		return v, nil
	}

	return nil, FlagNotFound
}

func (c *Ctx) IsFlagSet(name string) bool {
	f, err := c.Flag(name)
	if err != nil {
		return false
	}
	return f.IsSet()
}

func (c *Ctx) IsFlagRequired(name string) bool {
	f, err := c.Flag(name)
	if err != nil {
		return false
	}
	return f.IsRequired()
}

func (c *Ctx) RootCommand() CommandInterface {
	return c.rootCommand
}

func (c *Ctx) CurrentCommand() CommandInterface {
	return c.currentCommand
}

func (c *Ctx) Values() []string {
	return c.values
}

func (c *Ctx) setValues(values []string) {
	c.values = values
}

func (c *Ctx) setCurrentCommand(cmd CommandInterface) {
	c.currentCommand = cmd
}

func newCliCtx(
	flagMap map[string]flag.Flag,
	rootCommand CommandInterface,
) *Ctx {
	return &Ctx{
		flagMap:        flagMap,
		rootCommand:    rootCommand,
		currentCommand: rootCommand,
	}
}

type CommandHandlerFunc func(*Ctx) Result

type CommandInterface interface {
	Flags() []flag.Flag
	Commands() []CommandInterface
	CommandName() string
	CommandDescription() string
	HandlerFunc() CommandHandlerFunc
}

type Command struct {
	Name        string
	FlagSet     []flag.Flag
	Description string
	SubCommands []CommandInterface
	Handler     CommandHandlerFunc
}

func (c *Command) CommandName() string {
	return c.Name
}

func (c *Command) Flags() []flag.Flag {
	return c.FlagSet
}

func (c *Command) CommandDescription() string {
	return c.Description
}

func (c *Command) Commands() []CommandInterface {
	return c.SubCommands
}

func (c *Command) HandlerFunc() CommandHandlerFunc {
	return c.Handler
}

type Result struct {
	value  interface{}
	errors []error
}

func (r *Result) Value() interface{} {
	return r.value
}

func (r *Result) HasErrors() bool {
	return r.errors != nil && len(r.errors) > 0
}

func (r *Result) Errors() []error {
	return r.errors
}

func NilResult() Result {
	return Result{}
}

func ErrResult(err []error) Result {
	return Result{errors: err}
}

func ValueResult(v any) Result {
	return Result{value: v}
}

type CLI struct {
	args        []string
	GlobalFlags []flag.Flag
	Description string
	CommandSet  []CommandInterface
	Handler     CommandHandlerFunc
}

func (cli *CLI) CommandName() string {
	return ""
}

func (cli *CLI) CommandDescription() string {
	return cli.Description
}

func (cli *CLI) Run(args []string) Result {
	iter := arg.NewArgIterator(args)

	ctx := newCliCtx(nil, cli)
	errs := parse(ctx, iter)
	if errs != nil {
		return Result{
			errors: errs,
		}
	}

	h := ctx.CurrentCommand().HandlerFunc()
	if h != nil {
		return h(ctx)
	} else {
		return Result{
			errors: []error{fmt.Errorf("handler for command '%s' does not exist", ctx.CurrentCommand().CommandName())},
		}
	}
}

func (cli *CLI) Flags() []flag.Flag {
	return cli.GlobalFlags
}

func (cli *CLI) Commands() []CommandInterface {
	return cli.CommandSet
}

func (cli *CLI) HandlerFunc() CommandHandlerFunc {
	return cli.Handler
}
