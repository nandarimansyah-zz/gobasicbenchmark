package numericconversion_test

import (
	"testing"

	"github.com/nandarimansyah/gobasicbenchmark/numericconversion"
)

func BenchmarkParseBool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := numericconversion.MParseBool("true")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkParseInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := numericconversion.MParseInt("7182818284")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkParseFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := numericconversion.MParseFloat("3.1415926535")
		if err != nil {
			panic(err)
		}
	}
}
