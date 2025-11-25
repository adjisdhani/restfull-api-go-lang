package middleware

import (
	"belajar_golang_restful_api/helper"
	"belajar_golang_restful_api/model/web"
	"fmt"
	"net/http"
	"os"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (authMiddleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	config, err := helper.LoadConfigNew(".")
	helper.PanicIfError(err)

	fmt.Println("HEADER:", request.Header.Get("X-API-KEY"))
	fmt.Println("CONFIG:", config)
	fmt.Println("ENV RAW:", os.Getenv("X_API_KEY"))

	if request.Header.Get("X-API-KEY") == config.X_API_KEY {
		authMiddleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
