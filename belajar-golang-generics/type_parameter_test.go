package belajargolanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = Length("Manuel")
	assert.Equal(t, "Manuel", result)

	var resultNumber int = Length(100)
	assert.Equal(t, 100, resultNumber)
}
