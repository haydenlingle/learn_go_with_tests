package main

import "fmt"

const spanish = "Spanish"

const french = "French"
const englishGreetingPfx = "Hello, "
const spanishGreetingPfx = "Hola, "
const frenchGreetingPfx = "Bonjour, "

func Hello(name string, language string) string {
	if len(name) <= 0 {
		name = "World"
	}
	if language == spanish {
		return spanishGreetingPfx + name
	}
	if language == french {
		return frenchGreetingPfx + name
	}
	return englishGreetingPfx + name
}

func main() {
	fmt.Println(Hello("world", "spanish"))
}
