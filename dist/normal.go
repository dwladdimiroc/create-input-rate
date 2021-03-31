package dist

import (
	"gonum.org/v1/gonum/stat/distuv"
)

func Normal(n int) []Sample {
	constInput := 10000.0

	// use the defined variable
	norm := distuv.UnitNormal

	//norm.Prob()

	var samples = make([]Sample, n)
	for i := 0; i < n; i++ {
		samples[i].Time = i
		sample := (7 * float64(i) / float64(n)) - 3.5
		samples[i].Value = int(norm.Prob(sample) * constInput)
	}

	return stdSamples(samples)
}
