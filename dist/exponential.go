package dist

import "math"

func Exponential(n int) []Sample {
	var l = 0.0008
	var k = 400.0

	var samples = make([]Sample, n)
	for i := 0; i < n; i++ {
		samples[i].Time = i
		samples[i].Value = int(math.Floor(k * math.Exp(l*float64(i))))
		//samples[i].Value = int(math.Floor(math.Exp(float64(i)) * peak))
	}

	return stdSamples(samples)
}
