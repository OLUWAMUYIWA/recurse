package main

import (
	"testing"
)

func testSumPatsBtwGrid(t *testing.T) {
	if ans := sumPathsBtwGrid(3, 3); ans != 6 {
		t.Fatalf("FAIL! %d expected; %d gotten", 6, ans)
	}
	t.Logf("PASS!\n")
}
