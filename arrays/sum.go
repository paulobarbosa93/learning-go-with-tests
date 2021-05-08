package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberToSum ...[]int) (sums []int) {
	for _, number := range numberToSum {
		sums = append(sums, Sum(number))
	}
	return
}

func SumAllRest(numberToSum ...[]int) (sums []int) {
	for _, number := range numberToSum {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			final := number[1:]
			sums = append(sums, Sum(final))
		}
	}
	return
}
