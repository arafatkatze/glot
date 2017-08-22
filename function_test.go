package glot

import (
	"testing"
)

func TestAddFunc3d(t *testing.T) {
	dimensions := 3
	persist := false
	debug := false
	plot, _ := NewPlot(dimensions, persist, debug)
	fct := func(x, y float64) float64 { return x - y }
	groupName := "Stright Line"
	style := "lines"
	pointsY := []float64{1, 2, 3}
	pointsX := []float64{1, 2, 3, 4, 5}
	err := plot.AddFunc3d(groupName, style, pointsX, pointsY, fct)
	if err == nil {
		t.Error("TestAddFunc3d raises error when the size of X and Y arrays are not equal.")
	}
}
