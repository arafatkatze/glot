package glot

import (
	"math"
	"testing"
)

func TestMin(t *testing.T) {
	var v int
	v = min(1, 2)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestCmd(t *testing.T) {
	base2n := func(n float64) []float64 {
		res := []float64{}
		for i := 0.; i < n; i++ {
			res = append(res, math.Pow(2, i))
		}
		return res
	}
	var testArgs = PlotArgs{
		Debug:   true,
		Format:  Pdf,
		Persist: true,
		Style:   Points,
		Command: "plot ",
	}
	var pg = &PointGroup{
		name:       "TestCmd",
		dimensions: 1,
		style:      Points,
		castedData: base2n(9),
	}
	p, err := NewPlot(1, testArgs)
	if err != nil {
		t.Errorf("NewPlot failed creation with: %v", err)
	}
	p.plotX(pg)
}
