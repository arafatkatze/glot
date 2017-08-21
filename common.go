package glot

import (
	"fmt"
)

// SetTitle changes the label for the x-axis
func (plot *Plot) SetTitle(title string) error {
	return plot.Cmd(fmt.Sprintf("set title \"%s\" ", title))
}

// SetXLabel changes the label for the x-axis
func (plot *Plot) SetXLabel(label string) error {

	return plot.Cmd(fmt.Sprintf("set xlabel '%s'", label))
}

// SetYLabel changes the label for the y-axis
func (plot *Plot) SetYLabel(label string) error {
	return plot.Cmd(fmt.Sprintf("set ylabel '%s'", label))
}

func (plot *Plot) SetXrange(start int, end int) error {
	return plot.Cmd(fmt.Sprintf("set xrange [%d:%d]", start, end))
}

func (plot *Plot) SetYrange(start int, end int) error {
	return plot.Cmd(fmt.Sprintf("set yrange [%d:%d]", start, end))
}

func (plot *Plot) SetZrange(start int, end int) error {
	return plot.Cmd(fmt.Sprintf("set zrange [%d:%d]", start, end))
}

// SetZLabel changes the label for the z-axis
func (plot *Plot) SetZLabel(label string) error {
	return plot.Cmd(fmt.Sprintf("set xlabel '%s'", label))
}

// SetLabels Functions helps to set labels for x, y, z axis
func (plot *Plot) SetLabels(labels ...string) error {
	ndims := len(labels)
	if ndims > 3 || ndims <= 0 {
		return &gnuplotError{fmt.Sprintf("invalid number of dims '%v'", ndims)}
	}
	var err error

	for i, label := range labels {
		switch i {
		case 0:
			ierr := plot.SetXLabel(label)
			if ierr != nil {
				err = ierr
				return err
			}
		case 1:
			ierr := plot.SetYLabel(label)
			if ierr != nil {
				err = ierr
				return err
			}
		case 2:
			ierr := plot.SetZLabel(label)
			if ierr != nil {
				err = ierr
				return err
			}
		}
	}
	return nil
}

// SavePlot function is used to save the plot at this point.
// The plot is dynamic and additional pointgroups can be added and removed and different versions
// of the plot can be saved.
func (plot *Plot) SavePlot(filename string) (err error) {
	if plot.nplots == 0 {
		return &gnuplotError{fmt.Sprintf("This plot has 0 curves and therefore its a redundant plot and it can't be printed.")}
	}
	outputFormat := "set terminal " + plot.format
	plot.CheckedCmd(outputFormat)
	outputFileCommand := "set output" + "'" + filename + "'"
	plot.CheckedCmd(outputFileCommand)
	plot.CheckedCmd("replot  ")
	return nil
}

func (plot *Plot) SetFormat(newformat string) error {
	allowed := []string{
		"png", "pdf"}
	for _, s := range allowed {
		if newformat == s {
			plot.format = newformat
			return nil
		}
	}
	fmt.Printf("** Format '%v' not in allowed list %v\n", newformat, allowed)
	fmt.Printf("** default to 'png'\n")
	err := &gnuplotError{fmt.Sprintf("invalid format '%s'", newformat)}
	return err
}
