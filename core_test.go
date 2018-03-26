package glot

import (
	"bytes"
	"crypto/md5"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"testing"
)

func TestMin(t *testing.T) {
	var v int
	v = min(1, 2)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestCmd(t *testing.T) {
	base2n := func(n float64) []float64 {
		res := []float64{}
		for i := 0.; i < n; i++ {
			res = append(res, math.Pow(2, i))
		}
		return res
	}
	var testArgs = PlotArgs{
		Debug:   true,
		Format:  Pdf,
		Persist: true,
		Style:   Points,
		Command: "plot ",
	}
	var pg = &PointGroup{
		name:       "TestCmd",
		dimensions: 1,
		style:      Points,
		castedData: base2n(9),
	}
	p, err := NewPlot(1, testArgs)
	if err != nil {
		t.Errorf("NewPlot failed creation with: %v", err)
	}
	p.plotX(pg)
}

func TestGeneratePngOutput1d1(t *testing.T) {
	var testArgs = PlotArgs{
		Debug:   true,
		Format:  Png,
		Persist: true,
		Style:   Points,
	}
	plot, err := NewPlot(1, testArgs)
	if err != nil {
		t.Errorf("NewPlot failed creation with: %v", err)
	}
	pts := makeSomeInts(make([]int, 128), true)
	plot.AddPointGroup("SaveAs PNG test", Points, pts)

	tmpDir, err := ioutil.TempDir("", "glot")
	if err != nil {
		t.Errorf("%v", err)
	}
	tmpPlotFile := filepath.Join(tmpDir, "f.png")
	defer os.RemoveAll(tmpDir)
	plot.SavePlot(tmpPlotFile)
	plot.Close()

	// Open file, read it into a checksum function and generate checksum.
	fh, err := os.Open(tmpPlotFile)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer fh.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, fh); err != nil {
		t.Errorf("%v", err)
	}
	// Compare expected checksum against actual checksum.
	// If there is a mismatch, test should fail.
	expected := []byte{
		114, 231, 226, 101, 216, 182, 101, 224,
		51, 18, 92, 252, 249, 110, 248, 105,
	}
	if bytes.Compare(expected, hash.Sum(nil)) != 0 {
		t.Errorf("Failed checksum, expected: %x got: %x",
			expected, hash.Sum(nil))
	}
}

func TestGeneratePdfOutput1d1(t *testing.T) {
	var testArgs = PlotArgs{
		Debug:   true,
		Format:  Pdf,
		Persist: true,
		Style:   Points,
	}
	plot, err := NewPlot(1, testArgs)
	if err != nil {
		t.Errorf("NewPlot failed creation with: %v", err)
	}
	pts := makeSomeInts(make([]int, 128), true)
	plot.AddPointGroup("SaveAs PDF test", Points, pts)

	tmpDir, err := ioutil.TempDir("", "glot")
	if err != nil {
		t.Errorf("%v", err)
	}
	tmpPlotFile := filepath.Join(tmpDir, "f.pdf")
	defer os.RemoveAll(tmpDir)
	plot.SavePlot(tmpPlotFile)
	plot.Close()

	// Open file, read it into a checksum function and generate checksum.
	fh, err := os.Open(tmpPlotFile)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer fh.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, fh); err != nil {
		t.Errorf("%v", err)
	}
	// Compare expected checksum against actual checksum.
	// If there is a mismatch, test should fail.
	expected := []byte{
		196, 115, 172, 87, 72, 52, 213, 2,
		151, 79, 85, 79, 8, 55, 73, 228,
	}
	if bytes.Compare(expected, hash.Sum(nil)) != 0 {
		t.Errorf("Failed checksum, expected: %x got: %x",
			expected, hash.Sum(nil))
	}
}

func TestGenerateSvgOutput1d1(t *testing.T) {
	var testArgs = PlotArgs{
		Debug:   true,
		Format:  Svg,
		Persist: true,
		Style:   Points,
	}
	plot, err := NewPlot(1, testArgs)
	if err != nil {
		t.Errorf("NewPlot failed creation with: %v", err)
	}
	pts := makeSomeInts(make([]int, 128), true)
	plot.AddPointGroup("SaveAs SVG test", Points, pts)

	tmpDir, err := ioutil.TempDir("", "glot")
	if err != nil {
		t.Errorf("%v", err)
	}
	tmpPlotFile := filepath.Join(tmpDir, "f.svg")
	defer os.RemoveAll(tmpDir)
	plot.SavePlot(tmpPlotFile)
	plot.Close()

	// Open file, read it into a checksum function and generate checksum.
	fh, err := os.Open(tmpPlotFile)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer fh.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, fh); err != nil {
		t.Errorf("%v", err)
	}
	// Compare expected checksum against actual checksum.
	// If there is a mismatch, test should fail.
	expected := []byte{
		94, 214, 211, 163, 55, 159, 236, 46,
		130, 189, 78, 122, 220, 243, 2, 35,
	}
	if bytes.Compare(expected, hash.Sum(nil)) != 0 {
		t.Errorf("Failed checksum, expected: %x got: %x",
			expected, hash.Sum(nil))
	}
}

