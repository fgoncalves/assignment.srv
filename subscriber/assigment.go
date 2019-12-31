package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	assignment "assignment/proto/assignment"
)

type Assigment struct{}

func (e *Assigment) Handle(ctx context.Context, msg *assignment.ExperimentEvent) error {
	log.Log("Handler Received new experiment: ", msg.Id)
	return nil
}

func Handler(ctx context.Context, msg *assignment.ExperimentEvent) error {
	log.Log("Function Received message: ", msg.Id)
	return nil
}
