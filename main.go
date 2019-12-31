package main

import (
	"assignment/handler"
	"assignment/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

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
	micro.RegisterSubscriber("experiments", service.Server(), new(subscriber.Assigment))

	// Register Function as Subscriber
	micro.RegisterSubscriber("experiments", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
