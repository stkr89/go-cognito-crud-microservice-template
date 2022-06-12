package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stkr89/modelsvc/common"
	"github.com/stkr89/modelsvc/endpoints"
	"github.com/stkr89/modelsvc/middleware"
	"github.com/stkr89/modelsvc/types"
	"net/http"
)

type errorWrapper struct {
	Error string `json:"error"`
}

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := mux.NewRouter()
	m.Handle("/api/v1/model", httptransport.NewServer(
		endpoint.Chain(
			middleware.ValidateCreateInput(),
			middleware.ConformCreateInput(),
		)(endpoints.Create),
		decodeHTTPCreateRequest,
		encodeHTTPGenericResponse,
	)).Methods(http.MethodPost)
	m.Handle("/api/v1/model/{id}", httptransport.NewServer(
		endpoint.Chain(
			middleware.ValidateGetInput(),
			middleware.ConformGetInput(),
		)(endpoints.Get),
		decodeHTTPGetRequest,
		encodeHTTPGenericResponse,
	)).Methods(http.MethodGet)
	m.Handle("/api/v1/model", httptransport.NewServer(
		endpoint.Chain(
			middleware.ValidateListInput(),
			middleware.ConformListInput(),
		)(endpoints.List),
		decodeHTTPListRequest,
		encodeHTTPGenericResponse,
	)).Methods(http.MethodGet)
	m.Handle("/api/v1/model", httptransport.NewServer(
		endpoint.Chain(
			middleware.ValidateUpdateInput(),
			middleware.ConformUpdateInput(),
		)(endpoints.Update),
		decodeHTTPUpdateRequest,
		encodeHTTPGenericResponse,
	)).Methods(http.MethodPut)
	m.Handle("/api/v1/model/{id}", httptransport.NewServer(
		endpoint.Chain(
			middleware.ValidateDeleteInput(),
			middleware.ConformDeleteInput(),
		)(endpoints.Delete),
		decodeHTTPDeleteRequest,
		encodeHTTPGenericResponse,
	)).Methods(http.MethodDelete)

	return m
}

func err2code(err *common.Error) int {
	switch err.Key {
	case common.InvalidRequestBody:
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
		return nil, err
	}

	return &types.DeleteRequest{
		ID: id,
	}, err
}

func decodeHTTPListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := types.ListRequest{}
	return &req, nil
}

func decodeHTTPUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req types.UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func decodeHTTPGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return nil, common.NewError(common.InvalidID, "invalid id")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	return &types.GetRequest{
		ID: id,
	}, err
}

func decodeHTTPCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req types.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}
