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

func (plot *Plot) plotX(pg *PointGroup) error {
	f, err := writePointsHelper(pg)
	if err != nil {
		return err
	}
	fname := f.Name()
	plot.tmpfiles[fname] = f
	cmd := plot.plotcmd
	if plot.nplots > 0 {
		cmd = plotCommand
	}
	if pg.style == "" {
		pg.style = defaultStyle
	}
	var line string
	if pg.name == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, pg.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, pg.name, pg.style)
	}
	plot.nplots++
	return plot.Cmd(line)
}

func (plot *Plot) plotXY(pg *PointGroup) error {
	f, err := writePointsHelper(pg)
	if err != nil {
		return err
	}
	fname := f.Name()
	plot.tmpfiles[fname] = f
	cmd := plot.plotcmd
	if plot.nplots > 0 {
		cmd = plotCommand
	}

	if pg.style == "" {
		pg.style = "points"
	}
	var line string
	if pg.name == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, pg.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, pg.name, pg.style)
	}
	plot.nplots++
	return plot.Cmd(line)
}

func (plot *Plot) plotXYZ(pg *PointGroup) error {
	f, err := writePointsHelper(pg)
	if err != nil {
		return err
	}
	fname := f.Name()
	plot.tmpfiles[fname] = f
	cmd := "splot" // Force 3D plot
	if plot.nplots > 0 {
		cmd = plotCommand
	}

	var line string
	if pg.name == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, pg.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, pg.name, pg.style)
	}
	plot.nplots++
	return plot.Cmd(line)
}

func writePointsHelper(points *PointGroup) (f *os.File, err error) {
	var npoints int // number of records to write to file
	var pointString string
	f, err = ioutil.TempFile(os.TempDir(), gGnuplotPrefix)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	switch points.dimensions {
	case 1:
		x := points.castedData.([][]float64)[0]
		npoints = len(x)
		for i := 0; i < npoints; i++ {
			pointString += fmt.Sprintf("%v\n", x[i])
			if i%10000 == 0 { // flush every 10,000 lines
				f.WriteString(pointString)
				pointString = ""
			}
		}
		f.WriteString(pointString)
	case 2:
		x := points.castedData.([][]float64)[0]
		y := points.castedData.([][]float64)[1]
		npoints = minValue(len(x), len(y))
		for i := 0; i < npoints; i++ {
			pointString += fmt.Sprintf("%v %v\n", x[i], y[i])
			if i%10000 == 0 { // flush every 10,000 lines
				f.WriteString(pointString)
				pointString = ""
			}
		}
		f.WriteString(pointString)
	case 3:
		x := points.castedData.([][]float64)[0]
		y := points.castedData.([][]float64)[1]
		z := points.castedData.([][]float64)[2]
		npoints = minValue(len(x), len(y), len(z))
		for i := 0; i < npoints; i++ {
			pointString += fmt.Sprintf("%v %v %v\n", x[i], y[i], z[i])
			if i%10000 == 0 { // flush every 10,000 lines
				f.WriteString(pointString)
				pointString = ""
			}
		}
		f.WriteString(pointString)
	default:
		return nil, &gnuplotError{
			fmt.Sprintf("invalid number of dims '%v'", points.dimensions),
		}
	}
	return
}

func minValue(n ...int) int {
	v := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] < v {
			v = n[i] // swap in smaller value
		}
	}
	return v
}
