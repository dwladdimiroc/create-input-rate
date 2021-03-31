package dist

import (
	"gonum.org/v1/gonum/stat/distuv"
	"math"
	"math/rand"
)

func Poisson(n int) []Sample {
	var peak = 30000.0

	poisson := distuv.Poisson{
		Lambda: 5,
	}

	var numSamples = 10
	var probPoisson = make([]float64, numSamples)
	for x := 0; x < numSamples; x++ {
		probPoisson[x] = poisson.Prob(float64(x))
	}

	var numProb = 360
	probPoissonExtrapolate := extrapolate(probPoisson, numProb)

	var parts = n / numProb
	var samples = make([]Sample, n)
	for i := 0; i < parts; i++ {
		k := rand.Float64() * peak
		for j := 0; j < numProb; j++ {
			samples[i*numProb+j].Time = i*numProb + j
			samples[i*numProb+j].Value = int(math.Floor(probPoissonExtrapolate[j]*k)) + 200
		}
	}

	samples = joinFunctions(samples, parts)
	for i := len(samples) - 100; i < len(samples); i++ {
		samples[i].Time = i
		samples[i].Value = 200
	}

	return stdSamples(samples)
}

func extrapolate(data []float64, samples int) []float64 {
	var newData = make([]float64, samples)

	c := samples / len(data)
	for j := 0; j < len(data)-1; j++ {
		newData[j*c] = data[j]
		for i := 1; i < c; i++ {
			y2 := data[j+1]
			x2 := float64(c)
			y1 := data[j]
			x1 := float64(0)
			x := float64(i)
			y := y1 + (((y2 - y1) / (x2 - x1)) * (x - x1))
			newData[j*c+i] = y
			//fmt.Printf("numProb %v index %v\n",newData[j*c+i], j*c+i)
		}
	}

	return newData
}

func joinFunctions(samples []Sample, parts int) []Sample {
	var n = len(samples) / parts

	for i := 1; i < parts; i++ {
		nPart := n * i
		x3 := nPart - n/4
		y3 := samples[x3].Value
		x2 := nPart + n/4
		y2 := samples[x2].Value
		x1 := nPart
		y1 := math.Min(float64(y3), float64(y2)) / 2
		a, b, c := calculateParabola(float64(x1), y1, float64(x2), float64(y2), float64(x3), float64(y3))
		var par = Parabola{A: a, B: b, C: c}
		samples = joinParabola(par, n, nPart, samples)
	}

	return samples
}

type Parabola struct {
	A, B, C float64
}

func (p Parabola) f(x float64) float64 {
	return p.A*math.Pow(x, 2.0) + p.B*x + p.C
}

func calculateParabola(x1, y1, x2, y2, x3, y3 float64) (float64, float64, float64) {
	denominator := (x1 - x2) * (x1 - x3) * (x2 - x3)
	A := (x3*(y2-y1) + x2*(y1-y3) + x1*(y3-y2)) / denominator
	B := (x3*x3*(y1-y2) + x2*x2*(y3-y1) + x1*x1*(y2-y3)) / denominator
	C := (x2*x3*(x2-x3)*y1 + x3*x1*(x3-x1)*y2 + x1*x2*(x1-x2)*y3) / denominator

	return A, B, C
}

func joinParabola(par Parabola, n int, posN int, samples []Sample) []Sample {
	for i := 0; i < n/2; i++ {
		index := (posN - (n / 4)) + i
		y := par.f(float64(index))
		samples[index].Value = int(y)
	}
	return samples
}
