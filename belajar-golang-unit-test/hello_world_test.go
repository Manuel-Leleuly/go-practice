package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("Sebelum unit test")

	m.Run()

	// after
	fmt.Println("Setelah unit test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Manuel")
	if result != "Hello Manuel" {
		panic("Result is not 'Hello Manuel'")
	}
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Manuel")
	assert.Equal(t, "Hello Manuel", result, "should be 'Hello Manuel'")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on windows")
	}

	result := HelloWorld("Manuel")
	assert.Equal(t, "Hello Manuel", result, "should be 'Hello Manuel'")
}

func TestSubTest(t *testing.T) {
	t.Run("Manuel", func(t *testing.T) {
		result := HelloWorld("Manuel")
		require.Equal(t, "Hello Manuel", result, "Result must be 'Hello Manuel'")
	})

	t.Run("Leleuly", func(t *testing.T) {
		result := HelloWorld("Leleuly")
		require.Equal(t, "Hello Manuel", result, "Result must be 'Hello Manuel'")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
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
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Manuel")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Manuel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Manuel")
		}
	})

	b.Run("Leleuly", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Leleuly")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Manuel",
			request: "Manuel",
		},
		{
			name:    "Theodore",
			request: "Theodore",
		},
		{
			name:    "Leleuly",
			request: "Leleuly",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
