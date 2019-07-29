package render

import (
	"encoding/json"
	"net/http"

	"github.com/FernandoCagale/c4-active/pkg/infra/errors"
)

//Response response status code
func Response(w http.ResponseWriter, response interface{}, code int) {
	json, err := json.Marshal(response)
	if err != nil {
		ResponseError(w, errors.AddInternalServerError(err.Error()))
		return
	}
	addHeaderDefaults(w, code)
	w.Write(json)
}

//ResponseError error
func ResponseError(w http.ResponseWriter, response errors.ResponseError) {
	json, err := json.Marshal(response)
	if err != nil {
		ResponseError(w, errors.AddInternalServerError(err.Error()))
		return
	}
	addHeaderDefaults(w, response.Code)
	w.Write(json)
}

func addHeaderDefaults(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}
