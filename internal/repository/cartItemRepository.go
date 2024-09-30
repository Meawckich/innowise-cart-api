package repository

import (
	"cart-api/internal/pkg/model"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type ItemRepository interface {
	Create(item model.ItemDto, cartId int) (model.CartItem, error)
	Delete(cartId, id int) error
}

type PostgresItemRepository struct {
	pool *sql.DB
}

func NewPostgresItemRepository(dbPool *sql.DB) *PostgresItemRepository {
	return &PostgresItemRepository{
		pool: dbPool,
	}
}

func (c *PostgresItemRepository) Create(item model.ItemDto, cartId int) (model.CartItem, error) {
	var count int
	row := c.pool.QueryRow("SELECT count(id) FROM carts WHERE id = $1", cartId)

	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return model.CartItem{}, errors.New("cart id now found")
		}
		return model.CartItem{}, row.Err()
	}

	if err := row.Scan(&count); err != nil {
		return model.CartItem{}, err
	}

	_, err := c.pool.Exec("INSERT INTO items (id, product, quantity, cart_id) VALUES (DEFAULT, $1, $2, $3);", item.Product, item.Quantity, cartId)
	if err != nil {
		return model.CartItem{}, nil
	}

	var insertedId int

	row = c.pool.QueryRow("SELECT id FROM items ORDER BY id DESC LIMIT 1")

	if err := row.Scan(&insertedId); err != nil {
		return model.CartItem{}, nil
	}

	var insertedItem model.CartItem

	row = c.pool.QueryRow("SELECT id, product, quantity, cart_id FROM items where id = $1", insertedId)

	if err := row.Scan(&insertedItem.Id, &insertedItem.Product, &insertedItem.Quantity, &insertedItem.Cart_id); err != nil {
		return model.CartItem{}, err
	}

	return insertedItem, nil
}

func (c *PostgresItemRepository) Delete(cartId, id int) error {
	_, err := c.pool.Exec("DELETE FROM items WHERE id = $1 AND cart_id = $2", id, cartId)

	if err != nil {
		return err
	}

	return nil
}
