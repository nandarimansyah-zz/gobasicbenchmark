package main

import (
	"fmt"

	"github.com/nandarimansyah/gobasicbenchmark/stringconcat"

	"github.com/nandarimansyah/gobasicbenchmark/fibonacci"
)

func main() {
	fmt.Println("Hello World!!!")

	fmt.Println(fibonacci.Fib(40))

	text := "text"
	fmt.Println(stringconcat.SelfConcatOperator(text, 10))

	text2 := "text2"
	fmt.Println(stringconcat.SelfConcatBuffer(text2, 10))
}
