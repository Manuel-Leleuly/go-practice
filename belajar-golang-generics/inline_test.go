package belajargolanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin[int](100, 100))
	assert.Equal(t, int64(100), FindMin[int64](int64(100), int64(200)))
	assert.Equal(t, 100.0, FindMin(100.0, 200.0))
}

func GetFirst[T []E, E any](data T) E {
	return data[0]
}

func TestGetFirst(t *testing.T) {
	names := []string{"Manuel", "Theodore", "Leleuly"}
	first := GetFirst[[]string, string](names)
	assert.Equal(t, "Manuel", first)
}
