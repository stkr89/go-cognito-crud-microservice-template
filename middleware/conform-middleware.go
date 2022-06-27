package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/leebenson/conform"
	"github.com/stkr89/go-cognito-crud-microservice-template/common"
	"github.com/stkr89/go-cognito-crud-microservice-template/types"
)

func ConformListInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.ListRequest)
			err := conform.Strings(req)
			if err != nil {
				return nil, common.NewError(common.InvalidRequestBody, err.Error())
			}
			return next(ctx, req)
		}
	}
}

func ConformDeleteInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.DeleteRequest)
			err := conform.Strings(req)
			if err != nil {
				return nil, common.NewError(common.InvalidRequestBody, err.Error())
			}
			return next(ctx, req)
		}
	}
}

func ConformUpdateInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.UpdateRequest)
			err := conform.Strings(req)
			if err != nil {
				return nil, common.NewError(common.InvalidRequestBody, err.Error())
			}
			return next(ctx, req)
		}
	}
}

func ConformGetInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.GetRequest)
			err := conform.Strings(req)
			if err != nil {
				return nil, common.NewError(common.InvalidRequestBody, err.Error())
			}
			return next(ctx, req)
		}
	}
}

func ConformCreateInput() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(*types.CreateRequest)
			err := conform.Strings(req)
			if err != nil {
				return nil, common.NewError(common.InvalidRequestBody, err.Error())
			}
			return next(ctx, req)
		}
	}
}
