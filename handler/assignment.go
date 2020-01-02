package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	"assignment/domain"

	assignment "assignment/proto/assignment"
)

type Assignment struct {
	ExperimentsRepository *domain.ExperimentsRepository
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Assignment) Assign(ctx context.Context, req *assignment.Request, rsp *assignment.Response) error {
	experiment, err := e.ExperimentsRepository.GetExperiment(req.ExperimentName)
	if err != nil {
		log.Log("Received error " + err.Error())
		return err
	}

	user := domain.User{
		Id: req.UserId,
	}

	variant := domain.RandomAssign(experiment, user)
	rsp.Variant = variant.Name
	return nil
}
