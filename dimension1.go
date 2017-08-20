package glot

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Plot struct {
	proc       *plotter_process
	debug      bool
	plotcmd    string
	nplots     int // number of currently active plots
	tmpfiles   tmpfiles_db
	dimensions int
	Curves     map[string]*Curve
	format     string
	style      string
	empty      bool
	title      string
}

type Curve struct {
	name        string
	dimensions  int
	style       string // current plotting style
	data        interface{}
	casted_data interface{}
	labels      [3]string
	set         bool
	axisx       [2]float64
	axisy       [2]float64
	axisz       [2]float64
}



func (self *Plot) Addcurve(name string, style string, data interface{}) (err error) {
	_, exists := self.Curves[name]
	if exists {
			return &gnuplot_error{fmt.Sprintf("A curve with the name %s  already exists, please use another name of the curve or remove this curve before using another one with the same name.", name)}
	}

	curve := &Curve{name: name, dimensions: self.dimensions, data: data, set: true}
	allowed := []string{
		"lines", "points", "linepoints",
		"impulses", "dots",
		"steps", "fill solid", "histogram",
		"errorbars",
		"boxes", "lp"}
		curve.style = "points"
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
		curve.casted_data = data.([]float64)
		self.PlotX(curve)
		self.Curves[name] = curve
	case []float64:
		curve.casted_data = data.([]float64)
		self.PlotX(curve)
		self.Curves[name] = curve
	case []float32:
		original_slice := data.([]float32)
		type_casted_slice := make([]float64, len(original_slice))
		for i := 0; i < len(original_slice); i++ {
			type_casted_slice[i] = float64(original_slice[i])
		}
		curve.casted_data = type_casted_slice
		self.PlotX(curve)
		self.Curves[name] = curve
	case []int:
		original_slice := data.([]int)
		type_casted_slice := make([]float64, len(original_slice))
		for i := 0; i < len(original_slice); i++ {
			type_casted_slice[i] = float64(original_slice[i])
		}
		curve.casted_data = type_casted_slice
		self.PlotX(curve)
		self.Curves[name] = curve
	case []int8:
		original_slice := data.([]int8)
		type_casted_slice := make([]float64, len(original_slice))
		for i := 0; i < len(original_slice); i++ {
			type_casted_slice[i] = float64(original_slice[i])
		}
		curve.casted_data = type_casted_slice
		self.PlotX(curve)
		self.Curves[name] = curve
	case []int16:
		original_slice := data.([]int16)
		type_casted_slice := make([]float64, len(original_slice))
		for i := 0; i < len(original_slice); i++ {
			type_casted_slice[i] = float64(original_slice[i])
		}
		curve.casted_data = type_casted_slice
		self.PlotX(curve)
		self.Curves[name] = curve
	case []int32:
		original_slice := data.([]int32)
		type_casted_slice := make([]float64, len(original_slice))
		for i := 0; i < len(original_slice); i++ {
			type_casted_slice[i] = float64(original_slice[i])
		}
		curve.casted_data = type_casted_slice
		self.PlotX(curve)
		self.Curves[name] = curve
	case []int64:
		original_slice := data.([]int64)
		type_casted_slice := make([]float64, len(original_slice))
		for i := 0; i < len(original_slice); i++ {
			type_casted_slice[i] = float64(original_slice[i])
		}
		curve.casted_data = type_casted_slice
		self.PlotX(curve)
		self.Curves[name] = curve
	default:
		return &gnuplot_error{fmt.Sprintf("invalid number of dims ")}
	}
if discovered == 0 {
	fmt.Printf("** style '%v' not in allowed list %v\n", style, allowed)
	fmt.Printf("** default to 'points'\n")
	err = &gnuplot_error{fmt.Sprintf("invalid style '%s'", style)}
}
	return err
}

func NewPlot(dimensions int, persist, debug bool) (*Plot, error) {
	p := &Plot{proc: nil, debug: debug, plotcmd: "plot",
		nplots: 0, dimensions: dimensions, style: "points", format: "png"}
	if dimensions == 1 {
		p.style = "points" // Default style for a 1d plot.
	}
	p.Curves = make(map[string]*Curve) // Adding a mapping between a curve name and a curve
	p.tmpfiles = make(tmpfiles_db)
	proc, err := new_plotter_proc(persist)
	if err != nil {
		return nil, err
	}

	// Only 1,2,3 Dimensional curves are supported
	if dimensions > 3 || dimensions < 1 {
		return nil, &gnuplot_error{fmt.Sprintf("invalid number of dims '%v'", dimensions)}
	}
	p.proc = proc
	return p, nil

}

