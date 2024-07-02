package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{{
		name:    "aant",
		request: "aant",
	}, {
		name:    "sultan",
		request: "sultan",
	}}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSubBenchmark(b *testing.B) {
	b.Run("aant", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("aant")
		}
	})

	b.Run("sultan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("sultan")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("aant")
	}
}

func BenchmarkHelloWorldElaina(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("elaina")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{{
		name:     "HelloWorld(aant)",
		request:  "aant",
		expected: "Hello World aant",
	}, {
		name:     "HelloWorld(sultan)",
		request:  "sultan",
		expected: "Hello World sultan",
	}, {
		name:     "HelloWorld(rahmanya)",
		request:  "rahmanya",
		expected: "Hello World rahmanya",
	}, {
		name:     "HelloWorld(elaina)",
		request:  "elaina",
		expected: "Hello World elaina",
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Sultan", func(t *testing.T) {
		result := HelloWorld("Sultan")
		require.Equal(t, "Hello World Sultan", result, "this should equal")
	})
	t.Run("Elaina", func(t *testing.T) {
		result := HelloWorld("Elaina")
		require.Equal(t, "Hello World Elaina", result, "this should equal")
	})
}

// before/after function test
func TestMain(m *testing.M) {
	// hanya jalan di satu package
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on Windows")
	}

	result := HelloWorld("aant")
	require.Equal(t, "Hello World aant", result, "this should equal")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("aant")
	assert.Equal(t, "Hello World aant", result, "this should equal")
	fmt.Println("test hello world with assertion is done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("aant")
	require.Equal(t, "Hello World aant", result, "this should equal")
	fmt.Println("test hello world with require is done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("aant")
	if result != "Hello World aant" {
		// unit test failed
		//panic("ERROR")
		// Fail() akan tetap nge-print fmt/function selanjutnya
		//t.Fail()

		t.Error("ERROR function")
	}

	fmt.Println("Test hello world done")
}

func TestHelloWorldSultan(t *testing.T) {
	result := HelloWorld("sultan")
	if result != "Hello World sultan" {
		// unit test failed
		// FailNow() akan berhenti disaat ini juga
		//t.FailNow()
		t.Fatal("FATAL function")
	}
	fmt.Println("Test hello world sultan done")
}
