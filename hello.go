package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, lang string) string {

	prefix := englishHelloPrefix // default to english
	if lang == "Spanish" {
		prefix = spanishHelloPrefix
	}

	if name == "" && (lang == "English" || lang == "") {
		name = "World"
	}
	if name == "" && lang == "Spanish" {
		name = "el mundo"
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("", ""))
}
