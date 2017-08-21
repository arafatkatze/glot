// gnuplot is a simple minded set of functions to manage a gnuplot subprocess
// in order to plot data.
// See the gnuplot documentation page for the exact semantics of the gnuplot
// commands.
//  http://www.gnuplot.info/
package glot

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

var g_gnuplot_cmd string
var g_gnuplot_prefix string = "go-gnuplot-"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func init() {
	var err error
	g_gnuplot_cmd, err = exec.LookPath("gnuplot")
	if err != nil {
		fmt.Printf("** could not find path to 'gnuplot':\n%v\n", err)
		panic("could not find 'gnuplot'")
	}
}

type gnuplot_error struct {
	err string
}

func (e *gnuplot_error) Error() string {
	return e.err
}

type plotter_process struct {
	handle *exec.Cmd
	stdin  io.WriteCloser
}

func new_plotter_proc(persist bool) (*plotter_process, error) {
	proc_args := []string{}
	if persist {
		proc_args = append(proc_args, "-persist")
	}
	cmd := exec.Command(g_gnuplot_cmd, proc_args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	return &plotter_process{handle: cmd, stdin: stdin}, cmd.Start()
}

type tmpfiles_db map[string]*os.File

// Plotter is a handle to a gnuplot subprocess, forwarding commands
// via its stdin
type Plotter struct {
	proc     *plotter_process
	debug    bool
	plotcmd  string
	nplots   int    // number of currently active plots
	style    string // current plotting style
	tmpfiles tmpfiles_db
}

// Cmd sends a command to the gnuplot subprocess and returns an error
// if something bad happened in the gnuplot process.
// ex:
//   fname := "foo.dat"
//   err := p.Cmd("plot %s", fname)
//   if err != nil {
//     panic(err)
//   }
func (self *Plotter) Cmd(format string, a ...interface{}) error {
	cmd := fmt.Sprintf(format, a...) + "\n"
	n, err := io.WriteString(self.proc.stdin, cmd)

	if self.debug {
		//buf := new(bytes.Buffer)
		//io.Copy(buf, self.proc.handle.Stdout)
		fmt.Printf("cmd> %v", cmd)
		fmt.Printf("res> %v\n", n)
	}

	return err
}

// CheckedCmd is a convenience wrapper around Cmd: it will panic if the
// error returned by Cmd isn't nil.
// ex:
//   fname := "foo.dat"
//   p.CheckedCmd("plot %s", fname)
func (self *Plotter) CheckedCmd(format string, a ...interface{}) {
	err := self.Cmd(format, a...)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
}

// Close makes sure all resources used by the gnuplot subprocess are reclaimed.
// This method is typically called when the Plotter instance is not needed
// anymore. That's usually done via a defer statement:
//   p, err := gnuplot.NewPlotter(...)
//   if err != nil { /* handle error */ }
//   defer p.Close()
func (self *Plotter) Close() (err error) {
	if self.proc != nil && self.proc.handle != nil {
		self.proc.stdin.Close()
		err = self.proc.handle.Wait()
	}
	self.ResetPlot()
	return err
}

// PlotNd will create an n-dimensional plot (up to 3) with a title `title`
// and using the data from the var-arg `data`.
// example:
//  err = p.PlotNd(
//           "test Nd plot",
//           []float64{0,1,2,3}, // x-data
//           []float64{0,1,2,3}, // y-data
//           []float64{0,1,2,3}) // z-data
func (self *Plotter) PlotNd(title string, data ...[]float64) error {
	ndims := len(data)

	switch ndims {
	case 1:
		return self.PlotX(data[0], title)
	case 2:
		return self.PlotXY(data[0], data[1], title)
	case 3:
		return self.PlotXYZ(data[0], data[1], data[2], title)
	}

	return &gnuplot_error{fmt.Sprintf("invalid number of dims '%v'", ndims)}
}

// PlotX will create a 2-d plot using `data` as input and `title` as the plot
// title.
// The index of the element in the `data` slice will be used as the x-coordinate
// and its correspinding value as the y-coordinate.
// Example:
//  err = p.PlotX([]float64{10, 20, 30}, "my title")
func (self *Plotter) PlotX(data []float64, title string, params ...bool) error {
	f, err := ioutil.TempFile(os.TempDir(), g_gnuplot_prefix)
	if err != nil {
		return err
	}
	fname := f.Name()
	self.tmpfiles[fname] = f
	for _, d := range data {
		f.WriteString(fmt.Sprintf("%v\n", d))
	}
	f.Close()
	cmd := self.plotcmd
	if self.nplots > 0 {
		cmd = "replot"
	}

	var line string
	if title == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, self.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, title, self.style)
	}
	self.nplots += 1
	return self.Cmd(line)
}

