package cli

import (
	"reflect"
	"testing"

	"github.com/lai0n/go-jacli/pkg/test"
)

var (
	argFlag = "--test"
	argVal  = "test"
)

func TestCommandArg_String(t *testing.T) {
	arg := CommandArg{
		value:       &argFlag,
		hyphenCount: 2,
	}

	t.Run("it returns arg flag without hyphens", func(t *testing.T) {
		test.AssertEquals(t, argFlag[2:], arg.String())
	})

	arg.hyphenCount = 0
	t.Run("it returns arg", func(t *testing.T) {
		test.AssertEquals(t, argFlag, arg.String())
	})
}

func TestCommandArg_IsFlag(t *testing.T) {
	arg := CommandArg{}

	t.Run("it returns false when argument does not contain any hyphens", func(t *testing.T) {
		test.AssertFalse(t, arg.IsFlag())
	})

	arg.hyphenCount = uint8(2)
	t.Run("it returns true when argument contains hyphens", func(t *testing.T) {
		test.AssertTrue(t, arg.IsFlag())
	})
}

func TestCommandArg_newArg(t *testing.T) {
	t.Run("it constructs argument that is not a flag", func(t *testing.T) {
		arg := newArg(&argVal)
		test.AssertFalse(t, arg.IsFlag())
	})

	t.Run("it constructs argument that is a flag with correct hyphen count", func(t *testing.T) {
		arg := newArg(&argFlag)
		test.AssertTrue(t, arg.IsFlag())
		test.AssertEquals(t, uint8(2), arg.hyphenCount)
	})
}

func TestArgsIterator_HasNext(t *testing.T) {
	iter := ArgsIterator{iteratorIndex: -1}
	t.Run("it returns false on empty array", func(t *testing.T) {
		test.AssertFalse(t, iter.HasNext())
	})

	iter.Args = make([]string, 2)
	t.Run("it returns true when on new iterator and filled array", func(t *testing.T) {
		test.AssertTrue(t, iter.HasNext())
	})

	iter.iteratorIndex = 0
	t.Run("it returns true when there is one element left", func(t *testing.T) {
		test.AssertTrue(t, iter.HasNext())
	})

	iter.iteratorIndex = 1
	t.Run("it returns false when there is no element left", func(t *testing.T) {
		test.AssertFalse(t, iter.HasNext())
	})
}

func TestArgsIterator_Peek(t *testing.T) {
	iter := ArgsIterator{iteratorIndex: -1}
	t.Run("it returns nok on empty array", func(t *testing.T) {
		_, ok := iter.Peek()
		test.AssertFalse(t, ok)
	})

	iter.Args = []string{"ahoy", "captain"}
	t.Run("it returns true when on new iterator and filled array and does not increment 'iteratorIndex'", func(t *testing.T) {
		oldIndex := iter.iteratorIndex
		val, ok := iter.Peek()
		test.AssertTrue(t, ok)
		test.AssertEquals(t, "ahoy", val.String())
		test.AssertEquals(t, oldIndex, iter.iteratorIndex)
	})

	iter.iteratorIndex = 0
	t.Run("it returns elem when there is one element left", func(t *testing.T) {
		val, ok := iter.Peek()
		test.AssertTrue(t, ok)
		test.AssertEquals(t, "captain", val.String())
	})

	iter.iteratorIndex = 1
	t.Run("it does not return elem when there is no element left", func(t *testing.T) {
		val, ok := iter.Peek()
		test.AssertFalse(t, ok)
		test.AssertEquals(t, nil, val)
	})
}

func TestArgsIterator_Next(t *testing.T) {
	iter := ArgsIterator{iteratorIndex: -1}
	t.Run("it returns nok on empty array", func(t *testing.T) {
		_, ok := iter.Next()
		test.AssertFalse(t, ok)
	})

	iter.Args = []string{"ahoy", "captain"}
	t.Run("it returns true when on new iterator and filled array and increments 'iteratorIndex'", func(t *testing.T) {
		oldIndex := iter.iteratorIndex
		val, ok := iter.Next()
		test.AssertTrue(t, ok)
		test.AssertEquals(t, "ahoy", val.String())
		test.AssertEquals(t, oldIndex+1, iter.iteratorIndex)
	})

	iter.iteratorIndex = 0
	t.Run("it returns elem when there is one element left", func(t *testing.T) {
		val, ok := iter.Next()
		test.AssertTrue(t, ok)
		test.AssertEquals(t, "captain", val.String())
	})

	iter.iteratorIndex = 1
	t.Run("it does not return elem when there is no element left", func(t *testing.T) {
		val, ok := iter.Next()
		test.AssertFalse(t, ok)
		test.AssertEquals(t, nil, val)
	})
}

func TestNewArgsIterator(t *testing.T) {
	args := []string{argFlag}

	iter := newArgsIterator(args)
	t.Run("it constructs iterator correctly", func(t *testing.T) {
		want := args
		if !reflect.DeepEqual(want, iter.Args) {
			t.Errorf("Expected '%v' got '%v'", want, iter.Args)
		}
		test.AssertEquals(t, -1, iter.iteratorIndex)
	})
}
