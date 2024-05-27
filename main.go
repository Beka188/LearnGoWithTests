package main

import "fmt"

func Hello(name string) string {
	return "Hello " + name
}
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// hello()
}
