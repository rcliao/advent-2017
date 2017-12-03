package day2

import (
	"strconv"
	"strings"
)

// CorruptionChecksum does the checksum based on description at http://adventofcode.com/2017/day/2
func CorruptionChecksum(input string) int {
	lines := strings.Split(input, "\n")
	count := 0

	for _, line := range lines {
		parts := strings.Split(line, "\t")
		lineInts := parseStringsToInts(parts)
		count += getMaxDiffInLine(lineInts)
	}

	return count
}

// CorruptionChecksum2 does the checksum part 2 based on description at http://adventofcode.com/2017/day/2
func CorruptionChecksum2(input string) int {
	lines := strings.Split(input, "\n")
	count := 0

	for _, line := range lines {
		parts := strings.Split(line, "\t")
		lineInts := parseStringsToInts(parts)
		count += getDivisibleInLine(lineInts)
	}

	return count

}

func parseStringsToInts(list []string) []int {
	ls := make([]int, len(list))
	for i, s := range list {
		number, _ := strconv.Atoi(s)
		ls[i] = number
	}
	return ls
}

func getMaxDiffInLine(list []int) int {
	if len(list) == 0 {
		return 0
	}
	min := list[0]
	max := list[0]
	for _, i := range list {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}
	return max - min
}

func getDivisibleInLine(list []int) int {
	if len(list) == 0 {
		return 0
	}
	divisible := 0
	for i := range list {
		for j := range list {
			if i != j && list[i]%list[j] == 0 {
				divisible = list[i] / list[j]
			}
		}
	}
	return divisible
}
