package glot

import "testing"

func TestNewPlot(t *testing.T) {
	persist := false
	debug := true
	_, err := NewPlot(0, persist, debug)
	if err == nil {
		t.Error("Expected error when making a 0 dimensional plot.")
	}
}
