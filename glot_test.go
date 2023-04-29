package glot

import (
	"bytes"
	"crypto/sha1"
	"io"
	"math"
	"math/rand"
	"os"
	"testing"
)

func TestNewPlot(t *testing.T) {
	persist := false
	debug := true
	_, err := NewPlot(0, persist, debug)
	if err == nil {
		t.Error("Expected error when making a 0 dimensional plot.")
	}
}

func TestMinValue(t *testing.T) {
	tests := []struct {
		inputs   []int
		expected int
	}{
		{
			[]int{2, 1, 6, -1, -5}, -5,
		},
		{
			[]int{10, 1000, -1000, 1}, -1000,
		},
		{
			[]int{1, 2, 3, 4, 5}, 1,
		},
	}
	for _, test := range tests {
		result := minValue(test.inputs...)
		if result != test.expected {
			t.Errorf("expected %d, received: %d\n", test.expected, result)

		}
	}
}

func TestPopulateSlice(t *testing.T) {
	expected := []float64{
		0.028195518620623417, 0.9118100887768767,
		1.0880881171536092, 0.8212952160091804,
		-0.33866344995966, 1.4866716553726071,
		-0.7511760243714543, 1.8465685075494818,
		0.6862981484621127, 0.21741348950326284,
	}
	s := make([]float64, 10)
	r := rand.New(rand.NewSource(321)) // Should always expect same results
	populateSlice(s, r)
	for i := range s {
		if !floatsEq(s[i], expected[i]) {
			t.Errorf("expected: %f received: %f", expected[i], s[i])
		}
	}
}

// floatsEq compares two floating point values with precision defined by
// the value of ε.
func floatsEq(a, b float64) bool {
	ε := 0.00000001 // Error
	if math.Abs(a-b) < ε {
		return true
	}
	return false
}

func TestPlotX1(t *testing.T) {
	want := []byte{
		126, 112, 122, 148, 146, 237, 178, 64, 107, 104,
		12, 209, 108, 1, 136, 192, 195, 121, 229, 186,
	}
	p, _ := NewPlot(1, true, true)
	p.debug = true
	p.format = "pdf"
	points := [][]float64{generatePoints(2018, 1000)}
	p.AddPointGroup("testRandomNormal", "", points)
	for k := range p.tmpfiles {
		// Only one entry in map
		if ok, err := confirmCheckSum(k, want); err != nil {
			t.Errorf("%v", err)
		} else if !ok {
			t.Errorf("%v", "Failed to checksum generated data.")
		}
	}
}

func TestPlotXY1(t *testing.T) {
	want := []byte{
		59, 73, 90, 10, 15, 97, 254, 213, 151, 75,
		236, 203, 106, 196, 63, 67, 52, 94, 111, 33,
	}
	p, _ := NewPlot(2, true, true)
	p.debug = true
	p.format = "pdf"
	points := [][]float64{
		generatePoints(-1, 1000),
		generatePoints(1, 1000),
	}
	p.AddPointGroup("testRandomNormal", "", points)
	for k := range p.tmpfiles {
		// Only one entry in map
		if ok, err := confirmCheckSum(k, want); err != nil {
			t.Errorf("%v", err)
		} else if !ok {
			t.Errorf("%v", "Failed to checksum generated data.")
		}
	}
}

func TestPlotXYZ1(t *testing.T) {
	want := []byte{
		189, 141, 23, 248, 204, 159, 175, 78, 178, 209,
		162, 43, 8, 214, 143, 161, 186, 175, 247, 178,
	}
	p, _ := NewPlot(3, true, true)
	p.debug = true
	p.format = "pdf"
	points := [][]float64{
		generatePoints(-1, 1000),
		generatePoints(1, 1000),
		generatePoints(2, 1000),
	}
	p.AddPointGroup("testRandomNormal", "", points)
	for k := range p.tmpfiles {
		// Only one entry in map
		if ok, err := confirmCheckSum(k, want); err != nil {
			t.Errorf("%v", err)
		} else if !ok {
			t.Errorf("%v", "Failed to checksum generated data.")
		}
	}
}

// confirmCheckSum is used for testing expected checksum of file
// that we generated during the test against actual checksum.
// If these two don't agree, something has gone wrong.
func confirmCheckSum(fn string, expected []byte) (bool, error) {
	f, err := os.Open(fn)
	if err != nil {
		return false, err
	}
	defer f.Close()
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return false, err
	}
	if bytes.Compare(h.Sum(nil), expected) != 0 {
		return false, nil
	}
	return true, nil
}

func generatePoints(seed int64, length int) []float64 {
	if length == 0 {
		return []float64{}
	}
	pts := make([]float64, length)
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < length; i++ {
		pts[i] = r.NormFloat64()
	}
	return pts
}

func populateSlice(s []float64, rs *rand.Rand) ([]float64, *rand.Rand) {
	if len(s) == 0 {
		return s, rs
	}
	s[0] = rs.NormFloat64()
	return populateSlice(s[1:], rs)
}
