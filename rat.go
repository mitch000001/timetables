package timetables

import "math/big"

// NewRat returns a new Rat for the provided float. It behaves similar as big.Rat.SetFloat64
// in that it returns nil if f is not finite
func NewRat(f float64) *Rat {
	r := new(big.Rat)
	r = r.SetFloat64(f)
	if r == nil {
		return nil
	}
	return &Rat{r}
}

type Rat struct {
	*big.Rat
}

func (r *Rat) Add(x *Rat) *Rat {
	z := new(big.Rat)
	res := z.Add(r.Rat, x.Rat)
	return &Rat{res}
}

func (r *Rat) Sub(x *Rat) *Rat {
	z := new(big.Rat)
	res := z.Sub(r.Rat, x.Rat)
	return &Rat{res}
}

func (r *Rat) Mul(x *Rat) *Rat {
	z := new(big.Rat)
	res := z.Mul(r.Rat, x.Rat)
	return &Rat{res}
}

func (r *Rat) Div(x *Rat) *Rat {
	z := new(big.Rat)
	res := z.Quo(r.Rat, x.Rat)
	return &Rat{res}
}

func (r *Rat) Cmp(x *Rat) int {
	return r.Rat.Cmp(x.Rat)
}
