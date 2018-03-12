package glot

import "fmt"

// Func2d is a 2-d function which can be plotted with gnuplot
type Func2d func(x float64) float64

// Func3d is a 3-d function which can be plotted with gnuplot
type Func3d func(x float64, y float64) float64

// AddFunc2d is used to make a 2-d plot of the format y = Function(x)
//
// Usage
//  dimensions := 2
//  persist := false
//  debug := false
//  plot, _ := glot.NewPlot(dimensions, persist, debug)
//  fct := func(x float64) float64 { return (math.Exp(x)) }
//  groupName := "Exponential Curve"
//  style := "lines"
//  pointsX := []float64{1, 2, 3, 4, 5}
//  plot.AddFunc2d(groupName, style, pointsX, fct)
//  plot.SavePlot("1.png")
// Variable definitions
//  dimensions  :=> refers to the dimensions of the plot.
//  debug       :=> can be used by developers to check the actual commands sent to gnu plot.
//  persist     :=> used to make the gnu plot window stay open.
//  groupName   :=> Name of the curve
//  style       :=> Style of the curve
//  pointsX     :=> The x Value of the points to be plotted.  y = func(x) is plotted on the curve.
//  style       :=> Style of the curve
// NOTE: Currently only float64 type is supported for this function
func (plot *Plot) AddFunc2d(name string, style string, x []float64, fct Func2d) error {
	y := make([]float64, len(x))
	for index := range x {
		y[index] = fct(x[index])
	}
	combined := [][]float64{}
	combined = append(combined, x)
	combined = append(combined, y)
	plot.AddPointGroup(name, style, combined)
	return nil
}

// AddFunc3d is used to make a 3-d plot of the format z = Function(x,y)
//
// Usage
//  dimensions := 3
//  persist := false
//  debug := false
//  plot, _ := glot.NewPlot(dimensions, persist, debug)
//  fct := func(x, y float64) float64 { return x - y }
//  groupName := "Stright Line"
//  style := "lines"
//  pointsY := []float64{1, 2, 3, 4, 5}
//  pointsX := []float64{1, 2, 3, 4, 5}
//  plot.AddFunc3d(groupName, style, pointsX, pointsY, fct)
//  plot.SetXrange(0, 5)
//  plot.SetYrange(0, 5)
//  plot.SetZrange(0, 5)
//  plot.SavePlot("1.png")
// Variable definitions
//  dimensions  :=> refers to the dimensions of the plot.
//  debug       :=> can be used by developers to check the actual commands sent to gnu plot.
//  persist     :=> used to make the gnu plot window stay open.
//  groupName   :=> Name of the curve
//  style       :=> Style of the curve
//  pointsX     :=> The x Value of the points to be plotted.  y = func(x) is plotted on the curve.
// NOTE: Currently only float64 type is supported for this function
func (plot *Plot) AddFunc3d(name string, style string, x []float64, y []float64, fct Func3d) error {
	if len(x) != len(y) {
		return &gnuplotError{fmt.Sprintf("The length of the x-axis array and y-axis array are not same.")}
	}
	z := make([]float64, len(x))
	for index := range x {
		z[index] = fct(x[index], y[index])
	}
	combined := [][]float64{}
	combined = append(combined, x)
	combined = append(combined, y)
	combined = append(combined, z)
	plot.AddPointGroup(name, style, combined)
	return nil
}
