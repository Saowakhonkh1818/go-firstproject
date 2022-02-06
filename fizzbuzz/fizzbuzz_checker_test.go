package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzzAnswer(t *testing.T) {
	tests := []int{15, 20, 25, 100}

	for _, n := range tests {
		expected := fizzbuzzAnswer(n)
		assert.Equal(t, expected, FizzBuzz(n))
	}
}
