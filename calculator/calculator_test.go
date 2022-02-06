package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 3, Sum(1, 2))
}

func TestMinus(t *testing.T) {
	assert.Equal(t, 0, Minus(1, 1))

}
