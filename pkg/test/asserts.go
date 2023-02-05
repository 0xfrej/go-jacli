package test

import (
	"testing"
)

func AssertErrSame(t *testing.T, want error, got error) {
	t.Helper()
	if want.Error() != got.Error() {
		t.Errorf("Expected '%s' got '%s'", want, got)
	}
}

func AssertEquals[T comparable](t *testing.T, want T, got T) {
	t.Helper()
	if got != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}
}

func AssertNil(t *testing.T, got any) {
	t.Helper()
	if got != nil {
		t.Errorf("Expected 'nil' got '%v'", got)
	}
}

func AssertNotNil(t *testing.T, got any) {
	t.Helper()
	if got == nil {
		t.Errorf("Expected not nil got '%v'", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	AssertEquals(t, true, got)
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	AssertEquals(t, false, got)
}

func AssertPanic(t *testing.T, want any, f func()) {
	t.Helper()
	defer func() {
		if got := recover(); got == nil {
			t.Errorf("The code did not panic")
		} else if got != want {
			t.Errorf("Expected '%v' got '%v'", want, got)
		}
	}()
	f()
}
