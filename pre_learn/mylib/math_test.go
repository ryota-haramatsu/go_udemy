package mylib

import "testing"

// go test
func TestAverage(t *testing.T) {
	t.Skip("Done")
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
 