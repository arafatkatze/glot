package glot

import "testing"

func TestMin(t *testing.T) {
	var v int
	v = min(1, 2)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}
