package glot

import (
	"fmt"
)

// A PointGroup refers to a set of points that need to plotted.
// It could either be a set of points or a function of co-ordinates.
// For Example z = Function(x,y)(3 Dimensional) or  y = Function(x) (2-Dimensional)
type PointGroup struct {
	name       string      // Name of the curve
	dimensions int         // dimensions of the curve
	style      string      // current plotting style
	data       interface{} // Data inside the curve in any integer/float format
	castedData interface{} // The data inside the curve typecasted to float64
	set        bool        //
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
func (plot *Plot) AddPointGroup(name string, style string, data interface{}) (err error) {
	_, exists := plot.PointGroup[name]
	if exists {
		return &gnuplotError{fmt.Sprintf("A PointGroup with the name %s  already exists, please use another name of the curve or remove this curve before using another one with the same name.", name)}
	}

	curve := &PointGroup{name: name, dimensions: plot.dimensions, data: data, set: true}
	allowed := []string{
		"lines", "points", "linepoints",
		"impulses", "dots", "bar",
		"steps", "fill solid", "histogram", "circle",
		"errorbars", "boxerrorbars",
		"boxes", "lp"}
	curve.style = defaultStyle
	discovered := 0
	for _, s := range allowed {
		if s == style {
			curve.style = style
			err = nil
			discovered = 1
		}
	}
	switch data.(type) {
	case [][]float64:
		if plot.dimensions != len(data.([][]float64)) {
			return &gnuplotError{fmt.Sprintf("The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot.")}
		}
		curve.castedData = data.([][]float64)
		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]float32:
		if plot.dimensions != len(data.([][]float32)) {
			return &gnuplotError{fmt.Sprintf("The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot.")}
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
		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]int:
		if plot.dimensions != len(data.([][]int)) {
			return &gnuplotError{fmt.Sprintf("The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot.")}
		}
		originalSlice := data.([][]int)
		if len(originalSlice) != 2 {
			return &gnuplotError{fmt.Sprintf("this is not a 2d matrix")}
		}
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice
		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]int8:
		if plot.dimensions != len(data.([][]int8)) {
			return &gnuplotError{fmt.Sprintf("The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot.")}
		}
		originalSlice := data.([][]int8)
		if len(originalSlice) != 2 {
			return &gnuplotError{fmt.Sprintf("this is not a 2d matrix")}
		}
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice

		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]int16:
		if plot.dimensions != len(data.([][]int16)) {
			return &gnuplotError{fmt.Sprintf("The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot.")}
		}
		originalSlice := data.([][]int16)
		if len(originalSlice) != 2 {
			return &gnuplotError{fmt.Sprintf("this is not a 2d matrix")}
		}
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice

		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]int32:
		if plot.dimensions != len(data.([][]int32)) {
			return &gnuplotError{fmt.Sprintf("The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot.")}
		}
		originalSlice := data.([][]int32)
		if len(originalSlice) != 2 {
			return &gnuplotError{fmt.Sprintf("this is not a 2d matrix")}
		}
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice

		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case [][]int64:
		if plot.dimensions != len(data.([][]int64)) {
			return &gnuplotError{fmt.Sprintf("The dimensions of this PointGroup are not compatible with the dimensions of the plot.\nIf you want to make a 2-d curve you must specify a 2-d plot.")}
		}
		originalSlice := data.([][]int64)
		if len(originalSlice) != 2 {
			return &gnuplotError{fmt.Sprintf("this is not a 2d matrix")}
		}
		typeCasteSlice := make([][]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = make([]float64, len(originalSlice[i]))
			for j := 0; j < len(originalSlice[i]); j++ {
				typeCasteSlice[i][j] = float64(originalSlice[i][j])
			}
		}
		curve.castedData = typeCasteSlice

		if plot.dimensions == 2 {
			plot.plotXY(curve)
		} else {
			plot.plotXYZ(curve)
		}
		plot.PointGroup[name] = curve

	case []float64:
		curve.castedData = data.([]float64)
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []float32:
		originalSlice := data.([]float32)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []int:
		originalSlice := data.([]int)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []int8:
		originalSlice := data.([]int8)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []int16:
		originalSlice := data.([]int16)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []int32:
		originalSlice := data.([]int32)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	case []int64:
		originalSlice := data.([]int64)
		typeCasteSlice := make([]float64, len(originalSlice))
		for i := 0; i < len(originalSlice); i++ {
			typeCasteSlice[i] = float64(originalSlice[i])
		}
		curve.castedData = typeCasteSlice
		plot.plotX(curve)
		plot.PointGroup[name] = curve
	default:
		return &gnuplotError{fmt.Sprintf("invalid number of dims ")}

	}
	if discovered == 0 {
		fmt.Printf("** style '%v' not in allowed list %v\n", style, allowed)
		fmt.Printf("** default to 'points'\n")
		err = &gnuplotError{fmt.Sprintf("invalid style '%s'", style)}
	}
	return err
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
func (plot *Plot) ResetPointGroupStyle(name string, style string) (err error) {
	pointGroup, exists := plot.PointGroup[name]
	if !exists {
		return &gnuplotError{fmt.Sprintf("A curve with name %s does not exist.", name)}
	}
	plot.RemovePointGroup(name)
	pointGroup.style = style
	plot.plotX(pointGroup)
	return err
}