func TestGeneratePngOutput2d1(t *testing.T) {
	var testArgs = PlotArgs{
		Debug:   true,
		Format:  Png,
		Persist: true,
		Style:   Points,
	}
	plot, err := NewPlot(2, testArgs)
	if err != nil {
		t.Errorf("NewPlot failed creation with: %v", err)
	}
	pts := [][]int{makeSomeInts(make([]int, 32), true),
		makeSomeInts(make([]int, 32), false)}
	plot.AddPointGroup("SaveAs PNG test", Points, pts)

	tmpDir, err := ioutil.TempDir("", "glot")
	if err != nil {
		t.Errorf("%v", err)
	}
	tmpPlotFile := filepath.Join(tmpDir, "f.png")
	defer os.RemoveAll(tmpDir)
	plot.SavePlot(tmpPlotFile)
	plot.Close()

	// Open file, read it into a checksum function and generate checksum.
	fh, err := os.Open(tmpPlotFile)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer fh.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, fh); err != nil {
		t.Errorf("%v", err)
	}
	// Compare expected checksum against actual checksum.
	// If there is a mismatch, test should fail.
	expected := []byte{
		134, 219, 152, 31, 199, 131, 195, 224,
		140, 175, 50, 31, 115, 132, 118, 238,
	}
	if bytes.Compare(expected, hash.Sum(nil)) != 0 {
		t.Errorf("Failed checksum, expected: %x got: %x",
			expected, hash.Sum(nil))
	}
}

func TestGeneratePdfOutput2d1(t *testing.T) {
	var testArgs = PlotArgs{
		Debug:   true,
		Format:  Pdf,
		Persist: true,
		Style:   Points,
	}
	plot, err := NewPlot(2, testArgs)
	if err != nil {
		t.Errorf("NewPlot failed creation with: %v", err)
	}
	pts := [][]int{makeSomeInts(make([]int, 32), true),
		makeSomeInts(make([]int, 32), false)}
	plot.AddPointGroup("SaveAs PDF test", Points, pts)

	tmpDir, err := ioutil.TempDir("", "glot")
	if err != nil {
		t.Errorf("%v", err)
	}
	tmpPlotFile := filepath.Join(tmpDir, "f.pdf")
	defer os.RemoveAll(tmpDir)
	plot.SavePlot(tmpPlotFile)
	plot.Close()

	// Open file, read it into a checksum function and generate checksum.
	fh, err := os.Open(tmpPlotFile)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer fh.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, fh); err != nil {
		t.Errorf("%v", err)
	}
	// Compare expected checksum against actual checksum.
	// If there is a mismatch, test should fail.
	expected := []byte{
		17, 138, 162, 227, 90, 247, 240, 14,
		42, 73, 209, 125, 225, 253, 138, 13,
	}
	if bytes.Compare(expected, hash.Sum(nil)) != 0 {
		t.Errorf("Failed checksum, expected: %x got: %x",
			expected, hash.Sum(nil))
	}
}

func TestGenerateSvgOutput2d1(t *testing.T) {
	var testArgs = PlotArgs{
		Debug:   true,
		Format:  Svg,
		Persist: true,
		Style:   Points,
	}
	plot, err := NewPlot(2, testArgs)
	if err != nil {
		t.Errorf("NewPlot failed creation with: %v", err)
	}
	pts := [][]int{makeSomeInts(make([]int, 32), true),
		makeSomeInts(make([]int, 32), false)}
	plot.AddPointGroup("SaveAs SVG test", Points, pts)

	tmpDir, err := ioutil.TempDir("", "glot")
	if err != nil {
		t.Errorf("%v", err)
	}
	tmpPlotFile := filepath.Join(tmpDir, "f.svg")
	defer os.RemoveAll(tmpDir)
	plot.SavePlot(tmpPlotFile)
	plot.Close()

	// Open file, read it into a checksum function and generate checksum.
	fh, err := os.Open(tmpPlotFile)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer fh.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, fh); err != nil {
		t.Errorf("%v", err)
	}
	// Compare expected checksum against actual checksum.
	// If there is a mismatch, test should fail.
	expected := []byte{
		57, 24, 100, 136, 119, 161, 201, 19,
		64, 147, 209, 92, 18, 72, 21, 36,
	}
	if bytes.Compare(expected, hash.Sum(nil)) != 0 {
		t.Errorf("Failed checksum, expected: %x got: %x",
			expected, hash.Sum(nil))
	}
}

