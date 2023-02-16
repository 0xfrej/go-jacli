package cli

import (
	"errors"
	"fmt"
	"github.com/lai0n/go-jacli/cli/arg"
	"github.com/lai0n/go-jacli/cli/flag"
	"strings"
)

var (
	FlagNotFound            = errors.New("flag not found")
	MissingCtxForHelpRender = errors.New("missing ctx for help render")
)

type Ctx struct {
	flagMap        map[string]flag.Flag
	rootCommand    CommandInterface
	currentCommand CommandInterface
	values         []string
	commandChain   []CommandInterface
	globalCommands []CommandInterface
	globalFlags    []flag.Flag
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

func (c *Ctx) addToChain(cmd CommandInterface) {
	c.commandChain = append(c.commandChain, cmd)
}

func newCliCtx(
	flagMap map[string]flag.Flag,
	rootCommand CommandInterface,
	globalCommands []CommandInterface,
	globalFlags []flag.Flag,
) *Ctx {
	return &Ctx{
		flagMap:        flagMap,
		rootCommand:    rootCommand,
		currentCommand: rootCommand,
		globalCommands: globalCommands,
		globalFlags:    globalFlags,
	}
}

type HelpCtx struct {
	CommandChain              []string
	CurrentCommandName        string
	CurrentCommandDescription string
	RequiredFlags             []flag.Flag
	OptionalFlags             []flag.Flag
	AvailableCommands         map[string]string
}

type HelpRendererFunc func(*HelpCtx) string

type CommandHandlerFunc func(*Ctx) Result

type CommandInterface interface {
	Flags() []flag.Flag
	Commands() []CommandInterface
	CommandName() string
	HelpDescription() string
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

func (c *Command) HelpDescription() string {
	return c.Description
}

func (c *Command) Commands() []CommandInterface {
	return c.SubCommands
}

func (c *Command) HandlerFunc() CommandHandlerFunc {
	return c.Handler
}

type Result struct {
	value            interface{}
	errors           []error
	helpRenderer     HelpRendererFunc
	currentCtx       *Ctx
	shouldRenderHelp bool
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

func (r *Result) ShouldRenderHelp() bool {
	return r.shouldRenderHelp
}

func (r *Result) RenderHelp() string {
	if r.currentCtx == nil {
		panic(MissingCtxForHelpRender)
	}

	var requiredFlags []flag.Flag
	var optionalFlags []flag.Flag
	addedFlags := map[string]bool{}
	for _, f := range r.currentCtx.flagMap {
		if _, ok := addedFlags[f.NameList()[0]]; ok {
			continue
		}
		addedFlags[f.NameList()[0]] = true

		if f.IsRequired() {
			requiredFlags = append(requiredFlags, f)
		} else {
			optionalFlags = append(optionalFlags, f)
		}
	}

	var commandChain []string
	for _, c := range r.currentCtx.commandChain {
		commandChain = append(commandChain, c.CommandName())
	}

	availableCommands := map[string]string{}
	for _, c := range append(r.currentCtx.currentCommand.Commands(), r.currentCtx.globalCommands...) {
		availableCommands[c.CommandName()] = c.HelpDescription()
	}

	helpCtx := &HelpCtx{
		CommandChain:              commandChain,
		CurrentCommandName:        r.currentCtx.currentCommand.CommandName(),
		CurrentCommandDescription: r.currentCtx.currentCommand.HelpDescription(),
		RequiredFlags:             requiredFlags,
		OptionalFlags:             optionalFlags,
		AvailableCommands:         availableCommands,
	}

	return r.helpRenderer(helpCtx)
}

func HelpResult() Result {
	return Result{shouldRenderHelp: true}
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
	args              []string
	Name              string
	GlobalFlags       []flag.Flag
	GlobalCommands    []CommandInterface
	Description       string
	CommandSet        []CommandInterface
	FlagsSet          []flag.Flag
	Handler           CommandHandlerFunc
	HelpRenderer      HelpRendererFunc
	RenderHelpOnError bool
}

func (cli *CLI) CommandName() string {
	return cli.Name
}

func (cli *CLI) HelpDescription() string {
	return cli.Description
}

func (cli *CLI) getHelpHandler() HelpRendererFunc {
	if cli.HelpRenderer == nil {
		return defaultHelpHandler
	}
	return cli.HelpRenderer
}

func (cli *CLI) Run(args []string) Result {
	iter := arg.NewArgIterator(args)

	ctx := newCliCtx(nil, cli, cli.GlobalCommands, cli.GlobalFlags)
	errs := parse(ctx, iter)
	fmt.Printf("%v", errs)
	if errs != nil {
		return Result{
			errors:           errs,
			helpRenderer:     cli.getHelpHandler(),
			currentCtx:       ctx,
			shouldRenderHelp: cli.RenderHelpOnError,
		}
	}
	if isHelpSet(ctx) {
		return HelpResult()
	}

	h := ctx.CurrentCommand().HandlerFunc()
	if h == nil {
		return Result{
			errors:           []error{fmt.Errorf("handler for command '%s' does not exist", ctx.CurrentCommand().CommandName())},
			helpRenderer:     cli.getHelpHandler(),
			currentCtx:       ctx,
			shouldRenderHelp: cli.RenderHelpOnError,
		}
	}
	r := h(ctx)
	if r.errors != nil {
		r.shouldRenderHelp = cli.RenderHelpOnError
	}
	r.helpRenderer = cli.getHelpHandler()
	r.currentCtx = ctx
	return r
}

func (cli *CLI) Flags() []flag.Flag {
	return cli.FlagsSet
}

func (cli *CLI) Commands() []CommandInterface {
	return cli.CommandSet
}

func (cli *CLI) HandlerFunc() CommandHandlerFunc {
	return cli.Handler
}

func defaultHelpHandler(ctx *HelpCtx) string {
	builder := strings.Builder{}

	if 1 < len(ctx.CommandChain) {
		builder.WriteString(fmt.Sprintf("%s %s: <command>\n\n", ctx.CommandChain[0], strings.Join(ctx.CommandChain[1:], " ")))
	} else {
		builder.WriteString(fmt.Sprintf("%s:  <command>\n\n", ctx.CommandChain[0]))
	}

	if 0 < len(ctx.AvailableCommands) {
		builder.WriteString("Commands:\n")

		for c, d := range ctx.AvailableCommands {
			builder.WriteString(fmt.Sprintf("\t%s\t%s\n", c, d))
		}
		builder.WriteString("\n")
	}

	if 0 < len(ctx.RequiredFlags) {
		builder.WriteString("Required flags:\n")

		for _, f := range ctx.RequiredFlags {
			var names strings.Builder
			for i, n := range f.NameList() {
				if 0 < i {
					names.WriteString(", ")
				}
				var dash string
				if 1 < len(n) {
					dash = "--"
				} else {
					dash = "-"
				}
				names.WriteString(fmt.Sprintf("%s%s", dash, n))
			}
			builder.WriteString(fmt.Sprintf("\t%s\t%s\n", names.String(), f.HelpDescription()))
		}
		builder.WriteString("\n")
	}

	if 0 < len(ctx.RequiredFlags) {
		builder.WriteString("Optional flags:\n")

		for _, f := range ctx.OptionalFlags {
			var names strings.Builder
			for i, n := range f.NameList() {
				if 0 < i {
					names.WriteString(", ")
				}
				var dash string
				if 1 < len(n) {
					dash = "--"
				} else {
					dash = "-"
				}
				names.WriteString(fmt.Sprintf("%s%s", dash, n))
			}
			builder.WriteString(fmt.Sprintf("\t%s\t%s\n", names.String(), f.HelpDescription()))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

var HelpCommand = &Command{
	Name:        "help",
	Description: "Displays this help message",
}

func isHelpSet(ctx *Ctx) bool {
	if f, ok := ctx.flagMap[flag.HelpFlag.Names[0]]; ok {
		return f.IsSet()
	}
	return ctx.CurrentCommand().CommandName() == HelpCommand.Name
}
