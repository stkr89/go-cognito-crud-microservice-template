package common

import (
	"context"
	"encoding/json"
	"github.com/stkr89/modelsvc/types"
)

func GetUserFromContext(ctx context.Context) (*types.User, error) {
	userBytes := ctx.Value("user")
	if userBytes == nil {
		return nil, NewError(UserNotFoundInContext, UserNotFoundInContext)
	}

	var user types.User
	err := json.Unmarshal(userBytes.([]byte), &user)
	if err != nil {
		return nil, NewError(UserNotFoundInContext, UserNotFoundInContext)
	}
	return &user, nil
}
