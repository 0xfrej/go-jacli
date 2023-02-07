package cli

import (
	"errors"
	"fmt"
	"github.com/lai0n/go-jacli/cli/arg"
	"github.com/lai0n/go-jacli/cli/flag"
	"github.com/lai0n/go-jacli/pkg/iterator"
)

func parse(ctx *Ctx, iter iterator.Iterator[*arg.CommandArg]) []error {
	if a, ok := iter.Peek(); ok && !a.IsFlag() {
		if cmd := findCommandByName(a.String(), ctx.CurrentCommand().Commands()); cmd != nil {
			iter.Next()
			ctx.setCurrentCommand(cmd)
			parse(ctx, iter)
		}
	}

	flags := append(ctx.RootCommand().Flags(), ctx.CurrentCommand().Flags()...)
	errs := parseFlags(flags, ctx, iter)
	if errs != nil {
		return errs
	}

	errs = validateFlags(flags)
	if errs != nil {
		return errs
	}

	return nil
}

func findCommandByName(needle string, haystack []CommandInterface) CommandInterface {
	for _, v := range haystack {
		if v.CommandName() == needle {
			return v
		}
	}
	return nil
}

func parseFlags(flags []flag.Flag, runCtx *Ctx, iter iterator.Iterator[*arg.CommandArg]) []error {
	flagMap, err := buildFlagMap(flags)
	if err != nil {
		return []error{err}
	}
	runCtx.setFlags(flagMap)

	ctx := flag.NewParseCtx(iter, flagMap)

	var errorSet []error
	applyFlag := func(flagName string) {
		if f, ok := flagMap[flagName]; ok {
			errorSet = append(errorSet, f.Apply(ctx))
		}
	}

	var valueSet []string
	for ctx.Args().HasNext() {
		if a, ok := ctx.Args().Next(); ok {
			if a.IsFlag() {
				// when hyphen count equals one treat every character as a flag name.
				if a.HyphenCount() == 1 {
					for _, flagName := range a.String() {
						applyFlag(string(flagName))
					}
				} else {
					applyFlag(a.String())
				}
			} else {
				valueSet = append(valueSet, a.String())
			}
		}
	}
	runCtx.setValues(valueSet)

	return errorSet
}

func validateFlags(flagSet []flag.Flag) []error {
	var errorSet []error

	for _, f := range flagSet {
		if f.IsRequired() && !f.IsSet() {
			errorSet = append(errorSet, fmt.Errorf("flag '%s' is required", f.NameList()[0]))
		}
	}

	return errorSet
}

func buildFlagMap(flags []flag.Flag) (map[string]flag.Flag, error) {
	flagMap := make(map[string]flag.Flag)

	for _, f := range flags {
		names := f.NameList()
		if len(names) < 1 {
			return nil, errors.New("flag must have at least one name")
		}
		for _, name := range names {
			if _, hasFlag := flagMap[name]; hasFlag {
				return nil, fmt.Errorf("flag '%s' was already registered", name)
			}
			flagMap[name] = f
		}
	}
	return flagMap, nil
}
