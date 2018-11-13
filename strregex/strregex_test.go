package strregex_test

import (
	"testing"

	"github.com/nandarimansyah/gobasicbenchmark/strregex"
)

func BenchmarkMatchString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strregex.IsMatchUsingMatchString("nanda@gmail.com")
	}
}

func BenchmarkMatchStringCompiled(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strregex.IsMatchUsingMatchStringCompiled("nanda@gmail.com")
	}
}
