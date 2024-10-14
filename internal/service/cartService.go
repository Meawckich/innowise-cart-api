package service

import (
	"net/http"
	"strconv"

	"cart-api/internal/pkg/model"
	"cart-api/internal/repository"
)

type CartService struct {
	repo repository.ICartRepository
}

func NewCartService(repo repository.ICartRepository) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) CreateCart() (model.Cart, error) {
	newCart, err := s.repo.Create()
	if err != nil {
		return newCart, err
	}

	return newCart, nil
}

func (s *CartService) GetCarts() ([]model.Cart, error) {
	carts, err := s.repo.GetAll()
	if err != nil {
		return carts, err
	}

	return carts, nil
}

func (s *CartService) ViewCart(req *http.Request) (model.Cart, error) {
	cart := model.Cart{}

	idPath := req.PathValue("id")

	id, err := strconv.Atoi(idPath)
	if err != nil {
		return cart, err
	}

	cart, err = s.repo.GetById(id)

	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (s *CartService) DeleteCart(req *http.Request) error {
	pathId := req.PathValue("id")

	id, err := strconv.Atoi(pathId)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
