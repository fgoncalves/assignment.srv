package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	assignment "assignment/proto/assignment"
)

type Assignment struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Assignment) Assign(ctx context.Context, req *assignment.Request, rsp *assignment.Response) error {
	log.Log("Received Assignment.Call request")
	rsp.Variant = "Hello " + req.UserId + req.ExperimentName
	return nil
}