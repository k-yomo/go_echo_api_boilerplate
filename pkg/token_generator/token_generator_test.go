package token_generator

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGenerateRandomNumStr(t *testing.T) {
	testCases := []struct {
		input    uint
		expected int
	}{
		{1, 1},
		{6, 6},
		{11, 11},
	}

	for _, tc := range testCases {
		result := GenerateRandomNumStr(tc.input)
		_, err := strconv.Atoi(result)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, len(result))
	}
}

func TestGenerateRandomStr(t *testing.T) {
	testCases := []struct {
		input    uint
		expected int
	}{
		{0, 0},
		{1, 1},
		{6, 6},
	}

	for _, tc := range testCases {
		result := GenerateRandomBase64Token(tc.input)
		assert.Equal(t, tc.expected, len(result))
	}
}
