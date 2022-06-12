package common

type Error struct {
	Key     string
	Message string
}

func NewError(key string, message string) *Error {
	return &Error{Key: key, Message: message}
}

func (e Error) Error() string {
	return e.Message
}

var (
	InvalidRequestBody = "Invalid request body"
	SomethingWentWrong = "Something went wrong"
	InvalidID          = "Invalid ID"
	Unauthorized       = "Unauthorized"
)
