package glot

import "testing"

var Args = PlotArgs{
	Debug:   true,
	Persist: true,
}

func TestSetLabels3Dims(t *testing.T) {
	dimensions := 3
	plot, _ := NewPlot(dimensions, Args)
	err := plot.SetLabels()
	if err == nil {
		t.Error("SetLabels raises error when empty string is passed")
	}
}

func TestSetLabels2Dims(t *testing.T) {
	dimensions := 2
	plot, _ := NewPlot(dimensions, Args)
	err := plot.SetLabels()
	if err == nil {
		t.Error("SetLabels raises error when empty string is passed")
	}
}

func TestSetLabels1Dims(t *testing.T) {
	dimensions := 1
	plot, _ := NewPlot(dimensions, Args)
	err := plot.SetLabels()
	if err == nil {
		t.Error("SetLabels raises error when empty string is passed")
	}
}

func TestSetFormat3Dims(t *testing.T) {
	dimensions := 3

	// Test unsupported formats
	plot, _ := NewPlot(dimensions, Args)
	err := plot.SetFormat(0)
	if err == nil {
		t.Error("SetLabels raises error when non-supported format is passed as an argument.")
	}

	// Test supported formats
	plot, _ = NewPlot(dimensions, Args)
	err = plot.SetFormat(Pdf)
	if err != nil {
		t.Error("SetLabels failed to set supported format (Pdf).")
	}
	err = plot.SetFormat(Png)
	if err != nil {
		t.Error("SetLabels failed to set supported format (Png).")
	}
}

func TestSetFormat2Dims(t *testing.T) {
	dimensions := 2

	// Test unsupported formats
	plot, _ := NewPlot(dimensions, Args)
	err := plot.SetFormat(0)
	if err == nil {
		t.Error("SetLabels raises error when non-supported format is passed as an argument.")
	}

	// Test supported formats
	plot, _ = NewPlot(dimensions, Args)
	err = plot.SetFormat(Pdf)
	if err != nil {
		t.Error("SetLabels failed to set supported format (Pdf).")
	}
	err = plot.SetFormat(Png)
	if err != nil {
		t.Error("SetLabels failed to set supported format (Png).")
	}
}

func TestSetFormat1Dims(t *testing.T) {
	dimensions := 1

	// Test unsupported formats
	plot, _ := NewPlot(dimensions, Args)
	err := plot.SetFormat(0)
	if err == nil {
		t.Error("SetLabels raises error when non-supported format is passed as an argument.")
	}

	// Test supported formats
	plot, _ = NewPlot(dimensions, Args)
	err = plot.SetFormat(Pdf)
	if err != nil {
		t.Error("SetLabels failed to set supported format (Pdf).")
	}
	err = plot.SetFormat(Png)
	if err != nil {
		t.Error("SetLabels failed to set supported format (Png).")
	}
}

func TestLinesStyleBrewerQualitative1(t *testing.T) {
	dimensions := 1

	// Test unsupported formats
	plot1, _ := NewPlot(dimensions, Args)
	plot2, _ := NewPlot(dimensions, Args)
	t1 := `
	set multiplot layout 2,2 ; 
	plot sin(x) ls 1 ; plot sin(x/2) ls 2 ;
	plot sin(x/4) ls 3 ; plot cos(x/2) ls 4
	`
	t2 := `
	set multiplot layout 2,2 ; 
	plot sin(x) ls 5 ; plot sin(x/2) ls 6 ;
	plot sin(x/4) ls 7 ; plot cos(x/2) ls 8
	`
	plot1.Cmd(t1)
	plot2.Cmd(t2)

}
