package day1

import "testing"

var digitTests1 = []struct {
	in  string
	out int
}{
	{"1122", 3},
	{"1111", 4},
	{"1234", 0},
	{"91212129", 9},
}
var digitTests2 = []struct {
	in  string
	out int
}{
	{"1212", 6},
	{"1221", 0},
	{"123425", 4},
	{"123123", 12},
	{"12131415", 4},
}

func TestCountWithNextDigit1(t *testing.T) {
	for _, tt := range digitTests1 {
		d := CountWithNextDigit(tt.in)
		if tt.out != d {
			t.Errorf("CountWithNextDigit(%q) => %d, want %d", tt.in, d, tt.out)
		}
	}
}

func TestCountWithNextDigit2(t *testing.T) {
	for _, tt := range digitTests2 {
		d := CountWithNextDigitHalf(tt.in)
		if tt.out != d {
			t.Errorf("CountWithNextDigit(%q) => %d, want %d", tt.in, d, tt.out)
		}
	}
}
