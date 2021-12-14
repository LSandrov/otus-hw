package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "ine2dpoints", expected: "ineedpoints"},
		{input: "i4b3c2l1", expected: "iiiibbbccl"},
		{input: ".2-5.2", expected: "..-----.."},
		{input: "o_O", expected: "o_O"},
		{input: "世界世界世界世界", expected: "世界世界世界世界"},
		{input: "世2界4世5", expected: "世世界界界界世世世世世"},
		{input: "世2世界3世4世5", expected: "世世世界界界世世世世世世世世世"},
		{input: "◌́◌́", expected: "◌́◌́"},
		{input: "é2é", expected: "ééé"},
		{input: "Hello, world2", expected: "Hello, worldd"},
		{input: ",4.1", expected: ",,,,."},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
