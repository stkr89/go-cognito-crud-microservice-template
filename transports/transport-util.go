package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/stkr89/go-cognito-crud-microservice-template/types"
	"net/http"
	"os"
	"strings"
)

func parseHttpAccessToken(ctx context.Context, r *http.Request) context.Context {
	authHeader := r.Header.Get("Authorization")
	splitAuthHeader := strings.Split(authHeader, " ")
	if len(splitAuthHeader) != 2 {
		return ctx
	}

	pubKeyURL := "https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json"
	formattedURL := fmt.Sprintf(pubKeyURL, os.Getenv("AWS_COGNITO_REGION"), os.Getenv("AWS_COGNITO_USER_POOL_ID"))

	keySet, err := jwk.Fetch(r.Context(), formattedURL)
	if err != nil {
		return ctx
	}

	token, err := jwt.Parse(
		[]byte(splitAuthHeader[1]),
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
	)
	if err != nil {
		return ctx
	}

	tokenMap, err := token.AsMap(ctx)
	if err != nil {
		return ctx
	}

	id, err := uuid.Parse(tokenMap["sub"].(string))
	if err != nil {
		return ctx
	}

	user := types.User{
		ID:        id,
		FirstName: tokenMap["custom:firstName"].(string),
		LastName:  tokenMap["custom:lastName"].(string),
		Email:     tokenMap["email"].(string),
	}
	userBytes, err := json.Marshal(user)
	if err != nil {
		return ctx
	}

	ctx = context.WithValue(ctx, "user", userBytes)

	return ctx
}
