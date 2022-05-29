package primes

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name           string
	input          int
	expectedOutput []int
	expectedError  error
}{
	{
		name:           "case for 5",
		input:          5,
		expectedOutput: []int{2, 3, 5},
	},
	{
		name:           "case for 10",
		input:          10,
		expectedOutput: []int{2, 3, 5, 7},
	},
	{
		name:           "case for 15",
		input:          15,
		expectedOutput: []int{2, 3, 5, 7, 11, 13},
	},
	{
		name:          "case for negative number",
		input:         -10,
		expectedError: ErrNotPositive,
	},
}

func TestEratosthenes(t *testing.T) {

	for _, cs := range testCases {
		cs := cs
		t.Run(cs.name, func(t *testing.T) {
			res, err := Eratosthenes(cs.input)
			if err != nil {
				if !errors.Is(err, cs.expectedError) {
					t.Fatalf("unexpected error: %s", err.Error())
				}
				return
			}
			if !SlicesEqual(res, cs.expectedOutput) {
				t.Fatalf("wrong result, got: %v, expected: %v", res, cs.expectedOutput)
			}
		})
	}
}

func SlicesEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func TestEratosthenesWithTestify(t *testing.T) {

	for _, cs := range testCases {
		cs := cs
		t.Run(cs.name, func(t *testing.T) {
			res, err := Eratosthenes(cs.input)
			if err != nil {
				if !errors.Is(err, cs.expectedError) {
					t.Fatalf("unexpected error: %s", err.Error())
				}
				return
			}
			assert.ElementsMatch(t, cs.expectedOutput, res)
		})
	}
}
