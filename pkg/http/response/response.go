package response

import (
	"Odnostranishka/pkg/http/codes"
	httpModel "Odnostranishka/pkg/http/model"
	"encoding/json"
	lg "gitlab.com/nevasik7/logger"
	"net/http"
)

// Write data is the Data field of the corresponding response model
func Write(w http.ResponseWriter, code int, data any) {
	var resp httpModel.BaseResponse

	w.Header().Set("Content-Type", "application/json")

	if code == codes.NoError {
		w.WriteHeader(http.StatusOK)
		resp.Data = data
		resp.Success = true
		resp.Error = nil
	} else {
		if _, ok := codes.InternalErr[code]; !ok {
			resp.Error = &httpModel.ErrorResponse{
				Message: "internal server error",
				Code:    codes.UnknownError,
			}
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			internalErr := codes.InternalErr[code]
			resp.Error = &httpModel.ErrorResponse{
				Message: internalErr.Message,
				Code:    code,
			}
			w.WriteHeader(internalErr.Status)
		}
	}

	r, err := json.Marshal(resp)
	if err != nil {
		lg.Errorf("response marshalling failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(r)
	if err != nil {
		lg.Errorf("writing response bytes failed: %v", err)
	}
}
