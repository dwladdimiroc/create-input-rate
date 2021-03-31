package dist

import "math"

func Logarithm(n int) []Sample {
	var peak = 1000.0

	var samples = make([]Sample, n)
	for i := 0; i < n; i++ {
		samples[i].Time = i
		samples[i].Value = int(math.Floor(math.Log10(float64(i)) * peak))
	}

	return stdSamples(samples)
}
