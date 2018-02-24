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
	plotcmd    PlotCommand
	nplots     int                    // number of currently active plots
	tmpfiles   tmpfilesDb             // A temporary file used for saving data
	dimensions int                    // dimensions of the plot
	PointGroup map[string]*PointGroup // A map between Curve name and curve type. This maps a name to a given curve in a plot. Only one curve with a given name exists in a plot.
	format     PlotFormat             // The saving format of the plot. This could be PDF, PNG, JPEG and so on.
	style      PointStyle             // style of the plot
	title      string                 // The title of the plot.
}

const (
	// PlotDefaultStyle results in assignment of whatever style
	// the global plot object is using.
	PlotDefaultStyle = iota
	// Points is default style which which plots are created
	Points
	// Bar is a Barplot type
	Bar
	BoxErrorBars
	Boxplot
	Circle
	Dots
	ErrorBars
	FillSolid
	// Histogram, i.e. fancy barplot
	Histogram
	// Lines is a lineplot
	Lines
	// LinesPoints is a lineplot with points
	LinesPoints
	Steps
	InvalidPointStyle

	// Pdf is a pdf output format
	Pdf = iota
	// Png is a png output format
	Png
)

// PointStyle specifies which style to use for plotting a set of points.
// Points here imply data points not points as in points on the plot, since
// on the plot points may be represented by boxes, circles, etc.
type PointStyle uint

// String is an implementation of the Stringer Interface
// for PointStyle type.
func (p PointStyle) String() string {
	var m = map[PointStyle]string{
		Bar:          "bar",
		BoxErrorBars: "boxerrorbars",
		Boxplot:      "boxplot",
		Circle:       "circle",
		Dots:         "dots",
		ErrorBars:    "errorbars",
		FillSolid:    "fill solid",
		Histogram:    "histogram",
		Lines:        "lines",
		LinesPoints:  "linespoints",
		Points:       "points",
	}
	if _, ok := m[p]; !ok {
		return m[Points]
	}
	return m[p]
}

// PlotFormat ...
type PlotFormat uint

// String is an implementation of the Stringer Interface
// for PlotFormat type.
func (pf PlotFormat) String() string {
	switch pf {
	case Pdf:
		return "pdf"
	case Png:
		return "png"
	default:
		return "unsupported"
	}
}

// PlotCommand ...
type PlotCommand string

// PlotArgs ...
type PlotArgs struct {
	Debug   bool
	Format  PlotFormat
	Persist bool
	Command PlotCommand
	Style   PointStyle
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
func NewPlot(dimensions int, args PlotArgs) (*Plot, error) {
	p := &Plot{
		proc:  nil,
		debug: args.Debug,
		plotcmd: func() PlotCommand {
			if args.Command == "" {
				return "plot"
			}
			return args.Command
		}(),
		nplots:     0,
		dimensions: dimensions,
		style:      args.Style,
		format:     args.Format,
	}
	p.PointGroup = make(map[string]*PointGroup) // Adding a mapping between a curve name and a curve
	p.tmpfiles = make(tmpfilesDb)
	proc, err := newPlotterProc(args.Persist)
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

// ChangeStyle changes PointStyle currently configured on plot
// to style, assuming value is valid. If outside the allowed range,
// i.e. invalid value, use PlotDefaultStyle instead.
func (plot *Plot) ChangeStyle(style PointStyle) error {
	if style < PlotDefaultStyle || style >= InvalidPointStyle {
		plot.style = PlotDefaultStyle
		return &gnuplotError{
			fmt.Sprintf("invalid style given, using plot default"),
		}
	}
	plot.style = style
	return nil
}

func (plot *Plot) pointGroupSliceLen() int {
	pgs, err := plot.pointGroupSlice()
	if err != nil {
		return 0
	}
	return len(pgs)
}
func (plot *Plot) pointGroupSlice() ([]*PointGroup, error) {
	pgsl := []*PointGroup{}
	if len(plot.PointGroup) == 0 {
		return []*PointGroup{},
			&gnuplotError{fmt.Sprintf("no pointgroups were found")}
	}
	for _, pg := range plot.PointGroup {
		pgsl = append(pgsl, pg)
	}
	return pgsl, nil
}

func (plot *Plot) plotX(PointGroup *PointGroup) error {
	f, err := ioutil.TempFile(os.TempDir(), gGnuplotPrefix)
	if err != nil {
		return err
	}
	defer f.Close()

	fname := f.Name()
	plot.tmpfiles[fname] = f
	for _, d := range PointGroup.castedData.([]float64) {
		f.WriteString(fmt.Sprintf("%v\n", d))
	}

	var cmd PlotCommand
	if plot.nplots > 0 {
		cmd = ""
	} else {
		cmd = plot.plotcmd
	}

	if PointGroup.style < 0 || PointGroup.style >= InvalidPointStyle {
		PointGroup.style = plot.style
	}
	var line string
	if PointGroup.name == "" {

		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, PointGroup.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, PointGroup.name, PointGroup.style)
	}
	if plot.nplots > 0 {
		plot.plotcmd = plot.plotcmd + ", " + PlotCommand(line)
	} else {
		plot.plotcmd = PlotCommand(line)
	}
	plot.nplots++
	return plot.Cmd(line)
}

func (plot *Plot) plotXY(PointGroup *PointGroup) error {
	x := PointGroup.castedData.([][]float64)[0]
	y := PointGroup.castedData.([][]float64)[1]
	npoints := min(len(x), len(y))
	pointString := ""
	f, err := ioutil.TempFile(os.TempDir(), gGnuplotPrefix)
	if err != nil {
		return err
	}
	defer f.Close()

	fname := f.Name()
	plot.tmpfiles[fname] = f

	for i := 0; i < npoints; i++ {
		pointString += fmt.Sprintf("%v %v\n", x[i], y[i])
		// f.WriteString(fmt.Sprintf("%v %v\n", x[i], y[i]))
		if i%10000 == 0 { // flush every 10,000 lines
			f.WriteString(pointString)
			pointString = ""
		}
	}
	f.WriteString(pointString)

	var cmd PlotCommand
	if plot.nplots > 0 {
		cmd = ""
	} else {
		cmd = plot.plotcmd
	}

	var line string
	if PointGroup.name == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, PointGroup.style.String())
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, PointGroup.name, PointGroup.style)
	}
	if plot.nplots > 0 {
		plot.plotcmd = plot.plotcmd + ", " + PlotCommand(line)
	} else {
		plot.plotcmd = PlotCommand(line)
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
	defer f.Close()
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
