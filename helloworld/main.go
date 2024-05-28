package main

import (
	"fmt"
)

const (
	spanish = "Spanish"

	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}
func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// hello()
}
func Factorial(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 || x == 1 {
		return 1
	}
	return Factorial(x-1) * x
}
