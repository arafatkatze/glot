package glot

import "testing"
import "fmt"

func TestAverage(t *testing.T) {
 fmt.Println("sajjfs")
  var v float64
  v = Average([]float64{1,2})
  if v != 1.5 {
    t.Error("Expected 1.5, got ", v)
  }
}
