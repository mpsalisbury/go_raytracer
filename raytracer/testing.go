package raytracer

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

func approxOpts() []cmp.Option {
	return []cmp.Option{
		cmpopts.EquateApprox(0, 0.0001),
		cmp.AllowUnexported(Point{}, Vector{}, Matrix{}),
	}
}

func approxEq(x, y interface{}) bool {
	return cmp.Equal(x, y, approxOpts()...)
}

func approxError(got, want interface{}) string {
	diff := cmp.Diff(got, want, approxOpts()...)
	return fmt.Sprintf("got %v; want %v; diff %s", got, want, diff)
}

func assertPanic(t *testing.T, f func(), err string) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(err)
		}
	}()
	f()
}

func assertNoPanic(t *testing.T, f func(), err string) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf(err)
		}
	}()
	f()
}
