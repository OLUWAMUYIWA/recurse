package main

import (
	"testing"
)

func testSumNumbers(t *testing.T) {
	if ans := sum_numbers(5); ans != 15 {
		t.Fatalf("FAIL: %d expected; %d goten", 5, ans)
	}
	t.Logf("PASS!\n")
}
