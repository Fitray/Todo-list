package core_http_responce

import "net/http"

var (
	StatusCodeUninitialazed = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponceWritter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatusCodeUninitialazed,
	}
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (w *ResponseWriter) GetStatusCodeOrPanic() int {
	if w.statusCode == StatusCodeUninitialazed {
		panic("No status code set")
	}
	return w.statusCode
}
