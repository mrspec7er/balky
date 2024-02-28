package utils

import (
	"encoding/json"
	"net/http"
)

type Metadata struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Count int64 `json:"count"`
}

type Response struct{}

func (r Response) InternalServerErrorHandler(w http.ResponseWriter, status int, err error) {
	message := err.Error()
	response := struct {
		Status   bool        `json:"status"`
		Message  *string     `json:"message"`
		Data     interface{} `json:"data"`
		Metadata interface{} `json:"metadata"`
	}{
		Status:   false,
		Message:  &message,
		Data:     nil,
		Metadata: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		r.InternalServerErrorHandler(w, 500, err)
	}

	w.WriteHeader(status)
	w.Write(responseData)
}

func (r Response) NotFoundHandler(w http.ResponseWriter) {
	message := "Data Not Found"
	response := struct {
		Status   bool        `json:"status"`
		Message  *string     `json:"message"`
		Data     interface{} `json:"data"`
		Metadata interface{} `json:"metadata"`
	}{
		Status:   false,
		Message:  &message,
		Data:     nil,
		Metadata: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		r.InternalServerErrorHandler(w, 500, err)
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write(responseData)
}

func (r Response) BadRequestHandler(w http.ResponseWriter) {
	message := "Bad Request"
	response := struct {
		Status   bool        `json:"status"`
		Message  *string     `json:"message"`
		Data     interface{} `json:"data"`
		Metadata interface{} `json:"metadata"`
	}{
		Status:   false,
		Message:  &message,
		Data:     nil,
		Metadata: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		r.InternalServerErrorHandler(w, 500, err)
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write(responseData)
}

func (r Response) UnauthorizeUser(w http.ResponseWriter) {
	message := "Unauthorize user"
	response := struct {
		Status   bool        `json:"status"`
		Message  *string     `json:"message"`
		Data     interface{} `json:"data"`
		Metadata interface{} `json:"metadata"`
	}{
		Status:   false,
		Message:  &message,
		Data:     nil,
		Metadata: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		r.InternalServerErrorHandler(w, 500, err)
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write(responseData)
}

func (r Response) SuccessMessageResponse(w http.ResponseWriter, message string) {
	response := struct {
		Status   bool        `json:"status"`
		Message  *string     `json:"message"`
		Data     interface{} `json:"data"`
		Metadata interface{} `json:"metadata"`
	}{
		Status:   true,
		Message:  &message,
		Data:     nil,
		Metadata: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		r.InternalServerErrorHandler(w, 500, err)
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(responseData)
}

func (r Response) MutationSuccessResponse(w http.ResponseWriter, message string) {
	response := struct {
		Status   bool        `json:"status"`
		Message  *string     `json:"message"`
		Data     interface{} `json:"data"`
		Metadata interface{} `json:"metadata"`
	}{
		Status:   true,
		Message:  &message,
		Data:     nil,
		Metadata: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		r.InternalServerErrorHandler(w, 500, err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

func (r Response) GetSuccessResponse(w http.ResponseWriter, message *string, data interface{}, metadata *Metadata) {
	response := struct {
		Status   bool        `json:"status"`
		Message  *string     `json:"message"`
		Data     interface{} `json:"data"`
		Metadata interface{} `json:"metadata"`
	}{
		Status:   true,
		Message:  message,
		Data:     data,
		Metadata: metadata,
	}
	w.Header().Set("Content-Type", "application/json")
	responseData, err := json.Marshal(response)
	if err != nil {
		r.InternalServerErrorHandler(w, 500, err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
