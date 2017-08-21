package glot

// Func2d is a 2-d function which can be plotted with gnuplot
type Func2d func(x float64) float64

// Func3d is a 3-d function which can be plotted with gnuplot
type Func3d func(x float64, y float64) float64

// AddFunc2d is used
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

func (plot *Plot) AddFunc3d(name string, style string, x []float64, y []float64, fct Func3d) error {
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
