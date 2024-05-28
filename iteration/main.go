package iteration

func Repeat(char string, count int) (answer string) {
	for i := 0; i < count; i++ {
		answer += char
	}
	return
}
