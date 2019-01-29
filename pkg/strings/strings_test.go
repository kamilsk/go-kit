package strings_test

import (
	"strings"
	"testing"

	. "github.com/kamilsk/go-kit/pkg/strings"
	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		strings  []string
		expected string
	}{
		{name: "nothing to pass"},
		{"simple usage", []string{"127.0.0.1", ":", "80"}, "127.0.0.1:80"},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Concat(tc.strings...))
		})
	}
}

func TestFirst(t *testing.T) {
	tests := []struct {
		name     string
		strings  []string
		expected string
	}{
		{name: "nothing to pass"},
		{"simple usage", []string{"", "", "third"}, "third"},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, First(tc.strings...))
		})
	}
}

func BenchmarkConcat(b *testing.B) {
	b.Run("concat", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			Concat("127.0.0.1", ":", "80")
		}
	})
	b.Run("join", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			strings.Join([]string{"127.0.0.1", ":", "80"}, "")
		}
	})
}
