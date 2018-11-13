package stringconcat_test

import (
	"testing"

	"github.com/nandarimansyah/gobasicbenchmark/stringconcat"
)

const (
	TEST_STRING = "test"
	TEST_SIZE   = 2
)

func benchmarkConcat(size int, SelfConcat func(string, int) string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SelfConcat(TEST_STRING, size)
	}
}

func BenchmarkConcatOperator(b *testing.B) {
	benchmarkConcat(TEST_SIZE, stringconcat.SelfConcatOperator, b)
}

func BenchmarkConcatBuffer(b *testing.B) { benchmarkConcat(TEST_SIZE, stringconcat.SelfConcatBuffer, b) }
