package glot

import (
	"testing"
)

var funcTestArgs = PlotArgs{
	Debug:   true,
	Persist: true,
}

func TestAddFunc3d(t *testing.T) {
	dimensions := 3
	plot, _ := NewPlot(dimensions, funcTestArgs)
	fct := func(x, y float64) float64 { return x - y }
	groupName := "Straight Line"
	pointsY := []float64{1, 2, 3}
	pointsX := []float64{1, 2, 3, 4, 5}
	err := plot.AddFunc3d(groupName, Lines, pointsX, pointsY, fct)
	if err == nil {
		t.Error("TestAddFunc3d raises error when the size of X and Y arrays are not equal.")
	}
}
