package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	responseData struct {
		status int
		size   int
	}

	wrappedResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func (lrw *wrappedResponseWriter) Write(b []byte) (int, error) {
	size, err := lrw.ResponseWriter.Write(b)
	lrw.responseData.size += size
	return size, err
}

func (lrw *wrappedResponseWriter) WriteHeader(statusCode int) {
	lrw.ResponseWriter.WriteHeader(statusCode)
	lrw.responseData.status = statusCode
}

func LoggedHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()
		uri := r.RequestURI
		method := r.Method

		rd := &responseData{
			status: 0,
			size:   0,
		}
		lrw := wrappedResponseWriter{
			ResponseWriter: w,
			responseData:   rd,
		}

		h(&lrw, r)

		d := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"uri":      uri,
			"method":   method,
			"duration": d,
			"status":   rd.status,
			"size":     rd.size,
		}).Info("request completed")
	}
}
