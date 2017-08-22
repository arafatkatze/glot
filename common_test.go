package glot

import "testing"

func TestSetLabels(t *testing.T) {
	dimensions := 3
	persist := false
	debug := false
	plot, _ := NewPlot(dimensions, persist, debug)
	err := plot.SetLabels()
	if err == nil {
		t.Error("SetLabels raises error when empty string is passed")
	}
}

func TestSetFormat(t *testing.T) {
	dimensions := 3
	persist := false
	debug := false
	plot, _ := NewPlot(dimensions, persist, debug)
	err := plot.SetFormat("tls")
	if err == nil {
		t.Error("SetLabels raises error when non-supported format is passed as an argument.")
	}
}
