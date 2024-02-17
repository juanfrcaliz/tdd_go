package main

import "fmt"

const englishHelloPrfix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrfix + name
}

func main() {
	fmt.Println(Hello("Kiko"))
}
