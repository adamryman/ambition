// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: b8d375e642
// Version Date: Fri Aug 4 23:55:38 UTC 2017

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "github.com/myambition/ambition/services/rello/rello-service"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	CheckListWebhookEndpoint endpoint.Endpoint
	EmptyRPCEndpoint         endpoint.Endpoint
}

// Endpoints

func (e Endpoints) CheckListWebhook(ctx context.Context, in *pb.ChecklistUpdate) (*pb.Empty, error) {
	response, err := e.CheckListWebhookEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Empty), nil
}

func (e Endpoints) EmptyRPC(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	response, err := e.EmptyRPCEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Empty), nil
}

// Make Endpoints

func MakeCheckListWebhookEndpoint(s pb.RelloServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.ChecklistUpdate)
		v, err := s.CheckListWebhook(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeEmptyRPCEndpoint(s pb.RelloServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.Empty)
		v, err := s.EmptyRPC(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"CheckListWebhook": struct{}{},
		"EmptyRPC":         struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "CheckListWebhook" {
			e.CheckListWebhookEndpoint = middleware(e.CheckListWebhookEndpoint)
		}
		if inc == "EmptyRPC" {
			e.EmptyRPCEndpoint = middleware(e.EmptyRPCEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"CheckListWebhook": struct{}{},
		"EmptyRPC":         struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "CheckListWebhook" {
			e.CheckListWebhookEndpoint = middleware("CheckListWebhook", e.CheckListWebhookEndpoint)
		}
		if inc == "EmptyRPC" {
			e.EmptyRPCEndpoint = middleware("EmptyRPC", e.EmptyRPCEndpoint)
		}
	}
}
