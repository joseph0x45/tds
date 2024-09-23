package transport

import (
	"encoding/json"
	"net/http"
)

type Err struct {
	Code string `json:"code"`
}

type ErrorResponse struct {
	Error Err `json:"error"`
}

type ServiceResponse struct {
	StatusCode int
	Data       interface{}
	Err        error
}

func (sr ServiceResponse) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(sr.StatusCode)
	if sr.Err != nil {
		bytes, _ := json.Marshal(&ErrorResponse{
			Error: Err{Code: sr.Err.Error()},
		})
		w.Write(bytes)
		return
	}
	if sr.Data == nil {
		sr.Data = map[string]interface{}{}
	}
	if err := json.NewEncoder(w).Encode(sr.Data); err != nil {
		http.Error(w, "failed to encode response", 500)
	}
}

func WriteOKNoData(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func WriteError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	bytes, _ := json.Marshal(&ErrorResponse{
		Error: Err{Code: err.Error()},
	})
	w.Write(bytes)
}
