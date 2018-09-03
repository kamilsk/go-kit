package fn_test

import (
	"errors"
	"testing"

	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	testCases := []struct {
		name     string
		actions  []func() error
		panicked bool
	}{
		{
			name: "with panic",
			actions: []func() error{
				func() error { return nil },
				func() error { return errors.New("raise panic") },
				func() error { return nil },
			},
			panicked: true,
		},
		{
			name: "without panic",
			actions: []func() error{
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
			},
			panicked: false,
		},
	}

	for _, test := range testCases {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			var check = assert.NotPanics
			if tc.panicked {
				check = assert.Panics
			}
			check(t, func() { fn.Must(tc.actions...) })
		})
	}
}
