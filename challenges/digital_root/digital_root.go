package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DigitalRoot(n int) (sum int) {
	nStr := strconv.Itoa(n)
	length := len(nStr)
	chars := strings.Split(nStr, "")

	if length > 1 {
		for _, char := range chars {
			tmp, _ := strconv.Atoi(char)
			sum += tmp
		}
		return
	} else {
		return n
	}
}

func main() {
	fmt.Println(DigitalRoot(123))
}
