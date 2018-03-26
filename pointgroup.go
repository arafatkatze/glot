package glot

import "fmt"

const incorrectDimErrMsg = "The dimensions of this PointGroup are not " +
	"compatible with the dimensions of the plot.\n " +
	"If you want to make a %d-d curve you must specify a %d-d plot."

// A PointGroup refers to a set of points that need to be plotted.
// It could either be a set of points or a function of co-ordinates.
// For Example z = Function(x,y)(3 Dimensional) or  y = Function(x) (2-Dimensional)
type PointGroup struct {
	name       string      // Name of the curve
	dimensions int         // dimensions of the curve
	style      PointStyle  // current plotting style
	data       interface{} // Data inside the curve in any integer/float format
	castedData interface{} // The data inside the curve typecasted to float64
	set        bool        //
	index      int         // Relative index of pointgroup in the plot
}

func (pg *PointGroup) setIndex(idx int) {
	pg.index = idx
}

// AddPointGroup function adds a group of points to a plot.
//
// Usage
//  dimensions := 2
//  persist := false
//  debug := false
//  plot, _ := glot.NewPlot(dimensions, persist, debug)
//  plot.AddPointGroup("Sample1", "points", []int32{51, 8, 4, 11})
//  plot.AddPointGroup("Sample2", "points", []int32{1, 2, 4, 11})
//  plot.SavePlot("1.png")
func (plot *Plot) AddPointGroup(name string, style PointStyle, data interface{}) (curve *PointGroup, err error) {
	curve, exists := plot.PointGroup[name]
	if exists {
		return curve, &gnuplotError{fmt.Sprintf("A PointGroup with the name %s  already exists, please use another name of the curve or remove this curve before using another one with the same name.", name)}
	}

	curve = &PointGroup{name: name,
		dimensions: plot.dimensions,
		data:       data,
		set:        true,
		style:      style,
	}
	// We want to make sure that pointGroups are added to figure in a
	// consistent and repeatable manner. Because we are using maps, the
	// order is inherently unpredictable and using an index for each group
	// allows us to have reproducible plots.
	curve.setIndex(plot.pointGroupSliceLen())
	discovered := 0
	// If the style value is an empty string and there's only a single
	// dimension, assume histogram by default.

	if style < 0 || style >= InvalidPointStyle {
		switch plot.dimensions {
		case 0:
			return nil, &gnuplotError{
				fmt.Sprintf("Wrong number of dimensions in this plot."),
			}
		case 1:
			curve.style = Histogram
			discovered = 1
		case 2, 3:
			curve.style = Points
			discovered = 1
		}
	} else {
		discovered++
	}

	switch data.(type) {
	case [][]float64:
		if plot.dimensions != len(data.([][]float64)) {
			return nil, &gnuplotError{fmt.Sprintf(incorrectDimErrMsg,
				plot.dimensions, plot.dimensions)}
		}
		curve.castedData = data.([][]float64)
		plot.PointGroup[name] = curve

	case [][]float32:
		if plot.dimensions != len(data.([][]float32)) {
			return nil, &gnuplotError{fmt.Sprintf(incorrectDimErrMsg,
				plot.dimensions, plot.dimensions)}
		}
		originalSlice := data.([][]float32)
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve

	case [][]int:
		if plot.dimensions != len(data.([][]int)) {
			return nil, &gnuplotError{fmt.Sprintf(incorrectDimErrMsg,
				plot.dimensions, plot.dimensions)}
		}
		originalSlice := data.([][]int)
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve

	case [][]int8:
		if plot.dimensions != len(data.([][]int8)) {
			return nil, &gnuplotError{fmt.Sprintf(incorrectDimErrMsg,
				plot.dimensions, plot.dimensions)}
		}
		originalSlice := data.([][]int8)
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve

	case [][]int16:
		if plot.dimensions != len(data.([][]int16)) {
			return nil, &gnuplotError{fmt.Sprintf(incorrectDimErrMsg,
				plot.dimensions, plot.dimensions)}
		}
		originalSlice := data.([][]int16)
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve

	case [][]int32:
		if plot.dimensions != len(data.([][]int32)) {
			return nil, &gnuplotError{fmt.Sprintf(incorrectDimErrMsg,
				plot.dimensions, plot.dimensions)}
		}
		originalSlice := data.([][]int32)
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve

	case [][]int64:
		if plot.dimensions != len(data.([][]int64)) {
			return nil, &gnuplotError{fmt.Sprintf(incorrectDimErrMsg,
				plot.dimensions, plot.dimensions)}
		}
		originalSlice := data.([][]int64)
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve

	case []float64:
		curve.castedData = data.([]float64)
		plot.PointGroup[name] = curve
	case []float32:
		originalSlice := data.([]float32)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve
	case []int:
		originalSlice := data.([]int)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve
	case []int8:
		originalSlice := data.([]int8)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve
	case []int16:
		originalSlice := data.([]int16)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve
	case []int32:
		originalSlice := data.([]int32)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve
	case []int64:
		originalSlice := data.([]int64)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.PointGroup[name] = curve
	default:
		return nil, &gnuplotError{fmt.Sprintf("invalid number of dimensions")}
	}
	if discovered == 0 {
		fmt.Printf("** style '%s' not supported ", style.String())
		fmt.Printf("** default to 'points'\n")
		err = &gnuplotError{fmt.Sprintf("invalid style '%s'", style)}
	}
	switch curve.dimensions {
	case 1:
		plot.plotX(curve)
	case 2:
		plot.plotXY(curve)
	case 3:
		plot.plotXYZ(curve)
	default:
		return curve, &gnuplotError{fmt.Sprintf("invalid number of dimensions")}
	}
	return curve, err
}

// RemovePointGroup helps to remove a particular point group from the plot.
// This way you can remove a pointgroup if it's un-necessary.
//
// Usage
//  dimensions := 3
//  persist := false
//  debug := false
//  plot, _ := glot.NewPlot(dimensions, persist, debug)
//  plot.AddPointGroup("Sample1", "points", []int32{51, 8, 4, 11})
//  plot.AddPointGroup("Sample2", "points", []int32{1, 2, 4, 11})
//  plot.RemovePointGroup("Sample1")
func (plot *Plot) RemovePointGroup(name string) {
	delete(plot.PointGroup, name)
	plot.cleanplot()
	for _, pointGroup := range plot.PointGroup {
		plot.plotX(pointGroup)
	}
}

// ResetPointGroupStyle helps to reset the style of a particular point group in a plot.
// Using both AddPointGroup and RemovePointGroup you can add or remove point groups.
// And dynamically change the plots.
//
// Usage
//  dimensions := 2
//  persist := false
//  debug := false
//  plot, _ := glot.NewPlot(dimensions, persist, debug)
//  plot.AddPointGroup("Sample1", "points", []int32{51, 8, 4, 11})
//  plot.ResetPointGroupStyle("Sample1", "points")
func (plot *Plot) ResetPointGroupStyle(name string, style PointStyle) (err error) {
	pointGroup, exists := plot.PointGroup[name]
	if !exists {
		return &gnuplotError{fmt.Sprintf("A curve with name %s does not exist.", name)}
	}
	plot.RemovePointGroup(name)
	pointGroup.style = style
	plot.plotX(pointGroup)
	return err
}
