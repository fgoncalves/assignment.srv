package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	"assignment/domain"

	assignment "assignment/proto/assignment"
)

type Assignment struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Assignment) Assign(ctx context.Context, req *assignment.Request, rsp *assignment.Response) error {
	log.Log("Received Assignment.Call request")

	// TODO: We would need to find the experiment I suppose
	experiment := domain.Experiment{
		Name: req.ExperimentName,
		Variants: []domain.Variant{
			domain.Variant{
				Name:       "control",
				Percentage: 0.5,
			},
			domain.Variant{
				Name:       "test",
				Percentage: 0.5,
			},
		},
	}

	user := domain.User{
		Id: req.UserId,
	}

	variant := domain.RandomAssign(experiment, user)
	rsp.Variant = variant.Name
	return nil
}
