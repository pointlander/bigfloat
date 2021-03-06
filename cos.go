package bigfloat

import (
	"math/big"
)

// Cos compute the cos of the value
// https://people.math.sc.edu/girardi/m142/handouts/10sTaylorPolySeries.pdf
func Cos(z *big.Float) *big.Float {
	prec := z.Prec()
	one := big.NewFloat(1).SetPrec(prec)
	sign := true
	zz := big.NewFloat(0).SetPrec(prec)
	zz.Set(z)
	if zz.Cmp(big.NewFloat(0).SetPrec(0)) < 0 {
		zz.Neg(zz)
	}
	a, d, i := big.NewFloat(0).SetPrec(prec), big.NewFloat(2).SetPrec(prec),
		big.NewFloat(2).SetPrec(prec)
	for {
		b := Pow(zz, i)
		b = b.Quo(b, d)
		cp := big.NewFloat(0).SetPrec(prec).Set(a)
		if sign {
			a.Sub(a, b)
		} else {
			a.Add(a, b)
		}
		sign = !sign
		if cp.Cmp(a) == 0 {
			break
		}
		for j := 0; j < 2; j++ {
			i.Add(i, one)
			d.Mul(d, i)
		}
	}
	return a.Add(one, a)
}
