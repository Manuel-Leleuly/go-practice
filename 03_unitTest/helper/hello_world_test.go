package helper

import (
	"fmt"
	"runtime"
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

/*
	a skip function (t.Skip()) is used to skip a testing on a certain method.
	For example, if the testing function is not compatible with a certain OS, you can use skip function to skip the testing and move on to the next
*/
func TestSkip(t *testing.T) {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run on linux")
	}

	result := HelloWorld("Manuel")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

/*
	test main is used to execute commands or functions before and/or after the unit test.
	For instance, if you want to test the functions of communicating with a database,
	you can use test main to first initialize and connect to the database before the testing begins.
	Afterwards, you can create a function or a command to disconnect the project to the database.

	WARNING: test main will only work in a package that it is written.
	If you write a test main inside package A, it will not run on other packages besides A
*/
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

/*
	The Run() method is used to create a subtest inside a test
*/
func TestSubTest(t *testing.T) {
	t.Run("Manuel", func(t *testing.T) {
		result := HelloWorld("Manuel")
		require.Equal(t, "Hello Manuel", result, "Result must be 'Hello Manuel'")
	})
	t.Run("Theodore", func(t *testing.T) {
		result := HelloWorld("Theodore")
		require.Equal(t, "Hello Theodore", result, "Result must be 'Hello Theodore'")
	})
}

/*
	Sometimes creating a similar test function over and over but with a different parameter can be quite exhausting.
	Therefore, we can use test table to create slice of struct and fill it with the data.
	Afterwards, we can iterate each data and run the test. Therefore, we can only write the testing function once.
*/
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "Manuel",
			request:  "Manuel",
			expected: "Hello Manuel",
		},
		{
			name:     "Theodore",
			request:  "Theodore",
			expected: "Hello Theodore",
		},
		{
			name:     "Leleuly",
			request:  "Leleuly",
			expected: "Hello Leleuly",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be "+"'"+test.expected+"'")
		})
	}
}

/*
	benchmark is used to test the performance of our code

	to run benchmark = go test -v -bench=.
	to run benchmark only = go test -v -run=asdf -bench=. // set -run to a non-existence unit test
	to run selected benchmark = go test -v -bench=<the name of the benchmark function>
*/

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Manuel")
	}
}

/*
	Similar to unit test, benchmark also has Run() method so that we can create sub benchmarks
*/

func BenchmarkSub(b *testing.B) {
	b.Run("Manuel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Manuel")
		}
	})
	b.Run("Theodore", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Theodore")
		}
	})
}
