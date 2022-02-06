package fizzbuzz

import "strconv"

func FizzBuzz(a int) []string {
	var str []string
	for i := 1; i <= a; i++ {
		if i%3 == 0 && i%5 == 0 {
			str = append(str, "FizzBuzz")
		} else if i%3 == 0 {
			str = append(str, "Fizz")
		} else if i%5 == 0 {
			str = append(str, "Buzz")
		} else {
			str = append(str, strconv.Itoa(i))
		}
	}
	return str
}
