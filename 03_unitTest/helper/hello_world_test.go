package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
	Instead of using 'panic' as the way to tell the user that there's an error,
	we can instead use Fail() or FailNow().

	The difference between the two is that the first one will show the failure and continues the testing to the next line of code.
	Whereas the second one will show the failure and stop the testing process immediately.

	Instead of using Fail() or FailNow(), we can instead use Error() or Fatal().
	Error method will take an argument (for example string), displays the argument, and calls Fail().
	Fatal method will take an argument (for example string), displays the argument, and calls FailNow().

	To prevent creating custom conditional, we can use 'Testify' package created by stretchr to use assertion
*/

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Manuel")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with require done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Manuel")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Manuel'")
	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWordAssert(t *testing.T) {
	result := HelloWorld("Manuel")
	assert.Equal(t, "Hello Eko", result, "")

}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Manuel")
	if result != "Hello Manuel" {
		// error
		// panic("Result is not 'Hello Manuel'")
		// t.Fail()
		t.Error("Result must be 'Hello Manuel'")
	}
	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldTheodore(t *testing.T) {
	result := HelloWorld("Theodore")
	if result != "Hello Theodore" {
		// error
		// panic("Result is not 'Hello Theodore'")
		// t.FailNow()
		t.Fatal("Result must be 'Hello Theodore'")
	}
	fmt.Println("TestHelloWorldTheodore done")
}
