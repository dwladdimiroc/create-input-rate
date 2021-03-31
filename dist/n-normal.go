package dist

import "gonum.org/v1/gonum/stat/distuv"

func NNormal(m, n int) []Sample {
	constInput := 10000.0

	// use the defined variable
	norm := distuv.UnitNormal

	//norm.Prob()

	var samples = make([]Sample, n)
	for c := 0; c < m; c++ {
		var k int
		for i := 0; i < n/m; i++ {
			k = i + (n/m)*c
			samples[k].Time = k
			sample := (7 * float64(i) / float64(n/m)) - 3.5
			samples[k].Value = int(norm.Prob(sample) * (constInput / float64(c+1)))
		}
	}

	return stdSamples(samples)
}
