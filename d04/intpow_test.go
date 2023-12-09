package d04

import "testing"

func TestIntPow(t *testing.T) {
	testcases := [][]int{
		{2, 1, 2},
		{3, 0, 1},
		{2, 4, 16},
		{2, 11, 2048},
		{3, 3, 27},
	}
	for i, testcase := range testcases {
		got := intPow(testcase[0], testcase[1])
		if got != testcase[2] {
			t.Errorf("Testcase %d failed, got %d, want %d", i, got, testcase[2])
		}
	}
}