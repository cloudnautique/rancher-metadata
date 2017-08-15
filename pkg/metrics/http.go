package metrics

import (
	"net/http"
	"strconv"
	"strings"
)

type ObserveableResponseWriter struct {
	http.ResponseWriter

	status      int
	wroteHeader bool
}

func (orw *ObserveableResponseWriter) Status() int {
	return orw.status
}

func (orw *ObserveableResponseWriter) WriteHeader(code int) {
	orw.status = code
	orw.wroteHeader = true

	orw.ResponseWriter.WriteHeader(code)
}

func (orw *ObserveableResponseWriter) Write(b []byte) (int, error) {
	if !orw.wroteHeader {
		orw.WriteHeader(http.StatusOK)
	}

	return orw.ResponseWriter.Write(b)
}

func (orw *ObserveableResponseWriter) Header() http.Header {
	return orw.ResponseWriter.Header()
}

func NewStatusObserveableResponseWriter(w http.ResponseWriter) *ObserveableResponseWriter {
	return &ObserveableResponseWriter{
		ResponseWriter: w,
	}
}

func StatusLabelString(code int) string {
	// Faster to return knowns vs. convert
	switch code {
	case 200:
		return "200"
	case 404:
		return "404"
	case 500:
		return "500"

	default:
		return strconv.Itoa(code)
	}
}

func CleanMethodString(method string) string {
	// Faster to return knowns vs. convert
	switch method {
	case "GET":
		return "get"
	case "POST":
		return "post"
	case "HEAD":
		return "head"
	case "PUT":
		return "put"
	default:
		return strings.ToLower(method)
	}
}
