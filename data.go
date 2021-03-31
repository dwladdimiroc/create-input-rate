package main

import (
	"fmt"
	"github.com/dwladdimiroc/storm-input-data/dist"
	"github.com/dwladdimiroc/storm-input-data/files"
)

func main() {
	var lengthSamples = 4000

	samplesNorm := dist.Normal(lengthSamples)
	if err := files.WriteSamples("norm", samplesNorm); err != nil {
		fmt.Printf("error csv write: %v\n", err)
	}

	samplesNNorm := dist.NNormal(3, lengthSamples)
	if err := files.WriteSamples("n-norm", samplesNNorm); err != nil {
		fmt.Printf("error csv write: %v\n", err)
	}

	samplesLineal := dist.Lineal(lengthSamples)
	if err := files.WriteSamples("lineal", samplesLineal); err != nil {
		fmt.Printf("error csv write: %v\n", err)
	}

	samplesLogarithm := dist.Logarithm(lengthSamples)
	if err := files.WriteSamples("logarithm", samplesLogarithm); err != nil {
		fmt.Printf("error csv write: %v\n", err)
	}

	samplesExponential := dist.Exponential(lengthSamples)
	if err := files.WriteSamples("exponential", samplesExponential); err != nil {
		fmt.Printf("error csv write: %v\n", err)
	}

	samplesPoisson := dist.Poisson(lengthSamples)
	if err := files.WriteSamples("poisson", samplesPoisson); err != nil {
		fmt.Printf("error csv write: %v\n", err)
	}
}
