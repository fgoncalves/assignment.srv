package data

import (
	experiments "assignment/proto/assignment"
	"context"
)

type Service struct {
	Service experiments.ExperimentsService
}

func (s Service) GetExperiment(name string) (*experiments.ExperimentResponse, error) {
	request := experiments.ExperimentRequest{
		Name: name,
	}

	service := s.Service
	return service.Get(context.Background(), &request)
}