func TestGeneratePngOutput3d1(t *testing.T) {
	var testArgs = PlotArgs{
		Debug:   true,
		Format:  Png,
		Persist: true,
		Style:   Points,
	}
	plot, err := NewPlot(3, testArgs)
	if err != nil {
		t.Errorf("NewPlot failed creation with: %v", err)
	}
	pts := [][]int{
		makeSomeInts(make([]int, 32), true),
		makeSomeInts(make([]int, 32), false),
		makeSomeInts(make([]int, 32), false),
	}
	plot.AddPointGroup("SaveAs PNG test", Points, pts)
	tmpDir, err := ioutil.TempDir("", "glot")
	if err != nil {
		t.Errorf("%v", err)
	}
	tmpPlotFile := filepath.Join(tmpDir, "f.png")
	defer os.RemoveAll(tmpDir)
	plot.SavePlot(tmpPlotFile)
	plot.Close()

	// Open file, read it into a checksum function and generate checksum.
	fh, err := os.Open(tmpPlotFile)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer fh.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, fh); err != nil {
		t.Errorf("%v", err)
	}
	// Compare expected checksum against actual checksum.
	// If there is a mismatch, test should fail.
	expected := []byte{
		218, 234, 117, 20, 238, 2, 137, 102,
		25, 29, 222, 203, 71, 99, 184, 139,
	}
	if bytes.Compare(expected, hash.Sum(nil)) != 0 {
		t.Errorf("Failed checksum, expected: %x got: %x",
			expected, hash.Sum(nil))
	}
}

func TestGeneratePdfOutput3d1(t *testing.T) {
	var testArgs = PlotArgs{
		Debug:   true,
		Format:  Pdf,
		Persist: true,
		Style:   Points,
	}
	plot, err := NewPlot(3, testArgs)
	if err != nil {
		t.Errorf("NewPlot failed creation with: %v", err)
	}
	pts := [][]int{
		makeSomeInts(make([]int, 32), true),
		makeSomeInts(make([]int, 32), false),
		makeSomeInts(make([]int, 32), false),
	}
	plot.AddPointGroup("SaveAs PDF test", Points, pts)

	tmpDir, err := ioutil.TempDir("", "glot")
	if err != nil {
		t.Errorf("%v", err)
	}
	tmpPlotFile := filepath.Join(tmpDir, "f.pdf")
	defer os.RemoveAll(tmpDir)
	plot.SavePlot(tmpPlotFile)
	plot.Close()

	// Open file, read it into a checksum function and generate checksum.
	fh, err := os.Open(tmpPlotFile)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer fh.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, fh); err != nil {
		t.Errorf("%v", err)
	}
	// Compare expected checksum against actual checksum.
	// If there is a mismatch, test should fail.
	expected := []byte{
		113, 70, 64, 149, 3, 210, 177, 245,
		62, 237, 2, 31, 137, 88, 5, 69,
	}
	if bytes.Compare(expected, hash.Sum(nil)) != 0 {
		t.Errorf("Failed checksum, expected: %x got: %x",
			expected, hash.Sum(nil))
	}
}

func TestGenerateSvgOutput3d1(t *testing.T) {
	var testArgs = PlotArgs{
		Debug:   true,
		Format:  Svg,
		Persist: true,
		Style:   Points,
	}
	plot, err := NewPlot(3, testArgs)
	if err != nil {
		t.Errorf("NewPlot failed creation with: %v", err)
	}
	pts := [][]int{
		makeSomeInts(make([]int, 32), true),
		makeSomeInts(make([]int, 32), false),
		makeSomeInts(make([]int, 32), false),
	}
	plot.AddPointGroup("SaveAs SVG test", Points, pts)

	tmpDir, err := ioutil.TempDir("", "glot")
	if err != nil {
		t.Errorf("%v", err)
	}
	tmpPlotFile := filepath.Join(tmpDir, "f.svg")
	defer os.RemoveAll(tmpDir)
	plot.SavePlot(tmpPlotFile)
	plot.Close()

	// Open file, read it into a checksum function and generate checksum.
	fh, err := os.Open(tmpPlotFile)
	if err != nil {
		t.Errorf("%v", err)
	}
	defer fh.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, fh); err != nil {
		t.Errorf("%v", err)
	}
	// Compare expected checksum against actual checksum.
	// If there is a mismatch, test should fail.
	expected := []byte{
		232, 198, 196, 33, 52, 149, 67, 25,
		79, 84, 140, 42, 119, 223, 123, 203,
	}
	if bytes.Compare(expected, hash.Sum(nil)) != 0 {
		t.Errorf("Failed checksum, expected: %x got: %x",
			expected, hash.Sum(nil))
	}
}

func skipOneInts(n []int, acc int) []int {
	if len(n) == 0 {
		return n
	}
	n[0] = acc
	acc += 2
	return skipOneInts(n[1:], acc)
}

func makeSomeInts(n []int, e bool) []int {
	if e {
		skipOneInts(n, 0) // evens, start at 0 increment by 2
	} else {
		skipOneInts(n, 1) // odds, start at 1, increment by 2
	}
	return n
}
