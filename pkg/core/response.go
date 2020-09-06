package core

import (
	"encoding/json"
	"net/http"
)

func MethodNotAllowed(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	m := map[string]interface{}{"message": message}
	if rs, err := json.Marshal(m); err == nil {
		_, _ = w.Write(rs)
	}
}

func BadRequest(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	m := map[string]interface{}{"message": message}
	if rs, err := json.Marshal(m); err == nil {
		_, _ = w.Write(rs)
	}
}

func NotFound(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	m := map[string]interface{}{"message": message}
	if rs, err := json.Marshal(m); err == nil {
		_, _ = w.Write(rs)
	}
}

func Ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if rs, err := json.Marshal(data); err == nil {
		_, _ = w.Write(rs)
	}
}

func Created(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if rs, err := json.Marshal(data); err == nil {
		_, _ = w.Write(rs)
	}
}

func Unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte("401 Unauthorized"))
}
