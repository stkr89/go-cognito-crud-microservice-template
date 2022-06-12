package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/stkr89/modelsvc/middleware"
	v1 "github.com/stkr89/modelsvc/pb"
	"github.com/stkr89/modelsvc/types"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/stkr89/modelsvc/endpoints"
)

type gRPCServer struct {
	create gt.Handler
	get    gt.Handler
	list   gt.Handler
	update gt.Handler
	delete gt.Handler
	v1.UnimplementedModelSvcServer
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.Endpoints) v1.ModelSvcServer {
	return &gRPCServer{
		create: gt.NewServer(
			endpoint.Chain(
				middleware.ValidateCreateInput(),
				middleware.ConformCreateInput(),
			)(endpoints.Create),
			decodeCreateGRPCRequest,
			encodeCreateGRPCResponse,
		),
		get: gt.NewServer(
			endpoint.Chain(
				middleware.ValidateGetInput(),
				middleware.ConformGetInput(),
			)(endpoints.Get),
			decodeGetGRPCRequest,
			encodeGetGRPCResponse,
		),
		list: gt.NewServer(
			endpoint.Chain(
				middleware.ValidateListInput(),
				middleware.ConformListInput(),
			)(endpoints.List),
			decodeListGRPCRequest,
			encodeListGRPCResponse,
		),
		update: gt.NewServer(
			endpoint.Chain(
				middleware.ValidateUpdateInput(),
				middleware.ConformUpdateInput(),
			)(endpoints.Update),
			decodeUpdateGRPCRequest,
			encodeUpdateGRPCResponse,
		),
		delete: gt.NewServer(
			endpoint.Chain(
				middleware.ValidateDeleteInput(),
				middleware.ConformDeleteInput(),
			)(endpoints.Delete),
			decodeDeleteGRPCRequest,
			encodeDeleteGRPCResponse,
		),
	}
}

func (s gRPCServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	_, resp, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*v1.CreateResponse), nil
}

func (s gRPCServer) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	_, resp, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*v1.GetResponse), nil
}

func (s gRPCServer) List(ctx context.Context, req *v1.ListRequest) (*v1.ListResponse, error) {
	_, resp, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*v1.ListResponse), nil
}

func (s gRPCServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	_, resp, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*v1.UpdateResponse), nil
}

func (s gRPCServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	_, resp, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*v1.DeleteResponse), nil
}

func decodeDeleteGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	reqpb := request.(*v1.DeleteRequest)
	b, err := json.Marshal(reqpb)
	if err != nil {
		return nil, err
	}

	var req types.DeleteRequest
	err = json.Unmarshal(b, &req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func encodeDeleteGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	respb := response.(*types.DeleteRequest)
	b, err := json.Marshal(respb)
	if err != nil {
		return nil, err
	}

	var resp v1.DeleteResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func decodeUpdateGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	reqpb := request.(*v1.UpdateRequest)
	b, err := json.Marshal(reqpb)
	if err != nil {
		return nil, err
	}

	var req types.UpdateRequest
	err = json.Unmarshal(b, &req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func encodeUpdateGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	respb := response.(*types.UpdateResponse)
	b, err := json.Marshal(respb)
	if err != nil {
		return nil, err
	}

	var resp v1.UpdateResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func decodeListGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	reqpb := request.(*v1.ListRequest)
	b, err := json.Marshal(reqpb)
	if err != nil {
		return nil, err
	}

	var req types.ListRequest
	err = json.Unmarshal(b, &req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func encodeListGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	respb := response.(*types.ListResponse)
	b, err := json.Marshal(respb)
	if err != nil {
		return nil, err
	}

	var resp v1.ListResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func decodeGetGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	reqpb := request.(*v1.GetRequest)
	b, err := json.Marshal(reqpb)
	if err != nil {
		return nil, err
	}

	var req types.GetRequest
	err = json.Unmarshal(b, &req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func encodeGetGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	respb := response.(*types.GetResponse)
	b, err := json.Marshal(respb)
	if err != nil {
		return nil, err
	}

	var resp v1.GetResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func decodeCreateGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	reqpb := request.(*v1.CreateRequest)
	b, err := json.Marshal(reqpb)
	if err != nil {
		return nil, err
	}

	var req types.CreateRequest
	err = json.Unmarshal(b, &req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func encodeCreateGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	respb := response.(*types.CreateResponse)
	b, err := json.Marshal(respb)
	if err != nil {
		return nil, err
	}

	var resp v1.CreateResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
