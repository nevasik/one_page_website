package router

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var (
	v = validator.New(validator.WithRequiredStructEnabled())
)

func decodeAndValidate[reqType any](r *http.Request, req reqType) (err error) {
	if err = json.NewDecoder(r.Body).Decode(req); err != nil {
		return
	}
	return v.Struct(req)
}
