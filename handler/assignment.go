package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	assignment "assignment/proto/assignment"
)

type Assignment struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Assignment) Call(ctx context.Context, req *assignment.Request, rsp *assignment.Response) error {
	log.Log("Received Assignment.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Assignment) Stream(ctx context.Context, req *assignment.StreamingRequest, stream assignment.Assignment_StreamStream) error {
	log.Logf("Received Assignment.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&assignment.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Assignment) PingPong(ctx context.Context, stream assignment.Assignment_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&assignment.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
