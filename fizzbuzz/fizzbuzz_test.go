package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzz(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{
			name:     "3",
			n:        3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			name:     "5",
			n:        5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name:     "15",
			n:        15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			name:     "20",
			n:        20,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, FizzBuzz(test.n))
		})
	}
}
