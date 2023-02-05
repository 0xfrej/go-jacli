package cli

import (
	"fmt"
	"os"
	"testing"

	"github.com/lai0n/go-jacli/pkg/test"
)

var (
	flagNames  = []string{"test"}
	testValue  = "testVal111_2"
	testValue2 = "testKekMan"
)

func TestStringFlag_Apply(t *testing.T) {
	flag := StringFlag{
		Names: flagNames,
	}
	ctx := newCtx(t, []string{})

	t.Run("it panics when there is no other argument", func(t *testing.T) {
		err := flag.apply(ctx)
		test.AssertErrSame(t, fmt.Errorf("flag '%s' requires a parameter", flag.Names[0]), err)
	})

	resetCtx(t, ctx, []string{"--kek"})
	t.Run("it panics when there next argument is not a string", func(t *testing.T) {
		err := flag.apply(ctx)
		test.AssertErrSame(t, fmt.Errorf("flag '%s' requires a parameter", flag.Names[0]), err)
	})

	resetCtx(t, ctx, []string{testValue, testValue2})
	t.Run("it applies and consumes next argument used as value", func(t *testing.T) {
		err := flag.apply(ctx)
		test.AssertNil(t, err)
		test.AssertEquals(t, testValue, flag.value)

		if val, ok := ctx.Args.Next(); !ok || val.String() != testValue2 {
			t.Error("failed to consume argument used up as value")
		}
	})

	flag.TakesFile = true
	notPath := "whatever"
	resetCtx(t, ctx, []string{notPath})
	t.Run("it sets", func(t *testing.T) {
		// should fail
		err := flag.apply(ctx)
		test.AssertErrSame(t, fmt.Errorf("path '%s' does not exist", notPath), err)

		// should pass
		file := createTestFile(t)
		resetCtx(t, ctx, []string{file.Name()})
		err = flag.apply(ctx)
		test.AssertNil(t, err)
		test.AssertEquals(t, file.Name(), flag.value)
		file.Close()
		os.Remove(file.Name())
	})
}

func TestStringFlag_IsSet(t *testing.T) {
	flag := StringFlag{}

	t.Run("it returns false on empty string", func(t *testing.T) {
		test.AssertFalse(t, flag.IsSet())
	})

	flag.value = "im set"
	t.Run("it returns true on empty string", func(t *testing.T) {
		test.AssertTrue(t, flag.IsSet())
	})
}

func TestStringFlag_IsRequired(t *testing.T) {
	flag := StringFlag{}

	t.Run("it returns false on non-required flag", func(t *testing.T) {
		test.AssertFalse(t, flag.IsRequired())
	})

	flag.Required = true
	t.Run("it returns true on required flag", func(t *testing.T) {
		test.AssertTrue(t, flag.IsRequired())
	})
}

func TestStringFlag_Value(t *testing.T) {
	flag := StringFlag{
		value: testValue,
	}

	t.Run("it returns the value", func(t *testing.T) {
		test.AssertEquals(t, testValue, flag.Value())
	})
}

func newCtx(t *testing.T, args []string) *Ctx {
	t.Helper()
	return &Ctx{
		Args: &ArgsIterator{
			Args:          args,
			iteratorIndex: -1,
		},
	}
}

func resetCtx(t *testing.T, ctx *Ctx, args []string) {
	t.Helper()
	ctx.Args.Args = args
	ctx.Args.iteratorIndex = -1
}

func createTestFile(t *testing.T) *os.File {
	t.Helper()
	file, err := os.CreateTemp("", "*")
	if err != nil {
		t.Fatal("failed to create tmp file")
	}
	return file
}

func removeTestFile(t *testing.T, f *os.File) {
	t.Helper()
	f.Close()
	os.Remove(f.Name())
}
