package middleware

import (
	_ "Odnostranishka/internal/plugin/model"
	"Odnostranishka/pkg/http/codes"
	"bytes"
	lg "gitlab.com/nevasik7/logger"
	"io"
	"net/http"
)

func (m *mw) validateAndCheckSignature(r *http.Request) int {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		lg.Errorf("Error read body, err=%v", err)
		return codes.BodyBinding
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	return codes.NoError
}
