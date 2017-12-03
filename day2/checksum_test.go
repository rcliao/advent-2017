package day2

import "testing"

var tests1 = []struct {
	in  string
	out int
}{
	{"5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8", 18},
}
var tests2 = []struct {
	in  string
	out int
}{
	{"5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5", 9},
}

func TestCorruptionChecksum(t *testing.T) {
	for _, tt := range tests1 {
		e := CorruptionChecksum(tt.in)
		if e != tt.out {
			t.Errorf("CorruptionChecksum(%q) => %d, wants %d", tt.in, e, tt.out)
		}
	}
}

func TestCorruptionChecksum2(t *testing.T) {
	for _, tt := range tests2 {
		e := CorruptionChecksum2(tt.in)
		if e != tt.out {
			t.Errorf("CorruptionChecksum2(%q) => %d, wants %d", tt.in, e, tt.out)
		}
	}
}
