package helpers

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
)

type JsonResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ErrorRequest struct {
	ErrorMessage string `json:"error_message"`
	Field        string `json:"field"`
}

func ReadJSON[T any](w http.ResponseWriter, r *http.Request, data T) error {
	maxBytes := 1 << 31
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(data); err != nil {
		return err
	}

	if err := dec.Decode(&struct{}{}); err != nil {
		return errors.New("Body must have onl a single JSON value")
	}
	return nil
}

func WriteJSON[T any](w http.ResponseWriter, status int, data JsonResponse[T], headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(out); err != nil {
		return err
	}
	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	var statusCode = http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload = JsonResponse[error]{
		Success: true,
		Message: err.Error(),
		Data:    err,
	}

	return WriteJSON(w, statusCode, payload)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
