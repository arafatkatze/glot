package glot

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewPlot(t *testing.T) {
	persist := false
	debug := true
	_, err := NewPlot(0, persist, debug)
	if err == nil {
		t.Error("Expected error when making a 0 dimensional plot.")
	}
}

// Mock Close method for Plot, as its actual implementation is not provided
// This is to fulfill the `defer plot.Close()` requirement from the prompt.
// In a real scenario, this would be part of the Plot's definition in glot.go.
func (p *Plot) Close() {
	// If p.proc is not nil and has a close method, call it.
	// e.g., if p.proc != nil { p.proc.close() }
	// For this test environment, we can leave it empty or log.
	// fmt.Println("Plot.Close() called for test purposes.")
}

func TestAddSmaSeries_Float64Data_ImplicitX(t *testing.T) {
	plot, err := NewPlot(2, false, false)
	if err != nil {
		t.Fatalf("NewPlot failed: %v", err)
	}
	defer plot.Close()

	name := "sma_test1"
	style := "lines"
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	windowSize := 3

	err = plot.AddSmaSeries(name, style, data, windowSize)
	if err != nil {
		t.Fatalf("AddSmaSeries failed: %v", err)
	}

	pg, ok := plot.PointGroup[name]
	if !ok {
		t.Fatalf("PointGroup '%s' not found in plot", name)
	}

	expectedCastedData := [][]float64{{0.0, 1.0, 2.0}, {2.0, 3.0, 4.0}}
	if !reflect.DeepEqual(pg.castedData, expectedCastedData) {
		t.Errorf("For PointGroup '%s', expected castedData %v, got %v", name, expectedCastedData, pg.castedData)
	}
	if pg.style != style {
		t.Errorf("For PointGroup '%s', expected style '%s', got '%s'", name, style, pg.style)
	}
}

func TestAddSmaSeries_Float64SliceData_ExplicitX(t *testing.T) {
	plot, err := NewPlot(2, false, false)
	if err != nil {
		t.Fatalf("NewPlot failed: %v", err)
	}
	defer plot.Close()

	name := "sma_test2"
	style := "points"
	data := [][]float64{{10.0, 11.0, 12.0, 13.0, 14.0}, {1.0, 2.0, 3.0, 4.0, 5.0}}
	windowSize := 3

	err = plot.AddSmaSeries(name, style, data, windowSize)
	if err != nil {
		t.Fatalf("AddSmaSeries failed: %v", err)
	}

	pg, ok := plot.PointGroup[name]
	if !ok {
		t.Fatalf("PointGroup '%s' not found in plot", name)
	}

	expectedCastedData := [][]float64{{12.0, 13.0, 14.0}, {2.0, 3.0, 4.0}}
	if !reflect.DeepEqual(pg.castedData, expectedCastedData) {
		t.Errorf("For PointGroup '%s', expected castedData %v, got %v", name, expectedCastedData, pg.castedData)
	}
	if pg.style != style {
		t.Errorf("For PointGroup '%s', expected style '%s', got '%s'", name, style, pg.style)
	}
}

func TestAddSmaSeries_WindowSizeLeadsToEmptySMA(t *testing.T) {
	plot, err := NewPlot(2, false, false)
	if err != nil {
		t.Fatalf("NewPlot failed: %v", err)
	}
	defer plot.Close()

	name := "sma_empty_test"
	data := []float64{1.0, 2.0}
	windowSize := 3

	err = plot.AddSmaSeries(name, "lines", data, windowSize)
	if err == nil {
		t.Fatalf("AddSmaSeries expected to fail for windowSize leading to empty SMA, but it succeeded")
	}

	// Check specific error content if desired, e.g. "CalculateSMA returned empty slice"
	// For example: if !strings.Contains(err.Error(), "CalculateSMA returned empty slice") {
	//	 t.Errorf("Expected error to contain 'CalculateSMA returned empty slice', got: %v", err)
	// }


	if _, ok := plot.PointGroup[name]; ok {
		t.Errorf("PointGroup '%s' should not have been added to plot after error", name)
	}
}

func TestAddSmaSeries_WindowSizeOne(t *testing.T) {
	plot, err := NewPlot(2, false, false)
	if err != nil {
		t.Fatalf("NewPlot failed: %v", err)
	}
	defer plot.Close()

	name := "sma_win_one"
	style := "lines"
	data := [][]float64{{1.0, 2.0, 3.0}, {10.0, 20.0, 30.0}}
	windowSize := 1

	err = plot.AddSmaSeries(name, style, data, windowSize)
	if err != nil {
		t.Fatalf("AddSmaSeries failed: %v", err)
	}

	pg, ok := plot.PointGroup[name]
	if !ok {
		t.Fatalf("PointGroup '%s' not found in plot", name)
	}

	expectedCastedData := [][]float64{{1.0, 2.0, 3.0}, {10.0, 20.0, 30.0}}
	if !reflect.DeepEqual(pg.castedData, expectedCastedData) {
		t.Errorf("For PointGroup '%s', expected castedData %v, got %v", name, expectedCastedData, pg.castedData)
	}
	if pg.style != style {
		t.Errorf("For PointGroup '%s', expected style '%s', got '%s'", name, style, pg.style)
	}
}

func TestAddSmaSeries_InvalidDataType(t *testing.T) {
	plot, err := NewPlot(2, false, false)
	if err != nil {
		t.Fatalf("NewPlot failed: %v", err)
	}
	defer plot.Close()

	name := "sma_invalid_type"
	data := "not_a_slice" // Invalid data type
	windowSize := 3

	err = plot.AddSmaSeries(name, "lines", data, windowSize)
	if err == nil {
		t.Fatalf("AddSmaSeries expected to fail for invalid data type, but it succeeded")
	}

	if !strings.Contains(err.Error(), "unsupported data type") {
		t.Errorf("Expected error to contain 'unsupported data type', got: %v", err.Error())
	}

	if _, ok := plot.PointGroup[name]; ok {
		t.Errorf("PointGroup '%s' should not have been added to plot after error", name)
	}
}