// PlotXY will create a 2-d plot using `x` and `y` as input and `title` as
// the plot title.
// The values of the `x` slice will be used as x-coordinates and the matching
// values of `y` as y-coordinates (ie: for the same index).
// If the lengthes of the slices do not match, the range for the data will be
// the smallest size of the two slices.
// Example:
//  err = p.PlotXY(
//           []float64{10, 20, 30},
//           []float64{11, 22, 33, 44},
//           "my title")
func (self *Plotter) PlotXY(x, y []float64, title string) error {
	npoints := min(len(x), len(y))

	f, err := ioutil.TempFile(os.TempDir(), g_gnuplot_prefix)
	if err != nil {
		return err
	}
	fname := f.Name()
	self.tmpfiles[fname] = f

	for i := 0; i < npoints; i++ {
		f.WriteString(fmt.Sprintf("%v %v\n", x[i], y[i]))
	}

	f.Close()
	cmd := self.plotcmd
	if self.nplots > 0 {
		cmd = "replot"
	}

	var line string
	if title == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, self.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, title, self.style)
	}
	self.nplots += 1
	return self.Cmd(line)
}

// PlotXYZ will create a 3-d plot using `x`, `y` and `z` as input and
// `title` as the plot title.
// The data points to be plotted are the triplets (x[i], y[i], z[i]) where
// `i` runs from 0 to the smallest length of the 3 slices.
// Example:
//  err = p.PlotXYZ(
//           []float64{10, 20, 30},
//           []float64{11, 22, 33, 44},
//           []float64{111, 222, 333, 444, 555},
//           "my title")
func (self *Plotter) PlotXYZ(x, y, z []float64, title string) error {
	npoints := min(len(x), len(y))
	npoints = min(npoints, len(z))
	f, err := ioutil.TempFile(os.TempDir(), g_gnuplot_prefix)
	if err != nil {
		return err
	}
	fname := f.Name()
	self.tmpfiles[fname] = f

	for i := 0; i < npoints; i++ {
		f.WriteString(fmt.Sprintf("%v %v %v\n", x[i], y[i], z[i]))
	}

	f.Close()
	cmd := "splot" // Force 3D plot
	if self.nplots > 0 {
		cmd = "replot"
	}

	var line string
	if title == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, self.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, title, self.style)
	}
	self.nplots += 1
	return self.Cmd(line)
}

// Func is a 1-d function which can be plotted with gnuplot
type Func func(x float64) float64
type Func2d func(x float64) float64
type Func3d func(x float64, y float64) float64

// PlotFunc will create a 2-d plot using `data` as x-coordinates and `fct(x[i])`
// as the y-coordinates.
// Example:
//  fct := funct (x float64) float64 { return math.Exp(float64(x) + 2.) }
//  err = p.PlotFunc(
//           []float64{0,1,2,3,4,5},
//           fct,
//           "my title")
func (self *Plotter) PlotFunc(data []float64, fct Func, title string) error {

	f, err := ioutil.TempFile(os.TempDir(), g_gnuplot_prefix)
	if err != nil {
		return err
	}
	fname := f.Name()
	self.tmpfiles[fname] = f

	for _, x := range data {
		f.WriteString(fmt.Sprintf("%v %v\n", x, fct(x)))
	}

	f.Close()
	cmd := self.plotcmd
	if self.nplots > 0 {
		cmd = "replot"
	}

	var line string
	if title == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, self.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, title, self.style)
	}
	self.nplots += 1
	return self.Cmd(line)
}

