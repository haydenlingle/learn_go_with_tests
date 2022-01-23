package main

import "fmt"

const englishGreetingPfx = "Hello, "

func Hello(name string) string {
	if len(name) <= 0 {
		name = "World"
	}
	return englishGreetingPfx + name
}

func main() {
	fmt.Println(Hello("world"))
}
