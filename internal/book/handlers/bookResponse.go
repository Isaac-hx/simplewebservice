// This file handle response to json formatting
package handlerss

import (
	"encoding/json"
	"net/http"
)

type baseResponse struct {
	Message string
}

func Response(w http.ResponseWriter, codeStatus int, message string) error {
	w.WriteHeader(codeStatus)
	return json.NewEncoder(w).Encode(&baseResponse{Message: message})

}
