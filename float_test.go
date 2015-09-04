package timetables

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFloatMul(t *testing.T) {
	f := NewFloat(12.0)

	res := f.Mul(NewFloat(2.0))

	expected := big.NewFloat(12.0)
	expected = expected.Mul(expected, big.NewFloat(2.0))

	if !reflect.DeepEqual(expected, res.Float) {
		t.Logf("Expected float to equal\n%v\n\tgot:\n%v\n", expected, res.Float)
		t.Fail()
	}
}

func TestFloatDiv(t *testing.T) {
	f := NewFloat(12.0)

	res := f.Div(NewFloat(2.0))

	expected := big.NewFloat(12.0)
	expected = expected.Quo(expected, big.NewFloat(2.0))

	if !reflect.DeepEqual(expected, res.Float) {
		t.Logf("Expected float to equal\n%v\n\tgot:\n%v\n", expected, res.Float)
		t.Fail()
	}
}

func TestFloatAdd(t *testing.T) {
	f := NewFloat(12.0)

	res := f.Add(NewFloat(2.0))

	expected := big.NewFloat(12.0)
	expected = expected.Add(expected, big.NewFloat(2.0))

	if !reflect.DeepEqual(expected, res.Float) {
		t.Logf("Expected float to equal\n%v\n\tgot:\n%v\n", expected, res.Float)
		t.Fail()
	}
}

func TestFloatSub(t *testing.T) {
	f := NewFloat(12.0)

	res := f.Sub(NewFloat(2.0))

	expected := big.NewFloat(12.0)
	expected = expected.Sub(expected, big.NewFloat(2.0))

	if !reflect.DeepEqual(expected, res.Float) {
		t.Logf("Expected float to equal\n%v\n\tgot:\n%v\n", expected, res.Float)
		t.Fail()
	}
}

func TestFloatCmp(t *testing.T) {
	f := NewFloat(12)

	tests := []struct {
		in  *Float
		out int
	}{
		{
			NewFloat(12),
			0,
		},
		{
			NewFloat(12.01),
			-1,
		},
		{
			NewFloat(11),
			1,
		},
	}
	for _, test := range tests {
		res := f.Cmp(test.in)

		if test.out != res {
			t.Logf("Expected result to equal %d, got %d\n", test.out, res)
			t.Fail()
		}
	}
}
