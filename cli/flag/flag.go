package flag

import (
	"errors"
	"flag"
	"fmt"
	"github.com/lai0n/go-jacli/cli/arg"
	"github.com/lai0n/go-jacli/pkg/iterator"
	"strconv"
)

var (
	FlagImpossibleToCast = errors.New("not a value flag")
)

type ParseCtx struct {
	args  iterator.Iterator[*arg.CommandArg]
	flags map[string]flag.Flag
}

func (c *ParseCtx) Args() iterator.Iterator[*arg.CommandArg] {
	return c.args
}

func (c *ParseCtx) Flags() map[string]flag.Flag {
	return c.flags
}

type Flag interface {
	NameList() []string
	IsSet() bool
	IsRequired() bool
	Apply(*ParseCtx) error
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

	Names    []string
	Required bool
	value    string
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

type BoolFlag struct {
	ValueFlag[bool]

	Names []string
	value bool
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

type IntFlag struct {
	ValueFlag[int]

	Names    []string
	Required bool
	value    int
	wasSet   bool
}

func (f *IntFlag) NameList() []string {
	return f.Names
}

func (f *IntFlag) Value() int {
	return f.value
}

func (f *IntFlag) IsSet() bool {
	return f.wasSet
}

func (f *IntFlag) IsRequired() bool {
	return f.Required
}

func (f *IntFlag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.Atoi(a.String())
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid integer '%v'", f.Names[0], a.String())
	}
	f.value = v
	f.wasSet = true
	return nil
}

type Float32Flag struct {
	ValueFlag[float32]

	Names    []string
	Required bool
	value    float32
	wasSet   bool
}

func (f *Float32Flag) NameList() []string {
	return f.Names
}

func (f *Float32Flag) Value() float32 {
	return f.value
}

func (f *Float32Flag) IsSet() bool {
	return f.wasSet
}

func (f *Float32Flag) IsRequired() bool {
	return f.Required
}

func (f *Float32Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseFloat(a.String(), 32)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid float '%v'", f.Names[0], a.String())
	}
	f.value = float32(v)
	f.wasSet = true
	return nil
}

type Float64Flag struct {
	ValueFlag[float64]

	Names    []string
	Required bool
	value    float64
	wasSet   bool
}

func (f *Float64Flag) NameList() []string {
	return f.Names
}

func (f *Float64Flag) Value() float64 {
	return f.value
}

func (f *Float64Flag) IsSet() bool {
	return f.wasSet
}

func (f *Float64Flag) IsRequired() bool {
	return f.Required
}

func (f *Float64Flag) Apply(ctx *ParseCtx) error {
	a, ok := ctx.Args().Peek()
	if !ok || a.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args().Next() // take next argument
	v, e := strconv.ParseFloat(a.String(), 64)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid float '%v'", f.Names[0], a.String())
	}
	f.value = v
	f.wasSet = true
	return nil
}
