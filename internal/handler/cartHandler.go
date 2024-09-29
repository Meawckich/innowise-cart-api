package handler

import (
	"cart-api/internal/service"
	"encoding/json"
	"net/http"
)

type CartHandler struct {
	service *service.CartService
}

// in case when service needs an request rapams, it can passed with param into service method
func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{
		service: service,
	}
}

// CreateCart creates a new cart
//
//	@Summary		Creates a new cart and id generated
//	@Description	create cart
//	@Tags			carts
//	@Produce		json
//	@Success		200	{object}	model.Cart
//	@Failure		500	{object}	model.ResponseError
//	@Router			/carts [post]
func (h *CartHandler) CreateCart(w http.ResponseWriter, req *http.Request) {
	if newCart, err := h.service.CreateCart(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(newCart)
	}
}

// ViewCart shows cart
//
//	@Summary		Shows cart by id
//	@Description	get cart by id
//	@Tags			carts
//	@Produce		json
//	@Param			id	path		int	true	"id to find cart"
//	@Success		200	{object}	model.Cart
//	@Failure		400	{object}	model.ResponseError
//	@Failure		404	{object}	model.ResponseError
//	@Router			/carts/{id} [get]
func (h *CartHandler) ViewCart(w http.ResponseWriter, req *http.Request) {
	if cart, err := h.service.ViewCart(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(cart)
	}
}

// ListCarts lists all existing carts with items
//
//	@Summary		List carts
//	@Description	get all carts with composition item slice inside
//	@Tags			carts
//	@Produce		json
//	@Success		200	{array}		model.Cart
//	@Failure		500	{object}	model.ResponseError
//	@Router			/carts [get]
func (h *CartHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	if carts, err := h.service.GetCarts(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(carts)
	}
}

// DeleteCart  delete cart
//
//	@Summary		Delete a cart
//	@Description	delete a cart with its items recursively
//	@Tags			carts
//	@Produce		json
//	@Param			id	path		int	true	"Cart id"
//	@Success		200	{array}		byte
//	@Failure		400	{object}	model.ResponseError
//	@Failure		500	{object}	model.ResponseError
//	@Router			/carts/{id} [delete]
func (h *CartHandler) DeleteCart(w http.ResponseWriter, req *http.Request) {
	if err := h.service.DeleteCart(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
	}
}

func (h *CartHandler) GroupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /carts", h.CreateCart)
	mux.HandleFunc("GET /carts/{id}", h.ViewCart)
	mux.HandleFunc("GET /carts", h.GetAll)
	mux.HandleFunc("DELETE /carts/{id}", h.DeleteCart)
}
