package active

import (
	"net/http"

	"github.com/FernandoCagale/c4-active/api/handlers/active/response"
	"github.com/FernandoCagale/c4-active/api/infra"
	"github.com/FernandoCagale/c4-active/pkg/active"
	"github.com/FernandoCagale/c4-active/pkg/entity"
	"github.com/FernandoCagale/c4-active/pkg/infra/errors"
	"github.com/FernandoCagale/c4-active/pkg/infra/render"
)

//FindAll findAll
func FindAll(service active.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")

		limit, err := infra.GetLimit(r)
		if err != nil {
			render.ResponseError(w, errors.AddBadRequestError("Failed to decode request query"))
			return
		}

		page, err := infra.GetPage(r)
		if err != nil {
			render.ResponseError(w, errors.AddBadRequestError("Failed to decode request query"))
			return
		}

		query := entity.NewQuery(search, limit, page)

		actives, count, err := service.FindAll(query)
		if err != nil {
			render.ResponseError(w, errors.Err(err))
			return
		}

		render.Response(w, response.NewResponseFinAll(count, actives), http.StatusOK)
	})
}
