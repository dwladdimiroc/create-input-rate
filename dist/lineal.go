package dist

import "math"

func Lineal(n int) []Sample {
	var samples = make([]Sample, n)
	var peak = 5000.0
	var numbersPeaks = 5

	for c := 0; c < numbersPeaks; c++ {
		var k int
		for i := 0; i < n/numbersPeaks; i++ {
			k = i + (n/5)*c
			samples[k].Time = k
			samples[k].Value = fLineal(i, n, peak)
		}
	}

	return stdSamples(samples)
}

func fLineal(x, n int, max float64) int {
	time := float64(n) / 5.0
	y := (max / time) * float64(x)
	return int(math.Floor(y))
}
