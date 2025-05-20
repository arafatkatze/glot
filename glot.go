// Glot is a library for having simplified 1,2,3 Dimensional points/line plots
// It's built on top of Gnu plot and offers the ability to use Raw Gnu plot commands
// directly from golang.
// See the gnuplot documentation page for the exact semantics of the gnuplot
// commands.
//  http://www.gnuplot.info/

package glot

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Plot is the basic type representing a plot.
// Every plot has a set of Pointgroups that are simultaneously plotted
// on a 2/3 D plane given the plot type.
// The Plot dimensions must be specified at the time of construction
// and can't be changed later.  All the Pointgroups added to a plot must
// have same dimensions as the dimension specified at the
// the time of plot construction.
// The Pointgroups can be dynamically added and removed from a plot
// And style changes can also be made dynamically.
type Plot struct {
	proc       *plotterProcess
	debug      bool
	plotcmd    string
	nplots     int                    // number of currently active plots
	tmpfiles   tmpfilesDb             // A temporary file used for saving data
	dimensions int                    // dimensions of the plot
	PointGroup map[string]*PointGroup // A map between Curve name and curve type. This maps a name to a given curve in a plot. Only one curve with a given name exists in a plot.
	format     string                 // The saving format of the plot. This could be PDF, PNG, JPEG and so on.
	style      string                 // style of the plot
	title      string                 // The title of the plot.
}

// NewPlot Function makes a new plot with the specified dimensions.
//
// Usage
//  dimensions := 3
//  persist := false
//  debug := false
//  plot, _ := glot.NewPlot(dimensions, persist, debug)
// Variable definitions
//  dimensions  :=> refers to the dimensions of the plot.
//  debug       :=> can be used by developers to check the actual commands sent to gnu plot.
//  persist     :=> used to make the gnu plot window stay open.
func NewPlot(dimensions int, persist, debug bool) (*Plot, error) {
	p := &Plot{proc: nil, debug: debug, plotcmd: "plot",
		nplots: 0, dimensions: dimensions, style: "points", format: "png"}
	p.PointGroup = make(map[string]*PointGroup) // Adding a mapping between a curve name and a curve
	p.tmpfiles = make(tmpfilesDb)
	proc, err := newPlotterProc(persist)
	if err != nil {
		return nil, err
	}
	// Only 1,2,3 Dimensional plots are supported
	if dimensions > 3 || dimensions < 1 {
		return nil, &gnuplotError{fmt.Sprintf("invalid number of dims '%v'", dimensions)}
	}
	p.proc = proc
	return p, nil
}

func (plot *Plot) plotX(PointGroup *PointGroup) error {
	f, err := ioutil.TempFile(os.TempDir(), gGnuplotPrefix)
	if err != nil {
		return err
	}
	fname := f.Name()
	plot.tmpfiles[fname] = f
	for _, d := range PointGroup.castedData.([]float64) {
		f.WriteString(fmt.Sprintf("%v\n", d))
	}
	f.Close()
	cmd := plot.plotcmd
	if plot.nplots > 0 {
		cmd = plotCommand
	}
	if PointGroup.style == "" {
		PointGroup.style = defaultStyle
	}
	var line string
	if PointGroup.name == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, PointGroup.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, PointGroup.name, PointGroup.style)
	}
	plot.nplots++
	return plot.Cmd(line)
}

// ExamplePlot_AddSmaSeries demonstrates how to add a Simple Moving Average series to a plot.
func ExamplePlot_AddSmaSeries() {
	dimensions := 2
	persist := false // For actual display, you might want true
	debug := false

	plot, err := NewPlot(dimensions, persist, debug)
	if err != nil {
		// In a real application, handle the error appropriately.
		// For an example, panic might be acceptable or just printing.
		fmt.Printf("Error creating plot: %v\n", err)
		return
	}
	// defer plot.Close() // Assuming plot.Close() handles gnuplot process cleanup.
	// The actual Close method is not in the provided glot.go snippet.
	// In a real scenario, ensure this cleanup is correctly implemented.

	// Example data
	yValues := []float64{1.0, 2.0, 1.5, 2.5, 2.0, 3.0, 2.5, 3.5, 3.0, 4.0, 3.5}
	windowSize := 3
	seriesName := "Smoothed Data"

	err = plot.AddSmaSeries(seriesName, "lines", yValues, windowSize)
	if err != nil {
		fmt.Printf("Error adding SMA series: %v\n", err)
		return
	}

	// These Set methods are conceptual as they are not in the provided glot.go snippet.
	// In a full library, they would configure the plot's appearance.
	// plot.SetXLabel("Time")
	// plot.SetYLabel("Value")
	// plot.SetTitle("SMA Example")
	// As an alternative, one might use plot.Cmd for such settings:
	// plot.Cmd("set xlabel 'Time'")
	// plot.Cmd("set ylabel 'Value'")
	// plot.Cmd("set title 'SMA Example'")


	// In a real scenario, you would save the plot, e.g.:
	// if err := plot.SavePlot("sma_example.png"); err != nil {
	//     fmt.Printf("Error saving plot: %v\n", err)
	//     return
	// }
	// For this example, we'll just print a success message if the point group was added.
	if _, ok := plot.PointGroup[seriesName]; ok {
		fmt.Println("SMA series added and plot settings applied.")
	} else {
		fmt.Println("SMA series not found in plot.")
	}
	// Output: SMA series added and plot settings applied.
}

