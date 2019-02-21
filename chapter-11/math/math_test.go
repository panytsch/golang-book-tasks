package math

import "testing"

type testPair struct {
	values  []float64
	average float64
	max     float64
	min     float64
}

var testsValues = []testPair{
	{[]float64{1, 2}, 1.5, 2, 1},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 1, 1},
	{[]float64{-1, 1}, 0, 1, -1},
	{[]float64{}, 0, 0, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range testsValues {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range testsValues {
		v := Max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range testsValues {
		v := Min(pair.values)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}
