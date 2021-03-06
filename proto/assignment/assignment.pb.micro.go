// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/assignment/assignment.proto

package go_micro_srv_assignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Assignment service

type AssignmentService interface {
	Assign(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type assignmentService struct {
	c    client.Client
	name string
}

func NewAssignmentService(name string, c client.Client) AssignmentService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.assignment"
	}
	return &assignmentService{
		c:    c,
		name: name,
	}
}

func (c *assignmentService) Assign(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Assignment.Assign", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Assignment service

type AssignmentHandler interface {
	Assign(context.Context, *Request, *Response) error
}

func RegisterAssignmentHandler(s server.Server, hdlr AssignmentHandler, opts ...server.HandlerOption) error {
	type assignment interface {
		Assign(ctx context.Context, in *Request, out *Response) error
	}
	type Assignment struct {
		assignment
	}
	h := &assignmentHandler{hdlr}
	return s.Handle(s.NewHandler(&Assignment{h}, opts...))
}

type assignmentHandler struct {
	AssignmentHandler
}

func (h *assignmentHandler) Assign(ctx context.Context, in *Request, out *Response) error {
	return h.AssignmentHandler.Assign(ctx, in, out)
}
