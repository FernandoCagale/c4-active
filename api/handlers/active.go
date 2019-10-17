package handlers

import (
	"encoding/json"
	"github.com/FernandoCagale/c4-active/api/render"
	"github.com/FernandoCagale/c4-active/pkg/domain/active"
	"github.com/FernandoCagale/c4-active/pkg/entity"
	"github.com/FernandoCagale/c4-active/pkg/infra/errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ActiveHandler struct {
	usecase *active.ActiveUseCase
}

func NewActive(usecase *active.ActiveUseCase) *ActiveHandler {
	return &ActiveHandler{
		usecase: usecase,
	}
}

func (handler *ActiveHandler) Actives(w http.ResponseWriter, r *http.Request) {
	render.Response(w, map[string]bool{"ok": true}, http.StatusOK)
}

func (handler *ActiveHandler) Create(w http.ResponseWriter, r *http.Request) {
	var active *entity.Active

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&active); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := handler.usecase.Create(active); err != nil {
		switch err {
		case errors.ErrInvalidPayload:
			render.ResponseError(w, err, http.StatusBadRequest)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, active, http.StatusCreated)
}

func (handler *ActiveHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	var active entity.Active
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&active); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err = handler.usecase.Update(id, &active)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		case errors.ErrInvalidPayload:
			render.ResponseError(w, err, http.StatusBadRequest)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusNoContent)
}

func (handler *ActiveHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	err = handler.usecase.Delete(id)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusNoContent)
}

func (handler *ActiveHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	active, err := handler.usecase.FindByID(id)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, active, http.StatusOK)
}

func (handler *ActiveHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	actives, err := handler.usecase.FindAll()
	if err != nil {
		render.ResponseError(w, err, http.StatusInternalServerError)
		return
	}

	render.Response(w, actives, http.StatusOK)
}
