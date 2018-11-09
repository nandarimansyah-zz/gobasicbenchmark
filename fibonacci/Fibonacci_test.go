package fibonacci_test

import (
	"testing"

	"github.com/nandarimansyah/gobasicbenchmark/fibonacci"
)

// func BenchmarkFib10(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		fibbonaci.Fib(10)
// 	}
// }

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Fib(i)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
