package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"assignment/handler"

	assignment "assignment/proto/assignment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.assignment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	assignment.RegisterAssignmentHandler(service.Server(), new(handler.Assignment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
