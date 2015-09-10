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

func TestFloatMarshalBinary(t *testing.T) {
	float := NewFloat(8)

	floatBytes, err := float.MarshalBinary()

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	expectedBytes := []byte(float.Text('b', 53))

	if !reflect.DeepEqual(expectedBytes, floatBytes) {
		t.Logf("Expected marshaled value to equal\n%q\n\tgot\n%q\n", expectedBytes, floatBytes)
		t.Fail()
	}
}

func TestFloatUnmarhsalBinary(t *testing.T) {
	floatBytes := []byte(NewFloat(8).Text('b', 53))

	float := &Float{}
	err := float.UnmarshalBinary(floatBytes)

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	expectedFloat := NewFloat(8)

	if expectedFloat.Cmp(float) != 0 {
		t.Logf("Expected unmarshaled value to equal\n%q\n\tgot\n%q\n", expectedFloat.Text('f', 53), float.Text('f', 53))
		t.Fail()
	}
}