// AddSmaSeries calculates and adds a Simple Moving Average (SMA) series to the plot.
//
// Parameters:
//   name: The name of the SMA series.
//   style: The plotting style for the SMA series (e.g., "lines", "points").
//   data: The input data for which the SMA needs to be calculated.
//         It can be of two types:
//         - `[]float64`: Represents Y values. X values will be an implicit sequence (0, 1, 2,...).
//         - `[][]float64`: Represents X and Y values, where `data[0]` are X values and `data[1]` are Y values.
//   windowSize: The window size for the SMA calculation.
//
// Returns:
//   An error if the data type is unsupported, if windowSize is invalid,
//   or if there's an error adding the point group to the plot.
//
// Note:
//   The `CalculateSMA` function is expected to be available in the same package.
//   This method relies on `plot.AddPointGroup` to add the calculated SMA data to the plot.
func (plot *Plot) AddSmaSeries(name string, style string, data interface{}, windowSize int) error {
	var originalX, originalY, smaY, finalX []float64

	switch v := data.(type) {
	case []float64:
		originalY = v
		smaY = CalculateSMA(originalY, windowSize)
		if len(smaY) == 0 {
			return fmt.Errorf("failed to calculate SMA for %s: CalculateSMA returned empty slice, check window size (%d) and data length (%d)", name, windowSize, len(originalY))
		}
		// Generate implicit X values for SMA: 0, 1, ..., len(smaY)-1
		finalX = make([]float64, len(smaY))
		for i := range smaY {
			finalX[i] = float64(i)
		}

	case [][]float64:
		if len(v) != 2 {
			return fmt.Errorf("unsupported data format for [][]float64 in %s: expected two inner slices for X and Y values, got %d", name, len(v))
		}
		originalX = v[0]
		originalY = v[1]
		if len(originalX) != len(originalY) {
			return fmt.Errorf("data mismatch for %s: X and Y slices must have the same length (X: %d, Y: %d)", name, len(originalX), len(originalY))
		}

		smaY = CalculateSMA(originalY, windowSize)
		if len(smaY) == 0 {
			return fmt.Errorf("failed to calculate SMA for %s: CalculateSMA returned empty slice, check window size (%d) and data length (%d)", name, windowSize, len(originalY))
		}

		// Adjust X values: originalX[windowSize-1:]
		// The first SMA point corresponds to the window ending at originalY[windowSize-1],
		// so its X value is originalX[windowSize-1].
		if windowSize-1 >= len(originalX) {
			return fmt.Errorf("cannot determine X values for SMA series %s: windowSize (%d) is too large for the number of X data points (%d)", name, windowSize, len(originalX))
		}
		// Ensure we don't go out of bounds for originalX if smaY is shorter than expected (e.g. due to windowSize)
		endIndex := windowSize - 1 + len(smaY)
		if endIndex > len(originalX) {
			// This case should ideally not be hit if CalculateSMA is correct and len(smaY) = len(originalY) - windowSize + 1
			// However, as a safeguard:
			return fmt.Errorf("internal error or inconsistent data for %s: calculated SMA length implies out-of-bounds access to original X values", name)
		}
		finalX = originalX[windowSize-1 : endIndex]


	default:
		return fmt.Errorf("unsupported data type for AddSmaSeries (%s): %T. Supported types are []float64 and [][]float64", name, data)
	}

	if len(finalX) == 0 || len(smaY) == 0 {
		// This might happen if CalculateSMA returned an empty slice and it wasn't caught,
		// or if finalX somehow ended up empty.
		return fmt.Errorf("no data points to plot for SMA series %s after calculation (X: %d, Y: %d points)", name, len(finalX), len(smaY))
	}
	
	// Prepare data for AddPointGroup
	processedData := [][]float64{finalX, smaY}
	return plot.AddPointGroup(name, style, processedData)
}

func (plot *Plot) plotXY(PointGroup *PointGroup) error {
	x := PointGroup.castedData.([][]float64)[0]
	y := PointGroup.castedData.([][]float64)[1]
	npoints := min(len(x), len(y))

	f, err := ioutil.TempFile(os.TempDir(), gGnuplotPrefix)
	if err != nil {
		return err
	}
	fname := f.Name()
	plot.tmpfiles[fname] = f

	for i := 0; i < npoints; i++ {
		f.WriteString(fmt.Sprintf("%v %v\n", x[i], y[i]))
	}

	f.Close()
	cmd := plot.plotcmd
	if plot.nplots > 0 {
		cmd = plotCommand
	}

	if PointGroup.style == "" {
		PointGroup.style = "points"
	}
	var line string
	if PointGroup.name == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, PointGroup.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, PointGroup.name, PointGroup.style)
	}
	plot.nplots++
	return plot.Cmd(line)
}

func (plot *Plot) plotXYZ(points *PointGroup) error {
	x := points.castedData.([][]float64)[0]
	y := points.castedData.([][]float64)[1]
	z := points.castedData.([][]float64)[2]
	npoints := min(len(x), len(y))
	npoints = min(npoints, len(z))
	f, err := ioutil.TempFile(os.TempDir(), gGnuplotPrefix)
	if err != nil {
		return err
	}
	fname := f.Name()
	plot.tmpfiles[fname] = f

	for i := 0; i < npoints; i++ {
		f.WriteString(fmt.Sprintf("%v %v %v\n", x[i], y[i], z[i]))
	}

	f.Close()
	cmd := "splot" // Force 3D plot
	if plot.nplots > 0 {
		cmd = plotCommand
	}

	var line string
	if points.name == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, points.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, points.name, points.style)
	}
	plot.nplots++
	return plot.Cmd(line)
}
