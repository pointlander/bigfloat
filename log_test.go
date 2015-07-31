package floatutils_test

import (
	"math"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ALTree/floatutils"
)

func TestLog(t *testing.T) {
	for _, test := range []struct {
		x    string
		prec uint
		want string
	}{
		// 80 decimal digits are enough to give us 250 binary digits when parsed by the Parse function
		{"0.5", 250, "-0.6931471805599453094172321214581765680755001343602552541206800094933936219696947"},
		{"0.25", 250, "-1.3862943611198906188344642429163531361510002687205105082413600189867872439393894"},
		{"0.0125", 250, "-4.3820266346738816122696878190588939118276018917095387383953679294477534755864366"},

		{"2", 250, "0.6931471805599453094172321214581765680755001343602552541206800094933936219696947"},
		{"10", 250, "2.3025850929940456840179914546843642076011014886287729760333279009675726096773524"},
		{"512", 250, "6.2383246250395077847550890931235891126795012092422972870861200854405425977272524"},
		{"42e42", 250, "100.4462435240332870346734711985671787216063379441733308877145363337065173843891771"},
	} {
		want := new(big.Float).SetPrec(test.prec)
		want.Parse(test.want, 10)

		x := new(big.Float).SetPrec(test.prec)
		x.Parse(test.x, 10)

		z := floatutils.Log(x)

		// test if precision is correctly set
		if z.Prec() != test.prec {
			t.Errorf("Log(%v): got %d prec, want %d prec", x, z.Prec(), test.prec)
		}

		// test returned value
		if !compareFloats(want, z, test.prec, t) {
			t.Errorf("Log(%v): error is too big.\nwant = %.100e\ngot  = %.100e\n", x, z, want)
		}
	}
}

func TestLog32Small(t *testing.T) {
	for i := 0; i < 5e3; i++ {
		r := rand.Float32()*1e1 + 1
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floatutils.Log(x).Float32()
		want := math.Log(float64(r))
		if z != float32(want) || acc != big.Exact {
			t.Errorf("Log(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestLog32Big(t *testing.T) {
	for i := 0; i < 5e3; i++ {
		r := rand.Float32()*1e6 + 1e3
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floatutils.Log(x).Float32()
		want := math.Log(float64(r))
		if z != float32(want) || acc != big.Exact {
			t.Errorf("Log(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

// for the Log function we don't require complete compatibily
// with native float64 arithmetic. Let's settle for an error
// smaller than 1e-14

func TestLog64Small(t *testing.T) {
	for i := 0; i < 5e3; i++ {
		r := rand.Float64()*1e1 + 1
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floatutils.Log(x).Float64()
		want := math.Log(r)
		if math.Abs(z-want) > 1e-14 || acc != big.Exact {
			t.Errorf("Log(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestLog64Big(t *testing.T) {
	for i := 0; i < 5e3; i++ {
		r := rand.Float64()*1e6 + 1e3
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floatutils.Log(x).Float64()
		want := math.Log(r)
		if math.Abs(z-want) > 1e-14 || acc != big.Exact {
			t.Errorf("Log(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestLogSpecialValues(t *testing.T) {
	for i, f := range []float64{
		+0.0,
		-0.0,
		math.Inf(+1),
	} {
		x := big.NewFloat(f).SetPrec(53)
		z, acc := floatutils.Log(x).Float64()
		want := math.Log(f)
		if z != want || acc != big.Exact {
			t.Errorf("%d) Log(%g) =\n got %b (%s);\nwant %b (Exact)", i, f, z, acc, want)
		}
	}
}