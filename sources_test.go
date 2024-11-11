package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupping(t *testing.T) {
	input := []string{
		"abcde", "abcdef",
		"abcd2", "abdef",
		"abdfg", "abcd3",
		"bcdef", "123456",
		"1234567", "123455",
		"12345", "def",
		"123",
	}
	result := GroupStrings(input, 4)
	assert.NotNil(t, result)

	expected := map[string][]string{
		"abcd":  {"abcde", "abcdef", "abcd2", "abcd3"},
		"abdef": {"abdef"},
		"abdfg": {"abdfg"},
		"bcdef": {"bcdef"},
		"def":   {"def"},
		"12345": {"12345", "123456", "1234567", "123455"},
		"123":   {"123"},
	}
	assert.Equal(t, result, expected)
}
