package errgroup

import (
	"errors"
	"testing"
)

func TestErrGroup(t *testing.T) {
	err1 := errors.New("some error")

	ts := []struct {
		err error
	}{
		{err: nil},
		{err: nil},
		{err: err1},
		{err: nil},
	}

	g := New()

	for _, tc := range ts {
		err := tc.err
		g.Go(func() error { return err })
	}

	if gotErr := g.Wait(); gotErr != err1 {
		t.Errorf("errgroup should got error: %v, but not", err1)
	}
}
