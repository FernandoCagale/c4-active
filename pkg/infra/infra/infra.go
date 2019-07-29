package infra

import (
	"net/http"

	errors "github.com/FernandoCagale/c4-active/pkg/infra/errors"
	"github.com/FernandoCagale/c4-active/pkg/infra/render"
)

//NotFoundHandler NotFound
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.ResponseError(w, errors.AddNotFoundError("Not Found"))
	})
}

//NotAllowedHandler NotAllowed
func NotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.ResponseError(w, errors.AddMethodNotAllowedError("Method Not Allowed"))
	})
}
