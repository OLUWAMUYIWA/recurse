package main

import (
	"testing"
)

func testCountPartitions(t *testing.T) {
	if ans := countPartitions(5, 4); ans != 6 {
		t.Fatalf("FAIL! %d expected; %d gotten", 6, ans)
	}
	t.Logf("PASS!\n")
}
