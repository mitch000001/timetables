package timetables

import "math/big"

func NewFloat(in float64) *Float {
	return &Float{big.NewFloat(in)}
}

type Float struct {
	*big.Float
}

func (f *Float) Mul(x *Float) *Float {
	res := new(big.Float).Mul(f.Float, x.Float)
	return &Float{res}
}

func (f *Float) Div(x *Float) *Float {
	res := new(big.Float).Quo(f.Float, x.Float)
	return &Float{res}
}

func (f *Float) Add(x *Float) *Float {
	res := new(big.Float).Add(f.Float, x.Float)
	return &Float{res}
}

func (f *Float) Sub(x *Float) *Float {
	res := new(big.Float).Sub(f.Float, x.Float)
	return &Float{res}
}
