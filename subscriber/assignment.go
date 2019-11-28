package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	assignment "assignment/proto/assignment"
)

type Assignment struct{}

func (e *Assignment) Handle(ctx context.Context, msg *assignment.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *assignment.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
