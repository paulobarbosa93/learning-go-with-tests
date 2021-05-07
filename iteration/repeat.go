package iteration

func Repeat(character string, repetitionTimes int) string {
	var repetition string

	for i := 0; i < repetitionTimes; i++ {
		repetition += character
	}

	return repetition
}
