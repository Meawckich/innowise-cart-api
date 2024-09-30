package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cart-api/internal/pkg/model"
	"cart-api/internal/repository"
)

type ItemService struct {
	itemRepo repository.ItemRepository
	cartRepo repository.ICartRepository
}

func NewItemService(itemRepo repository.ItemRepository, cartRepo repository.ICartRepository) *ItemService {
	return &ItemService{
		itemRepo: itemRepo,
		cartRepo: cartRepo,
	}
}

func (s *ItemService) CreateItem(req *http.Request) (model.CartItem, error) {
	body := req.Body
	defer body.Close()

	var dto model.ItemDto

	dec := json.NewDecoder(body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&dto); err != nil {
		return model.CartItem{}, err
	}

	idPath := req.PathValue("cartId")

	cartId, err := strconv.Atoi(idPath)
	if err != nil {
		return model.CartItem{}, err
	}

	newItem, err := s.itemRepo.Create(dto, cartId)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *ItemService) DeleteItem(req *http.Request) error {
	pathCartId := req.PathValue("cartId")

	cartId, err := strconv.Atoi(pathCartId)
	if err != nil {
		return err
	}

	if _, err := s.cartRepo.GetById(cartId); err != nil {
		return err
	}

	pathItemId := req.PathValue("itemId")

	itemId, err := strconv.Atoi(pathItemId)
	if err != nil {
		return err
	}

	if err := s.itemRepo.Delete(cartId, itemId); err != nil {
		return err
	}

	return nil
}