func (self *Plot) PlotX(curves *Curve) error {
	f, err := ioutil.TempFile(os.TempDir(), g_gnuplot_prefix)
	if err != nil {
		return err
	}
	fname := f.Name()
	self.tmpfiles[fname] = f
	for _, d := range curves.casted_data.([]float64) {
		f.WriteString(fmt.Sprintf("%v\n", d))
	}
	f.Close()
	cmd := self.plotcmd
	if self.nplots > 0 {
		cmd = "replot"
	}
 if curves.style == "" {
		curves.style = "points"
	}
	var line string
	if curves.name == "" {
		line = fmt.Sprintf("%s \"%s\" with %s", cmd, fname, curves.style)
	} else {
		line = fmt.Sprintf("%s \"%s\" title \"%s\" with %s",
			cmd, fname, curves.name, curves.style)
	}
	self.nplots += 1
	return self.Cmd(line)
}

func (self *Plot) Cmd(format string, a ...interface{}) error {
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

func (self *Plot) CheckedCmd(format string, a ...interface{}) {
	err := self.Cmd(format, a...)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
}

func (self *Plot) SetFormat(newformat string) error {
	allowed := []string{
		"png", "pdf"}
	for _, s := range allowed {
		if newformat == s {
			self.format = newformat
			return nil
		}
	}
	fmt.Printf("** Format '%v' not in allowed list %v\n", newformat, allowed)
	fmt.Printf("** default to 'png'\n")
	err := &gnuplot_error{fmt.Sprintf("invalid format '%s'", newformat)}
	return err
}

func (self *Plot) cleanplot() (err error) {
	self.tmpfiles = make(tmpfiles_db)
	self.nplots = 0
	return err
}

func (self *Plot) ResetPlot() (err error) {
	self.cleanplot()
	self.Curves = make(map[string]*Curve) // Adding a mapping between a curve name and a curve

	return err
}

func (self *Plot) Removecurve(name string){
	 delete(self.Curves, name);
		self.cleanplot()
		for _, curve := range self.Curves {
			  self.PlotX(curve)
		}
}

func (self *Plot) Resetcurvestyle(name string, style string) (err error) {
	 curve, exists := self.Curves[name]
		if !exists {
				return &gnuplot_error{fmt.Sprintf("A curve with name %s does not exist.", name)}
		}
	 self.Removecurve(name);
		curve.style = style
		self.PlotX(curve)
		return err
}


func (self *Plot) Close() (err error) {
	if self.proc != nil && self.proc.handle != nil {
		self.proc.stdin.Close()
		err = self.proc.handle.Wait()
	}
	self.ResetPlot()
	return err
}


func (self *Plot) SavePlot(filename string) (err error) {
	if self.nplots == 0 {
		return  &gnuplot_error{fmt.Sprintf("This plot has 0 curves and therefore its a redundant plot and it can't be printed.")}
	}
	output_format := "set terminal " + self.format
	self.CheckedCmd(output_format)
	output_file_command := "set output" + "'" + filename + "'"
	self.CheckedCmd(output_file_command)
	self.CheckedCmd("replot  ")
	return nil
}



func (self *Plot) PlotFunc(data []float64, fct Func, title string) error {

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

// SetXLabel changes the label for the x-axis
func (self *Plot) SetTitle(title string) error {
	return self.Cmd(fmt.Sprintf("set title \"%s\" ", title))
}

// SetXLabel changes the label for the x-axis
func (self *Plot) SetXLabel(label string) error {

	return self.Cmd(fmt.Sprintf("set xlabel '%s'", label))
}

// SetYLabel changes the label for the y-axis
func (self *Plot) SetYLabel(label string) error {
	return self.Cmd(fmt.Sprintf("set ylabel '%s'", label))
}

func (self *Plot) SetXrange(start int, end int) error {
	return self.Cmd(fmt.Sprintf("set xrange [%d:%d]", start, end))
}

func (self *Plot) SetYrange(start int, end int) error {
	return self.Cmd(fmt.Sprintf("set yrange [%d:%d]", start, end))
}

func (self *Plot) SetZrange(start int, end int) error {
	return self.Cmd(fmt.Sprintf("set zrange [%d:%d]", start, end))
}


// SetZLabel changes the label for the z-axis
func (self *Plot) SetZLabel(label string) error {
return self.Cmd(fmt.Sprintf("set xlabel '%s'", label))
}


func (self *Plot) SetLabels(labels ...string) error {
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
