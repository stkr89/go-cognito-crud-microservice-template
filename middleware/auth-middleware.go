package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
	"github.com/stkr89/modelsvc/common"
)

func AuthenticateUser() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			errMsg := "unauthorized access"

			user, err := common.GetUserFromContext(ctx)
			if err != nil {
				return nil, common.NewError(common.Unauthorized, errMsg)
			}

			err = validator.New().Struct(user)
			if err != nil {
				return nil, common.NewError(common.Unauthorized, errMsg)
			}

			return next(ctx, request)
		}
	}
}
