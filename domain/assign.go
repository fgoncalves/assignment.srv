package domain

import (
	"math/rand"
	"sort"

	"gonum.org/v1/gonum/floats"
)

func weightedIndex(weights []float64, seed int64) int {
	cdf := make([]float64, len(weights))
	floats.CumSum(cdf, weights)

	rand.Seed(seed)
	randomWeight := rand.Float64() * cdf[len(cdf)-1]

	return sort.Search(len(cdf), func(i int) bool {
		return cdf[i] > randomWeight
	})
}

// Probably needs experiments here
func RandomAssign(experiment Experiment, user User) Variant {
	variants := experiment.Variants
	weights := make([]float64, len(variants))

	for i := range variants {
		weights[i] = variants[i].Percentage
	}

	seed := SeedFor(experiment, user)
	return variants[weightedIndex(weights, seed)]
}
