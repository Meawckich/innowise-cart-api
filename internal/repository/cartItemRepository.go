package repository

import (
	"cart-api/internal/pkg/model"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ItemRepository interface {
	Create(item model.ItemDto, cartId int) (model.CartItem, error)
	Delete(cartId, id int) error
}

type PostgresItemRepository struct {
	pool *sqlx.DB
}

func NewPostgresItemRepository(dbPool *sqlx.DB) *PostgresItemRepository {
	return &PostgresItemRepository{
		pool: dbPool,
	}
}

func (c *PostgresItemRepository) Create(item model.ItemDto, cartId int) (model.CartItem, error) {
	var count int
	if err := c.pool.Get(&count, "SELECT count(id) FROM carts WHERE id = $1", cartId); err != nil {
		return model.CartItem{}, err
	}

	if count == 0 {
		return model.CartItem{}, errors.New("cart id now found")
	}

	tx := c.pool.MustBegin()
	tx.MustExec("INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, $1, $2, $3);", item.Product, item.Quantity, cartId)
	err := tx.Commit()

	if err != nil {
		return model.CartItem{}, nil
	}

	row := c.pool.QueryRowx("SELECT id, product, quantity, cart_id FROM items ORDER BY id DESC LIMIT 1")

	var rowItem model.CartItem
	if err := row.StructScan(&rowItem); err != nil {
		return model.CartItem{}, err
	}

	return rowItem, nil
}

func (c *PostgresItemRepository) Delete(cartId, id int) error {
	tx := c.pool.MustBegin()
	tx.MustExec("DELETE FROM items WHERE id = $1 AND cart_id = $2", id, cartId)
	err := tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
