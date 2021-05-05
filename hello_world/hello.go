package hello_world

import "fmt"

const HELLO_PREFIX = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return HELLO_PREFIX + name
}

func main() {
	fmt.Println(Hello("Paul"))
}