// SetPlotCmd changes the command used for plotting by the gnuplot subprocess.
// Only valid plot commands are accepted (plot, splot)
func (self *Plotter) SetPlotCmd(cmd string) (err error) {
	switch cmd {
	case "plot", "splot":
		self.plotcmd = cmd
	default:
		err = errors.New("invalid plot cmd [" + cmd + "]")
	}
	return err
}

// SetStyle changes the style used by the gnuplot subprocess.
// Only valid styles are accepted:
//      "lines",
//      "points",
//      "linepoints",
// 		"impulses",
//      "dots",
// 		"steps",
// 		"errorbars",
// 		"boxes",
// 		"boxerrorbars",
// 		"pm3d"
func (self *Plotter) SetStyle(style string) (err error) {
	allowed := []string{
		"lines", "points", "linepoints",
		"impulses", "dots",
		"steps", "fill solid",
		"errorbars",
		"boxes", "lp",
		"boxerrorbars",
		"pm3d"}
	fmt.Println("These are dots")
	for _, s := range allowed {
		if s == style {
			self.style = style
			err = nil
			return err
		}
	}

	fmt.Printf("** style '%v' not in allowed list %v\n", style, allowed)
	fmt.Printf("** default to 'points'\n")
	self.style = "points"
	err = &gnuplot_error{fmt.Sprintf("invalid style '%s'", style)}

	return err
}

// SetXLabel changes the label for the x-axis
func (self *Plotter) SetXLabel(label string) error {
	return self.Cmd(fmt.Sprintf("set xlabel '%s'", label))
}

// SetYLabel changes the label for the y-axis
func (self *Plotter) SetYLabel(label string) error {
	return self.Cmd(fmt.Sprintf("set ylabel '%s'", label))
}

// SetZLabel changes the label for the z-axis
func (self *Plotter) SetZLabel(label string) error {
	return self.Cmd(fmt.Sprintf("set zlabel '%s'", label))
}

// SetLabels changes the labels for the x-,y- and z-axis in one go, depending
// on the size of the `labels` var-arg.
// Example:
//  err = p.SetLabels("x", "y", "z")
func (self *Plotter) SetLabels(labels ...string) error {
	ndims := len(labels)
	if ndims > 3 || ndims <= 0 {
		return &gnuplot_error{fmt.Sprintf("invalid number of dims '%v'", ndims)}
	}
	var err error = nil

	for i, label := range labels {
		switch i {
		case 0:
			ierr := self.SetXLabel(label)
			if ierr != nil {
				err = ierr
				return err
			}
		case 1:
			ierr := self.SetYLabel(label)
			if ierr != nil {
				err = ierr
				return err
			}
		case 2:
			ierr := self.SetZLabel(label)
			if ierr != nil {
				err = ierr
				return err
			}
		}
	}
	return nil
}

// ResetPlot clears up all plots and sets the Plotter state anew.
func (self *Plotter) ResetPlot() (err error) {
	for fname, fhandle := range self.tmpfiles {
		ferr := fhandle.Close()
		if ferr != nil {
			err = ferr
		}
		os.Remove(fname)
	}
	self.nplots = 0
	return err
}

// NewPlotter creates a new Plotter instance.
//  - `fname` is the name of the file containing commands (should be empty for now)
//  - `persist` is a flag to run the gnuplot subprocess with '-persist' so the
//    plot window isn't closed after sending a command
//  - `debug` is a flag to tell go-gnuplot to print out every command sent to
//    the gnuplot subprocess.
// Example:
//  p, err := gnuplot.NewPlotter("", false, false)
//  if err != nil { /* handle error */ }
//  defer p.Close()
func NewPlotter(fname string, persist, debug bool) (*Plotter, error) {
	p := &Plotter{proc: nil, debug: debug, plotcmd: "plot",
		nplots: 0, style: "points"}
	p.tmpfiles = make(tmpfiles_db)
	if fname != "" {
		panic("NewPlotter with fname is not yet supported")
	} else {
		proc, err := new_plotter_proc(persist)
		if err != nil {
			return nil, err
		}
		p.proc = proc
	}
	return p, nil
}
