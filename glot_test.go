package glot

import (
	"testing"
)

var glotTestArgs = PlotArgs{
	Debug:   false,
	Persist: false,
}

func TestNewPlotZeroDimensions(t *testing.T) {
	_, err := NewPlot(0, glotTestArgs)
	if err == nil {
		t.Error("Expected error when making a 0 dimensional plot.")
	}
}

func TestNewPlotOneDimensions(t *testing.T) {
	p, err := NewPlot(1, glotTestArgs)
	p.AddPointGroup("First Group", 0,
		[]float64{pow(2, 1), pow(2, 2), pow(2, 3)})
	if err != nil {
		t.Error(err)
	}
}

func pow(n float64, p int) float64 {
	if p == 0 {
		return n
	}
	return n * (pow(n, p-1))
}
