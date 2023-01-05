package math

import "testing"

func TestAverage(t *testing.T) {
	var v float64;
	v = Average([]float64{1,2});
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v);
	}
}

func TestMin(t *testing.T) {
	v := Min([]float64{1,2});
	if v != 1 {
		t.Error("Expected 1, got ", v);
	}
}

func TestMax(t *testing.T) {
	v := Max([]float64{1,2});
	if v != 2 {
		t.Error("Expected 2, got ", v);
	}
}