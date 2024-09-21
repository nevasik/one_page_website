package router

import (
	"Odnostranishka/pkg/http/codes"
	"Odnostranishka/pkg/http/response"
	"context"
	"fmt"
	lg "gitlab.com/nevasik7/logger"
	"net/http"
)

func FromBody[reqType any, respType any](ucFn func(ctx context.Context, req reqType) (respType, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req  reqType
			resp any
			err  error
			code int
		)

		if err = decodeAndValidate(r, &req); err != nil {
			lg.Error(err)
			response.Write(w, codes.BodyBinding, nil)
			return
		}

		if code != codes.NoError {
			response.Write(w, code, nil)
			return
		}

		resp, code = ucFn(r.Context(), req)

		lg.Tracef("RESP %v\n%+v", r.URL.Path, fmt.Sprintf("%+v", resp))

		response.Write(w, code, resp)
	}
}
