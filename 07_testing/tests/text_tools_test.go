package main

import (
	"golang-study/07_testing/text_tools"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFooFunc(t *testing.T) {
	assert.Equal(t, "foobar", text_tools.Concat("foo", "bar"))
}

func TestFetchUrlFunc(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"aaa github.com/demidovich aaa", "github.com/demidovich"},
		{"github.com/Demidovich?test=1 aaa", "github.com/Demidovich?test=1"},
		{"aaa github.com/Demidovich?test=1", "github.com/Demidovich?test=1"},
	}

	for _, tt := range tests {
		actual := text_tools.FetchUrl(tt.input)
		assert.Equal(t, tt.expected, actual[0])
	}
}
