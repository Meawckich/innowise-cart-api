package endpoint

import (
	"cart-api/internal/pkg/common/db/repository"
	"cart-api/internal/pkg/common/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type CartItemHandler struct {
	pool *sqlx.DB
}

func NewCartItemHandler(dbPool *sqlx.DB) *CartItemHandler {
	return &CartItemHandler{
		pool: dbPool,
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
func (c *CartItemHandler) AddToCart(repo repository.ItemRepository) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		body := req.Body
		defer body.Close()

		var dto model.ItemDto

		dec := json.NewDecoder(body)
		dec.DisallowUnknownFields()

		if err := dec.Decode(&dto); err != nil {
			http.Error(res, "Cannot decode body", http.StatusBadRequest)
			return
		}

		idPath := req.PathValue("cartId")
		cartId, err := strconv.Atoi(idPath)
		if err != nil {
			return
		}

		item, err := repo.Create(dto, cartId)

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)

			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(item)
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
func (c *CartItemHandler) RemoveFromCart(itemRepo repository.ItemRepository, cartRepo repository.ICartRepository) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		pathCartId := req.PathValue("cartId")
		cartId, err := strconv.Atoi(pathCartId)
		if err != nil {
			model.NewResponseError(http.StatusBadRequest, "Invalid cart id").ShowError(res)
			return
		}
		if _, err := cartRepo.GetById(cartId); err != nil {
			http.Error(res, "cannot find cart by given id", http.StatusBadRequest)
			return
		}

		pathItemId := req.PathValue("itemId")
		itemId, err := strconv.Atoi(pathItemId)
		if err != nil {
			http.Error(res, "Invalid item id", http.StatusBadRequest)
			return
		}

		if err := itemRepo.Delete(cartId, itemId); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte("{}"))
	}
}
