package belajargolanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First, Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return first
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Manuel",
		Second: "Leleuly",
	}
	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Manuel",
		Second: "Leleuly",
	}

	assert.Equal(t, "Eko", data.ChangeFirst("Eko"))
	assert.Equal(t, "Hello Manuel", data.SayHello("Manuel"))
}
