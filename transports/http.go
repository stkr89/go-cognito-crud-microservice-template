package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stkr89/go-cognito-crud-microservice-template/common"
	"github.com/stkr89/go-cognito-crud-microservice-template/endpoints"
	"github.com/stkr89/go-cognito-crud-microservice-template/middleware"
	"github.com/stkr89/go-cognito-crud-microservice-template/types"
	"net/http"
)

type errorWrapper struct {
	Error string `json:"error"`
}

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	commonOptions := []httptransport.ServerOption{
		httptransport.ServerBefore(parseHttpAccessToken),
		httptransport.ServerErrorEncoder(errorEncoder),
	}

	m := mux.NewRouter()
	m.Handle("/api/model/v1", httptransport.NewServer(
		endpoint.Chain(
			middleware.AuthenticateUser(),
			middleware.ValidateCreateInput(),
			middleware.ConformCreateInput(),
		)(endpoints.Create),
		decodeHTTPCreateRequest,
		encodeHTTPGenericResponse,
		commonOptions...,
	)).Methods(http.MethodPost)
	m.Handle("/api/model/v1/{id}", httptransport.NewServer(
		endpoint.Chain(
			middleware.AuthenticateUser(),
			middleware.ValidateGetInput(),
			middleware.ConformGetInput(),
		)(endpoints.Get),
		decodeHTTPGetRequest,
		encodeHTTPGenericResponse,
		commonOptions...,
	)).Methods(http.MethodGet)
	m.Handle("/api/model/v1", httptransport.NewServer(
		endpoint.Chain(
			middleware.AuthenticateUser(),
			middleware.ValidateListInput(),
			middleware.ConformListInput(),
		)(endpoints.List),
		decodeHTTPListRequest,
		encodeHTTPGenericResponse,
		commonOptions...,
	)).Methods(http.MethodGet)
	m.Handle("/api/model/v1", httptransport.NewServer(
		endpoint.Chain(
			middleware.AuthenticateUser(),
			middleware.ValidateUpdateInput(),
			middleware.ConformUpdateInput(),
		)(endpoints.Update),
		decodeHTTPUpdateRequest,
		encodeHTTPGenericResponse,
		commonOptions...,
	)).Methods(http.MethodPut)
	m.Handle("/api/model/v1/{id}", httptransport.NewServer(
		endpoint.Chain(
			middleware.AuthenticateUser(),
			middleware.ValidateDeleteInput(),
			middleware.ConformDeleteInput(),
		)(endpoints.Delete),
		decodeHTTPDeleteRequest,
		encodeHTTPGenericResponse,
		commonOptions...,
	)).Methods(http.MethodDelete)

	return m
}

func err2code(err *common.Error) int {
	switch err.Key {
	case common.InvalidRequestBody, common.InvalidID:
		return http.StatusBadRequest
	case common.Unauthorized:
		return http.StatusUnauthorized
	}

	return http.StatusInternalServerError
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err.(*common.Error)))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeHTTPDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, common.NewError(common.InvalidID, "invalid id")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, common.NewError(common.InvalidID, "invalid id")
	}

	return &types.DeleteRequest{
		ID: id,
	}, nil
}

func decodeHTTPListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := types.ListRequest{}
	return &req, nil
}

func decodeHTTPUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req types.UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, common.NewError(common.InvalidRequestBody, "invalid request body")
	}
	return &req, nil
}

func decodeHTTPGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, common.NewError(common.InvalidID, "invalid id")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, common.NewError(common.InvalidID, "invalid id")
	}

	return &types.GetRequest{
		ID: id,
	}, nil
}

func decodeHTTPCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req types.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, common.NewError(common.InvalidRequestBody, "invalid request body")
	}
	return &req, nil
}
