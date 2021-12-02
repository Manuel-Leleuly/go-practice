package test

import (
	"fmt"
	"programmerzamannow/belajar-golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	fmt.Println(err)
	fmt.Println(simpleService)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}