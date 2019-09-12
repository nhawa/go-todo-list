package lib

import (
	"encoding/json"
	"net/http"
)

type (
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Reason  string `json:"reason,omitempty"`
	}

	ResponseData struct {
		Status   *Status     `json:"status"`
		Data     interface{} `json:"data,omitempty"`
		httpCode int
	}

	Response interface {
		JSON(w http.ResponseWriter)
	}
)

// JsonResponse send response
func CreateResponse(httpCode int, responseCode int, message string, data interface{}) Response {
	return &ResponseData{
		Status: &Status{
			Code:    responseCode,
			Message: message,
		},
		Data:     data,
		httpCode: httpCode,
	}
}

func CreateFailResponse(httpCode int, responseCode int, message string, data interface{}, reason string) Response {
	return &ResponseData{
		Status: &Status{
			Code:    responseCode,
			Message: message,
			Reason:  reason,
		},
		Data:     data,
		httpCode: httpCode,
	}

}

// JSON Kembalikan Response berupa JSON
func (response *ResponseData) JSON(w http.ResponseWriter) {
	w.WriteHeader(response.httpCode)
	json.NewEncoder(w).Encode(response)
}
