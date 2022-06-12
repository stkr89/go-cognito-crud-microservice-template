package types

import "github.com/google/uuid"

type User struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
}

type CreateRequest struct {
}

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type GetResponse struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type ListRequest struct {
}

type ListResponse struct {
	Data []*ListResponseData `json:"data"`
}

type ListResponseData struct {
	ID uuid.UUID `json:"id"`
}

type UpdateRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type UpdateResponse struct {
	ID uuid.UUID `json:"id"`
}

type DeleteRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
