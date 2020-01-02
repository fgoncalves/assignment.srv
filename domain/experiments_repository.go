package domain

import (
	"assignment/data"
)

type ExperimentsRepository struct {
	Service *data.Service
}

func (r *ExperimentsRepository) GetExperiment(name string) (Experiment, error) {
	var experiment Experiment

	service := r.Service

	resp, err := service.GetExperiment(name)

	if err != nil {
		return experiment, err
	}

	variants := make([]Variant, len(resp.Variants))
	for i, variant := range resp.Variants {
		variants[i] = Variant{
			Name:       variant.Name,
			Percentage: variant.Percentage,
		}
	}

	experiment = Experiment{
		Name:     resp.Name,
		Variants: variants,
	}

	return experiment, nil
}
