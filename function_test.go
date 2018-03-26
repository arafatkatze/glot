package glot

import (
	"math"
	"math/rand"
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

func TestAddFunc2d1(t *testing.T) {
	dimensions := 2
	plot, _ := NewPlot(dimensions, funcTestArgs)
	fct := func(x float64) float64 { return math.Pow(math.E, x) }

	groupName := "Natural Exponential"
	pointsX := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	err := plot.AddFunc2d(groupName, Lines, pointsX, fct)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestAddFunc2d2(t *testing.T) {
	dimensions := 2
	plot, _ := NewPlot(dimensions, funcTestArgs)
	fct := func(x float64) float64 { return x * x }
	rand.Seed(111)
	groupName := "Exponential Floating Point Squared"
	pointsX := make([]float64, 1024)
	expFloat64(pointsX)
	err := plot.AddFunc2d(groupName, Points, pointsX, fct)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestAddFunc2d3(t *testing.T) {
	dimensions := 2
	plot, _ := NewPlot(dimensions, funcTestArgs)
	fct := func(x float64) float64 { return 1 / x }

	groupName := "Squared Integers"
	pointsX := make([]float64, 100)
	recSquared(pointsX)
	err := plot.AddFunc2d(groupName, Lines, pointsX, fct)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestAddFunc3d1(t *testing.T) {
	dimensions := 3
	plot, _ := NewPlot(dimensions, funcTestArgs)
	fct := func(x, y float64) float64 { return x - y }
	groupName := "Straight Line"
	pointsY := []float64{1, 2, 3, 4, 5}
	pointsX := []float64{1, 2, 3, 4, 5}
	err := plot.AddFunc3d(groupName, Lines, pointsX, pointsY, fct)
	if err != nil {
		t.Error(err.Error())
	}
}

func expFloat64(n []float64) []float64 {
	if len(n) == 0 {
		return n
	}
	n[0] = rand.ExpFloat64()
	return expFloat64(n[1:])
}
func recSquared(n []float64) float64 {
	if len(n) == 0 {
		return 0
	}
	n[len(n)-1] = float64((len(n) - 1) * (len(n) - 1))
	return recSquared(n[:len(n)-1])
}
