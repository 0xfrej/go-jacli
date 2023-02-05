package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Flag interface {
	NameList() []string
	IsSet() bool
	IsRequired() bool
	apply(*Ctx) error
}

type ValueFlag[T any] interface {
	Value() T
}

type StringFlag struct {
	Flag
	ValueFlag[string]

	Names     []string
	Required  bool
	value     string
	TakesFile bool
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

func (f *StringFlag) apply(ctx *Ctx) error {
	arg, ok := ctx.Args.Peek()
	if !ok || arg.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args.Next() // take next argument
	f.value = arg.String()

	if f.TakesFile {
		if _, err := os.Stat(f.value); os.IsNotExist(err) {
			return fmt.Errorf("path '%s' does not exist", f.value)
		} else if err != nil {
			return err
		}
	}

	return nil
}

type BoolFlag struct {
	Flag
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

func (f *BoolFlag) apply(_ *Ctx) error {
	f.value = true
	return nil
}

type IntFlag struct {
	Flag
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

func (f *IntFlag) apply(ctx *Ctx) error {
	arg, ok := ctx.Args.Peek()
	if !ok || arg.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args.Next() // take next argument
	v, e := strconv.Atoi(arg.String())
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid integer '%v'", f.Names[0], arg.String())
	}
	f.value = v
	f.wasSet = true
	return nil
}

type Float32Flag struct {
	Flag
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

func (f *Float32Flag) apply(ctx *Ctx) error {
	arg, ok := ctx.Args.Peek()
	if !ok || arg.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args.Next() // take next argument
	v, e := strconv.ParseFloat(arg.String(), 32)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid float '%v'", f.Names[0], arg.String())
	}
	f.value = float32(v)
	f.wasSet = true
	return nil
}

type Float64Flag struct {
	Flag
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

func (f *Float64Flag) apply(ctx *Ctx) error {
	arg, ok := ctx.Args.Peek()
	if !ok || arg.IsFlag() {
		return fmt.Errorf("flag '%s' requires a parameter", f.Names[0])
	}

	ctx.Args.Next() // take next argument
	v, e := strconv.ParseFloat(arg.String(), 64)
	if e != nil {
		return fmt.Errorf("flag '%s' contains invalid float '%v'", f.Names[0], arg.String())
	}
	f.value = v
	f.wasSet = true
	return nil
}

func parseFlags(flags []Flag, iterator *ArgsIterator) []error {
	flagMap, err := buildFlagMap(flags)
	if err != nil {
		return []error{err}
	}

	ctx := &Ctx{
		Args:  iterator,
		Flags: flagMap,
	}

	var errorSet []error
	applyFlag := func(flagName string) {
		if flag, ok := flagMap[flagName]; ok {
			errorSet = append(errorSet, flag.apply(ctx))
		}
	}

	for ctx.Args.HasNext() {
		if arg, ok := ctx.Args.Next(); ok && arg.IsFlag() {
			// when hyphen count equals one treat every character as a flag name.
			if arg.hyphenCount == 1 {
				for _, flagName := range arg.String() {
					applyFlag(string(flagName))
				}
			} else {
				applyFlag(arg.String())
			}
		}
	}

	return errorSet
}

func validateFlags(flagSet []Flag) []error {
	var errorSet []error

	for _, flag := range flagSet {
		if flag.IsRequired() && !flag.IsSet() {
			errorSet = append(errorSet, fmt.Errorf("flag '%s' is required", flag.NameList()[0]))
		}
	}

	return errorSet
}

func buildFlagMap(flags []Flag) (map[string]Flag, error) {
	flagMap := make(map[string]Flag)

	for _, flag := range flags {
		names := flag.NameList()
		if len(names) < 1 {
			return nil, errors.New("flag must have at least one name")
		}
		for _, name := range names {
			if _, hasFlag := flagMap[name]; hasFlag {
				return nil, fmt.Errorf("flag '%s' was already registered", name)
			}
			flagMap[name] = flag
		}
	}
	return flagMap, nil
}
