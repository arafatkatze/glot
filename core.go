package glot

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

var gGnuplotCmd string
var gGnuplotPrefix = "go-gnuplot-"

const defaultStyle = "points" // The default style for a curve
const plotCommand = "replot"  // The default style for a curve

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Function to intialize the package and check for GNU plot installation
// This raises an error if GNU plot is not installed
func init() {
	var err error
	gGnuplotCmd, err = exec.LookPath("gnuplot")
	if err != nil {
		fmt.Printf("** could not find path to 'gnuplot':\n%v\n", err)
		panic("could not find 'gnuplot'")
	}
}

type gnuplotError struct {
	err string
}

func (e *gnuplotError) Error() string {
	return e.err
}

// plotterProcess is the type for handling gnu commands.
type plotterProcess struct {
	handle *exec.Cmd
	stdin  io.WriteCloser
}

// newPlotterProc function makes the plotterProcess struct
func newPlotterProc(persist bool) (*plotterProcess, error) {
	procArgs := []string{}
	if persist {
		procArgs = append(procArgs, "-persist")
	}
	cmd := exec.Command(gGnuplotCmd, procArgs...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	return &plotterProcess{handle: cmd, stdin: stdin}, cmd.Start()
}

// Cmd sends a command to the gnuplot subprocess and returns an error
// if something bad happened in the gnuplot process.
// ex:
//   fname := "foo.dat"
//   err := p.Cmd("plot %s", fname)
//   if err != nil {
//     panic(err)
//   }
func (plot *Plot) Cmd(format string, a ...interface{}) error {
	cmd := fmt.Sprintf(format, a...) + "\n"
	n, err := io.WriteString(plot.proc.stdin, cmd)
	if plot.debug {
		//buf := new(bytes.Buffer)
		//io.Copy(buf, plot.proc.handle.Stdout)
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
func (plot *Plot) CheckedCmd(format string, a ...interface{}) {
	err := plot.Cmd(format, a...)
	if err != nil {
		errString := fmt.Sprintf("** err: %v\n", err)
		panic(errString)
	}
}

// A map between os files and file names
type tmpfilesDb map[string]*os.File

// Close makes sure all resources used by the gnuplot subprocess are reclaimed.
// This method is typically called when the Plotter instance is not needed
// anymore. That's usually done via a defer statement:
//   p, err := gnuplot.NewPlotter(...)
//   if err != nil { /* handle error */ }
//   defer p.Close()
func (plot *Plot) Close() (err error) {
	if plot.proc != nil && plot.proc.handle != nil {
		plot.proc.stdin.Close()
		err = plot.proc.handle.Wait()
	}
	plot.ResetPlot()
	return err
}

func (plot *Plot) cleanplot() (err error) {
	plot.tmpfiles = make(tmpfilesDb)
	plot.nplots = 0
	return err
}

// ResetPlot is used to reset the whole plot.
// This removes all the PointGroup's from the plot and makes it new.
// Usage
//  plot.ResetPlot()
func (plot *Plot) ResetPlot() (err error) {
	plot.cleanplot()
	plot.PointGroup = make(map[string]*PointGroup) // Adding a mapping between a curve name and a curve
	return err
}
