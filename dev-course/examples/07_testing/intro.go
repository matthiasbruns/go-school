package main

import (
	"fmt"
	"testing"
)

//START_1 OMIT
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) { // HL1
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans) // HL1
	}
}

//END_1 OMIT

func TestIntMinTableDriven(t *testing.T) {
	//START_2 OMIT
	var tests = []struct { // HL2
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{-1, 0, -1},
	}

	for _, tt := range tests { // HL2
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) { // HL2
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
	//END_2 OMIT
}

//START_3 OMIT
func BenchmarkIntMin(b *testing.B) { // HL3
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

//END_3 OMIT
