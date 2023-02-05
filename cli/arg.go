package cli

import (
	"fmt"
	"github.com/lai0n/go-jacli/pkg/iterator"
)

type Ctx struct {
	Args  *ArgsIterator
	Flags map[string]Flag
}

type CommandArg struct {
	fmt.Stringer

	value       *string
	hyphenCount uint8
}

func (a *CommandArg) String() string {
	return (*(a.value))[a.hyphenCount:]
}

func (a *CommandArg) IsFlag() bool {
	return a.hyphenCount > 0
}

func newArg(str *string) *CommandArg {
	hyphenCount := uint8(0)
	for _, char := range *str {
		if char == '-' {
			hyphenCount += 1
			continue
		}
		break
	}

	return &CommandArg{
		value:       str,
		hyphenCount: hyphenCount,
	}
}

// ArgsIterator iterates an array of string and producing CommandArg as a result.
type ArgsIterator struct {
	iterator.Iterator[CommandArg]

	Args []string
	// Constructor should set value to -1 indicating that iteration did not start
	iteratorIndex int
}

func (a *ArgsIterator) HasNext() bool {
	l := len(a.Args)
	return l > 0 && a.iteratorIndex+1 < l
}

func (a *ArgsIterator) Next() (elem *CommandArg, ok bool) {
	if a.HasNext() {
		a.iteratorIndex += 1
		return newArg(&a.Args[a.iteratorIndex]), true
	}
	return nil, false
}

func (a *ArgsIterator) Peek() (elem *CommandArg, ok bool) {
	if a.HasNext() {
		return newArg(&a.Args[a.iteratorIndex+1]), true
	}
	return nil, false
}

func newArgsIterator(argSet []string) *ArgsIterator {
	return &ArgsIterator{
		Args:          argSet,
		iteratorIndex: -1,
	}
}
