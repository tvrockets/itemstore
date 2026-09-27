package handler

import (
	"encoding/json"
	"errors"
	"github.com/itemstore/tvrockets/internal/model"
	"github.com/itemstore/tvrockets/internal/service"
	"net/http"
)

type ItemHandler struct {
	service *service.ItemService
}

func NewItemHandler(service *service.ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

func (i *ItemHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /items", i.GetAll)
	mux.HandleFunc("GET /items/{id}", i.GetByID)
	mux.HandleFunc("POST /items", i.Create)
	mux.HandleFunc("PUT /items/{id}", i.Update)
	mux.HandleFunc("DELETE /items/{id}", i.Delete)
}

func (i *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	items, err := i.service.GetAll()
	if err != nil {
		handleError(w, err)
		return
	}
	respondJSON(w, http.StatusOK, items)
}
func (i *ItemHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	item, err := i.service.GetByID(id)
	if err != nil {
		handleError(w, err)
		return
	}
	respondJSON(w, http.StatusOK, item)
}

func (i *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdItem, err := i.service.Create(item)
	if err != nil {
		handleError(w, err)
		return
	}
	respondJSON(w, http.StatusCreated, createdItem)
}

func (i *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updated, err := i.service.Update(id, item)
	if err != nil {
		handleError(w, err)
		return
	}
	respondJSON(w, http.StatusOK, updated)
}

func (i *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := i.service.Delete(id)
	if err != nil {
		handleError(w, err)
		return
	}
}

func handleError(w http.ResponseWriter, err error) {
	var ve *model.ValidationError
	switch {
	case errors.As(err, &ve):
		http.Error(w, ve.Error(), http.StatusBadRequest)
	case errors.Is(err, model.ErrNotFound):
		http.Error(w, "не найдено", http.StatusNotFound)
	case errors.Is(err, model.ErrAlreadyExists):
		http.Error(w, "уже существует", http.StatusConflict)
	default:
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
