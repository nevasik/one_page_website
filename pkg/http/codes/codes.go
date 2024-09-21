package codes

import (
	httpModel "Odnostranishka/pkg/http/model"
	"net/http"
)

const (
	UnknownError = -1
	NoError      = 0

	// Unauthorized general error code
	Unauthorized = 10000

	// InputError general error code
	InputError  = 20000
	BodyBinding = 20001

	// Forbidden general error code
	Forbidden = 30000

	// InternalError general error code
	InternalError = 40000
)

var InternalErr = map[int]httpModel.InternalError{
	// 1XXXX credentials validation errors
	Unauthorized: {
		Message: "unauthorized",
		Status:  http.StatusUnauthorized,
	},

	// 2XXXX errors caused by user input
	InputError: {
		Message: "wrong input",
		Status:  http.StatusBadRequest,
	},
	BodyBinding: {
		Message: "can't bind body to request model",
		Status:  http.StatusUnprocessableEntity,
	},

	// 3XXXX no access to resources
	Forbidden: {
		Message: "forbidden",
		Status:  http.StatusForbidden,
	},

	// 4XXXX internal errors
	UnknownError: {
		Message: "unknown internal error",
		Status:  http.StatusInternalServerError,
	},

	InternalError: {
		Message: "internal error",
		Status:  http.StatusInternalServerError,
	},
}
