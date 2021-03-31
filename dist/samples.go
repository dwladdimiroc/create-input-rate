package dist

type Sample struct {
	Time  int `csv:"time"`
	Value int `csv:"value"`
}

func stdSamples(samples []Sample) []Sample {
	for i := 0; i < len(samples); i++ {
		samples[i].Value = samples[i].Value / 5
	}
	return samples
}
