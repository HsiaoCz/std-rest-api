package wrst

import "testing"

type Value struct {
	input  Input
	output int
}

type Input struct {
	x int
	y int
}

func TestCalculateValues(t *testing.T) {
	// t.Fatal("not")
	v := Value{
		input: Input{
			x: 10,
			y: 12,
		},
		output: 22,
	}
	got := CalculateValue(v.input.x, v.input.y)
	if got != v.output {
		t.Error("the CalculateValue function is bad")
		return
	}
}
