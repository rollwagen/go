package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name + "!"
}

//lint:ignore
func main() {
	fmt.Println(Hello(""))
}
