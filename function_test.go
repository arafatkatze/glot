package glot

import (
	"reflect"
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

// TestCalculateSMABasicCases tests basic functionality of CalculateSMA with valid inputs.
func TestCalculateSMABasicCases(t *testing.T) {
	tests := []struct {
		name       string
		data       []float64
		windowSize int
		expected   []float64
	}{
		{
			name:       "Simple case",
			data:       []float64{1, 2, 3, 4, 5},
			windowSize: 3,
			expected:   []float64{2, 3, 4},
		},
		{
			name:       "Different numbers and window size",
			data:       []float64{10, 20, 30, 40, 50, 60},
			windowSize: 2,
			expected:   []float64{15, 25, 35, 45, 55},
		},
		{
			name:       "No data points",
			data:       []float64{},
			windowSize: 3,
			expected:   []float64{},
		},
		{
			name:       "Window size larger than data length",
			data:       []float64{1, 2},
			windowSize: 3,
			expected:   []float64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CalculateSMA(tt.data, tt.windowSize)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("CalculateSMA(%v, %d) = %v; want %v", tt.data, tt.windowSize, actual, tt.expected)
			}
		})
	}
}

// TestCalculateSMAEdgeWindowSize tests edge cases for windowSize parameter.
func TestCalculateSMAEdgeWindowSize(t *testing.T) {
	tests := []struct {
		name       string
		data       []float64
		windowSize int
		expected   []float64
	}{
		{
			name:       "Window size = 1",
			data:       []float64{1, 2, 3, 4, 5},
			windowSize: 1,
			expected:   []float64{1, 2, 3, 4, 5},
		},
		{
			name:       "Window size = data length",
			data:       []float64{1, 2, 3, 4, 5},
			windowSize: 5,
			expected:   []float64{3}, // (1+2+3+4+5)/5 = 3
		},
		{
			name:       "Window size = data length, single element",
			data:       []float64{10},
			windowSize: 1,
			expected:   []float64{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CalculateSMA(tt.data, tt.windowSize)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("CalculateSMA(%v, %d) = %v; want %v", tt.data, tt.windowSize, actual, tt.expected)
			}
		})
	}
}

// TestCalculateSMAEdgeData tests edge cases for the input data slice.
func TestCalculateSMAEdgeData(t *testing.T) {
	tests := []struct {
		name       string
		data       []float64
		windowSize int
		expected   []float64
	}{
		{
			name:       "Empty data slice",
			data:       []float64{},
			windowSize: 3,
			expected:   []float64{},
		},
		{
			name:       "Data slice with fewer elements than windowSize",
			data:       []float64{1, 2},
			windowSize: 3,
			expected:   []float64{},
		},
		{
			name:       "Data slice with fewer elements than windowSize, windowSize = 1", // This case is valid.
			data:       []float64{},
			windowSize: 1,
			expected:   []float64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CalculateSMA(tt.data, tt.windowSize)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("CalculateSMA(%v, %d) = %v; want %v", tt.data, tt.windowSize, actual, tt.expected)
			}
		})
	}
}

// TestCalculateSMAInvalidWindowSize tests invalid windowSize values.
func TestCalculateSMAInvalidWindowSize(t *testing.T) {
	tests := []struct {
		name       string
		data       []float64
		windowSize int
		expected   []float64
	}{
		{
			name:       "Window size = 0",
			data:       []float64{1, 2, 3, 4, 5},
			windowSize: 0,
			expected:   []float64{},
		},
		{
			name:       "Window size < 0",
			data:       []float64{1, 2, 3, 4, 5},
			windowSize: -1,
			expected:   []float64{},
		},
		{
			name:       "Window size = 0, empty data",
			data:       []float64{},
			windowSize: 0,
			expected:   []float64{},
		},
		{
			name:       "Window size < 0, empty data",
			data:       []float64{},
			windowSize: -2,
			expected:   []float64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CalculateSMA(tt.data, tt.windowSize)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("CalculateSMA(%v, %d) = %v; want %v", tt.data, tt.windowSize, actual, tt.expected)
			}
		})
	}
}
