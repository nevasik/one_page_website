package middleware

import (
	"Odnostranishka/pkg/http/codes"
	"Odnostranishka/pkg/http/response"
	"net/http"
)

type I interface {
	Verify(next http.Handler) http.Handler
}

type mw struct {
}

func New() I {
	return &mw{}
}

func (m *mw) Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := m.validateAndCheckSignature(r)
		if code != codes.NoError {
			response.Write(w, code, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
