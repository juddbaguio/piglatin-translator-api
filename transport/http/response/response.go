package response

import (
	"encoding/json"
	"net/http"
)

func coreResponse(statusCode int, w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonWriter := json.NewEncoder(w)
	jsonWriter.Encode(data)
}

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func NewCustomSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	coreResponse(statusCode, w, Response{
		StatusCode: statusCode,
		Data:       data,
	})
}

func NewCustomErrResponse(w http.ResponseWriter, statusCode int, err error) {
	coreResponse(statusCode, w, Response{
		StatusCode: statusCode,
		Message:    err.Error(),
	})
}

func NewInternalServerError(w http.ResponseWriter, err error) {
	response := Response{
		StatusCode: http.StatusInternalServerError,
	}

	if err != nil {
		response.Message = err.Error()
	}

	coreResponse(http.StatusInternalServerError, w, response)
}

func NewSuccess(w http.ResponseWriter, data interface{}) {
	response := Response{
		StatusCode: http.StatusOK,
	}

	if data != nil {
		response.Data = data
	}

	coreResponse(http.StatusOK, w, response)
}

func NewCreated(w http.ResponseWriter, message interface{}, data interface{}) {
	response := Response{
		StatusCode: http.StatusCreated,
		Message:    message,
		Data:       data,
	}

	coreResponse(http.StatusCreated, w, response)
}
