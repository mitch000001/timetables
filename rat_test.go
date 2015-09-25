package timetables

import (
	"math"
	"math/big"
	"reflect"
	"testing"
)

func TestNewRat(t *testing.T) {
	f := 0.86

	rat := NewRat(f)

	expectedRat := &Rat{
		new(big.Rat).SetFloat64(f),
	}

	if !reflect.DeepEqual(expectedRat, rat) {
		t.Logf("Expected NewRat to return\n%+v\n\tgot\n%+v\n", expectedRat, rat)
		t.Fail()
	}

	// f not finit
	f = math.Inf(0)

	rat = NewRat(f)

	if rat != nil {
		t.Logf("Expected rat to be nil, was not: %+v\n", rat)
		t.Fail()
	}
}

func TestRatAdd(t *testing.T) {
	rat := NewRat(3.0)

	res := rat.Add(NewRat(2))

	bigRat := big.NewRat(3, 1)

	expected := new(big.Rat).Add(bigRat, big.NewRat(2, 1))

	if !reflect.DeepEqual(expected, res.Rat) {
		t.Logf("Expected Add to return\n%+v\n\tgot\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestRatSub(t *testing.T) {
	rat := NewRat(3.0)

	res := rat.Sub(NewRat(2))

	bigRat := big.NewRat(3, 1)

	expected := new(big.Rat).Sub(bigRat, big.NewRat(2, 1))

	if !reflect.DeepEqual(expected, res.Rat) {
		t.Logf("Expected Add to return\n%+v\n\tgot\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestRatMul(t *testing.T) {
	rat := NewRat(3.0)

	res := rat.Mul(NewRat(2))

	bigRat := big.NewRat(3, 1)

	expected := new(big.Rat).Mul(bigRat, big.NewRat(2, 1))

	if !reflect.DeepEqual(expected, res.Rat) {
		t.Logf("Expected Add to return\n%+v\n\tgot\n%+v\n", expected, res)
		t.Fail()
	}
}

func TestRatDiv(t *testing.T) {
	rat := NewRat(3.0)

	res := rat.Div(NewRat(2))

	bigRat := big.NewRat(3, 1)

	expected := new(big.Rat).Quo(bigRat, big.NewRat(2, 1))

	if !reflect.DeepEqual(expected, res.Rat) {
		t.Logf("Expected Add to return\n%+v\n\tgot\n%+v\n", expected, res)
		t.Fail()
	}

	// div by zero
	defer func() {
		if r := recover(); r == nil {
			t.Logf("Expected error, got nil\n")
			t.Fail()
		}
	}()

	res = rat.Div(NewRat(0))

	if res != nil {
		t.Logf("Expected Div to throw panic\n")
		t.Fail()
	}
}

func TestRatCmp(t *testing.T) {
	f := NewRat(12)

	tests := []struct {
		in  *Rat
		out int
	}{
		{
			NewRat(12),
			0,
		},
		{
			NewRat(12.01),
			-1,
		},
		{
			NewRat(11),
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
