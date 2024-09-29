package handler

import (
	"cart-api/internal/handler/middleware"
	"cart-api/internal/service"
	"encoding/json"
	"net/http"
)

type ItemHandler struct {
	service *service.ItemService
}

func NewItemHandler(service *service.ItemService) *ItemHandler {
	return &ItemHandler{
		service: service,
	}
}

// DeleteItem godoc
//
//	@Summary		Delete an item from cart
//	@Description	delete item from cart, using cart id and item id
//	@Tags			items
//	@Produce		json
//	@Param			cartId	path		int	true	"Cart id"
//	@Param			itemId	path		int	true	"Item id"
//	@Success		200		{array}		byte
//	@Failure		400		{object}	model.ResponseError
//	@Router			/carts/{cartId}/items/{itemId} [delete]
func (h *ItemHandler) RemoveFromCart(w http.ResponseWriter, req *http.Request) {
	if err := h.service.DeleteItem(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
	}
}

// AddToCart add a new item into exists cart
//
//	@Summary	Add item to cart
//	@Tags		items
//	@Accept		json
//	@Produce	json
//	@Param		cartId	path		int				true	"cart id"
//	@Param		item	body		model.ItemDto	true	"item to add"
//	@Success	200		{object}	model.CartItem
//	@Failure	400		{object}	model.ResponseError
//	@Failire	500 {object} model.ResponseError
//	@Router		/carts/{cartId}/items [post]
func (h *ItemHandler) AddToCart(w http.ResponseWriter, req *http.Request) {
	if newItem, err := h.service.CreateItem(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newItem)
	}

}

func (h *ItemHandler) GroupRoutes(mux *http.ServeMux) {
	wrappedPath := middleware.NewValiDateMiddleWare(h.AddToCart)

	mux.Handle("POST /carts/{cartId}/items", wrappedPath)
	mux.HandleFunc("DELETE /carts/{cartId}/items/{itemId}", h.RemoveFromCart)
}
