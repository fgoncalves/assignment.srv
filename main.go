package main

import (
	"assignment/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"

	assignment "assignment/proto/assignment"

	"assignment/data"
	"assignment/domain"
)

func newAssignmentHandler() handler.Assignment {
	experimentsService := grpc.NewService()
	experimentsService.Init()

	cl := assignment.NewExperimentsService("go.micro.srv.experiments", experimentsService.Client())

	service := data.Service{
		Service: cl,
	}

	repository := domain.ExperimentsRepository{
		Service: &service,
	}

	return handler.Assignment{
		ExperimentsRepository: &repository,
	}
}

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.assignment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	handler := newAssignmentHandler()

	// Register Handler
	assignment.RegisterAssignmentHandler(service.Server(), &handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
