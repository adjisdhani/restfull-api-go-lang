package middleware

import (
	"belajar_golang_restful_api/helper"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggerMiddleware struct {
	Handler http.Handler
	Logger  *logrus.Logger
}

func NewLoggerMiddleware(handler http.Handler) *LoggerMiddleware {
	return &LoggerMiddleware{
		Handler: handler,
		Logger:  helper.NewLogger(),
	}
}

func (middleware *LoggerMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()

	middleware.Handler.ServeHTTP(writer, request)

	duration := time.Since(start)

	middleware.Logger.WithFields(logrus.Fields{
		"method":   request.Method,
		"uri":      request.RequestURI,
		"duration": duration.Milliseconds(),
	}).Info("HTTP Request")
}
