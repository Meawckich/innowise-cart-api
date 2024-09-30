package handler

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func GroupSwaggerRoute(mux *http.ServeMux) {
	mux.Handle("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))
}
