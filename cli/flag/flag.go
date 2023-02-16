//go:generate go run ../../internal/generate/flags.go

package flag

import (
	"errors"
	"fmt"
	"github.com/lai0n/go-jacli/cli/arg"
	"github.com/lai0n/go-jacli/pkg/iterator"
)

var (
	FlagImpossibleToCast = errors.New("not a value flag")
)

var (
	HelpFlag = &BoolFlag{
		Names:       []string{"help", "h"},
		Description: "Display help page",
	}
)

type ParseCtx struct {
	args  iterator.Iterator[*arg.CommandArg]
	flags map[string]Flag
}

func (c *ParseCtx) Args() iterator.Iterator[*arg.CommandArg] {
	return c.args
}

func (c *ParseCtx) Flags() map[string]Flag {
	return c.flags
}

func NewParseCtx(args iterator.Iterator[*arg.CommandArg], flags map[string]Flag) *ParseCtx {
	return &ParseCtx{
		args:  args,
		flags: flags,
	}
}

type Flag interface {
	NameList() []string
	IsSet() bool
	IsRequired() bool
	Apply(*ParseCtx) error
	HelpDescription() string
}

type ValueFlag[T any] interface {
	Flag
	Value() T
}

func AsValueFlag[T any](flag Flag) (ValueFlag[T], error) {
	if f, ok := flag.(ValueFlag[T]); ok {
		return f, nil
	}
	return nil, FlagImpossibleToCast
}

func AsFlag[T Flag](flag Flag) (T, error) {
	if f, ok := flag.(T); ok {
		return f, nil
	}
	var r T
	return r, FlagImpossibleToCast
}

type StringFlag struct {
	ValueFlag[string]

	Names       []string
	Required    bool
	Description string
	value       string
}

func (f *StringFlag) NameList() []string {
	return f.Names
}

func (f *StringFlag) Value() string {
	return f.value
}

func (f *StringFlag) IsSet() bool {
	return f.value != ""
}

func (f *StringFlag) IsRequired() bool {
	return f.Required
}

func (f *StringFlag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	f.value = a.String()

	return nil
}

func (f *StringFlag) HelpDescription() string {
	return f.Description
}

type BoolFlag struct {
	ValueFlag[bool]

	Names       []string
	Description string
	value       bool
}

func (f *BoolFlag) NameList() []string {
	return f.Names
}

func (f *BoolFlag) Value() bool {
	return f.value
}

func (f *BoolFlag) IsSet() bool {
	return true
}

func (f *BoolFlag) IsRequired() bool {
	return false
}

func (f *BoolFlag) Apply(_ *ParseCtx) error {
	f.value = true
	return nil
}

func (f *BoolFlag) HelpDescription() string {
	return f.Description
}
