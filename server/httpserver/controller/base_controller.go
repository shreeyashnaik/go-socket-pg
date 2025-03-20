package controller

import (
	"chat-app/common/schemas"
	"encoding/json"
	"net/http"
)

type BaseController struct {
}

func (b *BaseController) ResponseOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(schemas.GenericHTTPResponse{
		Success: true,
		Data:    data,
	})
}

func (b *BaseController) ResponseCreated(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(schemas.GenericHTTPResponse{
		Success: true,
		Data:    data,
	})
}

func (b *BaseController) ResponseBadRequest(w http.ResponseWriter, errorMessage string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(schemas.GenericHTTPResponse{
		Success: false,
		Error:   errorMessage,
		Data:    data,
	})
}

func (b *BaseController) ResponseMethodNotAllowed(w http.ResponseWriter, errorMessage string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(schemas.GenericHTTPResponse{
		Success: false,
		Error:   errorMessage,
		Data:    data,
	})
}

func (b *BaseController) ResponseInternalServerError(w http.ResponseWriter, errorMessage string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(schemas.GenericHTTPResponse{
		Success: false,
		Error:   errorMessage,
		Data:    data,
	})
}
