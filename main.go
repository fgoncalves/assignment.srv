package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"assignment/handler"
	"assignment/subscriber"

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

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.assignment", service.Server(), new(subscriber.Assignment))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.assignment", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
