package middleware

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/stkr89/modelsvc/common"
	"github.com/stkr89/modelsvc/types"
)

func AuthenticateUser() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			userBytes := ctx.Value("user")

			if userBytes == nil {
				return nil, common.NewError(common.Unauthorized, "unauthorized access")
			}

			var user types.User
			err := json.Unmarshal(userBytes.([]byte), &user)
			if err != nil {
				return nil, common.NewError(common.Unauthorized, "unauthorized access")
			}

			return next(ctx, request)
		}
	}
}
