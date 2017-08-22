package glot

import "testing"

func TestResetPointGroupStyle(t *testing.T) {
	dimensions := 2
	persist := false
	debug := false
	plot, _ := NewPlot(dimensions, persist, debug)
	plot.AddPointGroup("Sample1", "points", []int32{51, 8, 4, 11})
	err := plot.ResetPointGroupStyle("Sam", "lines")
	if err == nil {
		t.Error("The specified pointgroup to be reset does not exist")
	}
}
