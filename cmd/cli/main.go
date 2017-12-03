package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/rcliao/advent-2017/day1"
)

func main() {
	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(dat), "\n")
	fmt.Println(input)
	fmt.Println(day1.CountWithNextDigitHalf(input))
}
