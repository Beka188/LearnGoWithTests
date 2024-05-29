package arays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	var answer []int
	for _, numbers := range numbersToSum {
		answer = append(answer, Sum(numbers))
	}
	return answer
}

func SumAllTails(numbersToSum ...[]int) (answer []int) {
	for _, numbers := range numbersToSum {
		current := 0
		for i, number := range numbers {
			if i != 0 {
				current += number
			}
		}
		answer = append(answer, current)
	}
	return
}
