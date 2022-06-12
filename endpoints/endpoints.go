package endpoints

import (
	"context"
	"github.com/stkr89/modelsvc/types"

	"github.com/go-kit/kit/endpoint"
	"github.com/stkr89/modelsvc/service"
)

type Endpoints struct {
	Create endpoint.Endpoint
	Get    endpoint.Endpoint
	List   endpoint.Endpoint
	Update endpoint.Endpoint
	Delete endpoint.Endpoint
}

func MakeEndpoints(s service.ModelService) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		List:   makeListEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

func makeDeleteEndpoint(s service.ModelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*types.DeleteRequest)
		return nil, s.Delete(ctx, req)
	}
}

func makeUpdateEndpoint(s service.ModelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*types.UpdateRequest)
		return s.Update(ctx, req)
	}
}

func makeListEndpoint(s service.ModelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*types.ListRequest)
		return s.List(ctx, req)
	}
}

func makeGetEndpoint(s service.ModelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*types.GetRequest)
		return s.Get(ctx, req)
	}
}

func makeCreateEndpoint(s service.ModelService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*types.CreateRequest)
		return s.Create(ctx, req)
	}
}
