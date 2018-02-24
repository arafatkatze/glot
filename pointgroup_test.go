package glot

import (
	"testing"
)

func squareInt(n int) int {
	return n * n
}
func cubeInt(n int) int {
	return n * squareInt(n)
}

func squareInt16(n int16) int16 {
	return n * n
}
func cubeInt16(n int16) int16 {
	return n * squareInt16(n)
}

func TestResetPointGroupStyle(t *testing.T) {
	args := PlotArgs{
		Debug:   false,
		Persist: false,
	}
	dimensions := 2
	plot, _ := NewPlot(dimensions, args)
	plot.AddPointGroup("Sample1", Points, []int32{51, 8, 4, 11})
	err := plot.ResetPointGroupStyle("Sam", Lines)
	if err == nil {
		t.Error("The specified pointgroup to be reset does not exist")
	}
}

func TestTwoPointGroups(t *testing.T) {
	args := PlotArgs{
		Debug:   true,
		Persist: true,
		Format:  Pdf,
		Style:   Points,
	}

	plot, _ := NewPlot(1, args)
	plot.AddPointGroup("TestGroup_1", Points, []float64{
		-0.512695,
		0.591778,
		-0.0939544,
		-0.510766,
		-0.859442,
		0.0340482,
		0.887461,
		0.277168,
		-0.998753,
		0.356656,
	})
	plot.AddPointGroup("TestGroup_2", Points, []float64{
		0.712863,
		0.975935,
		0.875864,
		0.737082,
		-0.185717,
		-0.936551,
		0.779397,
		0.916793,
		0.622004,
		-0.0860084,
	})
}

func TestThreePointGroupsFloat64(t *testing.T) {
	args := PlotArgs{
		Debug:   true,
		Persist: true,
		Format:  Pdf,
		Style:   Points,
	}

	plot, _ := NewPlot(1, args)
	plot.AddPointGroup("TestGroup_1", Points, []float64{
		-0.512695,
		0.591778,
		-0.0939544,
		-0.510766,
		-0.859442,
		0.0340482,
		0.887461,
		0.277168,
		-0.998753,
		0.356656,
	})
	plot.AddPointGroup("TestGroup_2", Points, []float64{
		0.712863,
		0.975935,
		0.875864,
		0.737082,
		-0.185717,
		-0.936551,
		0.779397,
		0.916793,
		0.622004,
		-0.0860084,
	})
	plot.AddPointGroup("TestGroup_3", LinesPoints, []float64{
		0.28927,
		-0.945002,
		-0.904681,
		0.924912,
		0.990415,
		0.326935,
		-0.927919,
		0.994446,
		0.270194,
		-0.0378568,
	})

}

func TestThreePointGroupsFloatMixed(t *testing.T) {
	args := PlotArgs{
		Debug:   true,
		Persist: true,
		Format:  Pdf,
		Style:   Points,
	}

	plot, _ := NewPlot(1, args)
	plot.AddPointGroup("TestGroup^1", Points, []float32{
		-0.512695,
		0.591778,
		-0.0939544,
		-0.510766,
		-0.859442,
		0.0340482,
		0.887461,
		0.277168,
		-0.998753,
		0.356656,
	})
	plot.AddPointGroup("TestGroup^2", Points, []float64{
		0.712863,
		0.975935,
		0.875864,
		0.737082,
		-0.185717,
		-0.936551,
		0.779397,
		0.916793,
		0.622004,
		-0.0860084,
	})
	plot.AddPointGroup("TestGroup^3", LinesPoints, []float32{
		0.28927,
		-0.945002,
		-0.904681,
		0.924912,
		0.990415,
		0.326935,
		-0.927919,
		0.994446,
		0.270194,
		-0.0378568,
	})
}

func TestOnePointGroupInt8(t *testing.T) {
	args := PlotArgs{
		Debug:   true,
		Persist: true,
		Format:  Pdf,
		Style:   Points,
	}

	plot, _ := NewPlot(1, args)
	plot.AddPointGroup("TestGroup_1", Points, []int8{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		-9, -8, -7, -6, -5, -4, -3, -2, -1, 0,
	})

}
func TestOnePointGroupInt16(t *testing.T) {
	args := PlotArgs{
		Debug:   true,
		Persist: true,
		Format:  Pdf,
		Style:   Points,
	}

	plot, _ := NewPlot(1, args)
	plot.AddPointGroup("TestGroup_1", Points, []int16{
		1, 2, 8, 16, 32, 64, 128, 256, 512, 1024,
	})

}

func TestTwoPointGroupsInt16(t *testing.T) {
	args := PlotArgs{
		Debug:   true,
		Persist: true,
		Format:  Pdf,
		Style:   Points,
	}

	plot, _ := NewPlot(1, args)
	plot.AddPointGroup("PowerOfTwo^1", Points, []int16{
		1, 2, 8, 16, 32, 64, 128, 256, 512, 1024,
	})
	plot.AddPointGroup("Cubed^2", Points, []int16{
		0,
		1,
		cubeInt16(2),
		cubeInt16(3),
		cubeInt16(4),
		cubeInt16(5),
		cubeInt16(6),
		cubeInt16(7),
		cubeInt16(8),
		cubeInt16(9),
	})

}
