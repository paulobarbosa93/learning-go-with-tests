package main

import "fmt"

const ESPANHOL = "espanhol"
const FRANCES = "frances"
const HELLO_PREFIX_EN = "Hello, "
const HELLO_PREFIX_ES = "Hola, "
const HELLO_PREFIX_fn = "Bonjour, "

// public function
func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return salutionPrefix(lang) + name
}

// private function
func salutionPrefix(lang string) (prefix string) {
	switch lang {
	case ESPANHOL:
		prefix = HELLO_PREFIX_ES
	case FRANCES:
		prefix = HELLO_PREFIX_fn
	default:
		prefix = HELLO_PREFIX_EN
	}

	return
}

func main() {
	fmt.Println(Hello("Paul", "espanhol"))
}
