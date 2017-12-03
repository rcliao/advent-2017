package day1

import (
	"fmt"
	"strconv"
)

// CountWithNextDigit counts total of those digit that is the same as the next digit
func CountWithNextDigit(input string) int {
	return count(input, 1)
}

// CountWithNextDigitHalf is nearly the same as the CountWithNextDigit except input
// is now a half circular list
func CountWithNextDigitHalf(input string) int {
	return count(input, len(input)/2)
}

func count(input string, step int) int {
	count := 0
	for pos, char := range input {
		digit, _ := strconv.Atoi(fmt.Sprintf("%c", char))
		nextPos := pos + step
		if nextPos >= len(input)-1 {
			nextPos = nextPos % len(input)
		}
		nextDigit, _ := strconv.Atoi(fmt.Sprintf("%c", input[nextPos]))
		if nextDigit == digit {
			count += digit
		}
	}
	return count
}
