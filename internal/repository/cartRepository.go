package repository

import (
	"cart-api/internal/pkg/model"
	"database/sql"
	"errors"
)

type ICartRepository interface {
	GetById(id int) (model.Cart, error)
	GetAll() ([]model.Cart, error)
	Create() (model.Cart, error)
	Delete(id int) error
}

type PostgresCartRepository struct {
	pool *sql.DB
}

func NewPostgresCartRepository(dbPool *sql.DB) *PostgresCartRepository {
	return &PostgresCartRepository{
		pool: dbPool,
	}
}

func (c *PostgresCartRepository) GetById(id int) (model.Cart, error) {
	var cartId int
	row := c.pool.QueryRow("SELECT id FROM carts WHERE id = $1 LIMIT 1", id)

	if err := row.Scan(&cartId); err != nil {
		return model.Cart{}, err
	}

	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return model.Cart{}, errors.New("cart id now found")
		}
		return model.Cart{}, row.Err()
	}

	cart := model.Cart{}

	itemRows, err := c.pool.Query("SELECT id, product, quantity, cart_id FROM items where cart_id = $1", cartId)

	if err != nil {
		return model.Cart{}, err
	}

	defer itemRows.Close()

	var items []model.CartItem
	for itemRows.Next() {
		var item model.CartItem

		if err := itemRows.Scan(&item.Id, &item.Product, &item.Quantity, &item.Cart_id); err != nil {
			return model.Cart{}, err
		}

		items = append(items, item)
	}

	cart.Id = id

	if len(items) == 0 {
		cart.Items = make([]model.CartItem, 0)
	} else {
		cart.Items = items
	}

	return cart, nil
}

func (c *PostgresCartRepository) GetAll() ([]model.Cart, error) {
	var count int
	row := c.pool.QueryRow("SELECT count(id) FROM carts;")
	if row.Err() != nil {
		if row.Err() == sql.ErrNoRows {
			return []model.Cart{}, errors.New("carts not found")
		}
		return []model.Cart{}, row.Err()
	}

	err := row.Scan(&count)
	if err != nil {
		return []model.Cart{}, err
	}

	rows, err := c.pool.Query("Select id from carts")

	if err != nil {
		return []model.Cart{}, err
	}

	defer rows.Close()

	carts := make([]model.Cart, 0, count)

	//remove append
	for rows.Next() {
		cart := &model.Cart{}

		if err := rows.Scan(&cart.Id); err != nil {
			return []model.Cart{}, err
		}

		itemRows, err := c.pool.Query("SELECT id, product, quantity, cart_id FROM items where cart_id = $1", cart.Id)
		if err != nil {
			break
		}

		defer itemRows.Close()

		var items []model.CartItem

		for itemRows.Next() {

			var item model.CartItem
			if err := itemRows.Scan(&item.Id, &item.Product, &item.Quantity, &item.Cart_id); err != nil {
				break
			}
			items = append(items, item)
		}

		if len(items) == 0 {
			cart.Items = make([]model.CartItem, 0)
		} else {
			cart.Items = items

		}

		carts = append(carts, *cart)
	}

	return carts, nil

}

func (c *PostgresCartRepository) Create() (model.Cart, error) {
	_, err := c.pool.Exec("INSERT INTO carts DEFAULT VALUES")
	if err != nil {
		return model.Cart{}, err
	}

	row := c.pool.QueryRow("SELECT id FROM carts ORDER BY id DESC LIMIT 1")

	var insertedId int
	if err := row.Scan(&insertedId); err != nil {
		return model.Cart{}, err
	}

	items := make([]model.CartItem, 0)

	return model.Cart{Id: int(insertedId), Items: items}, nil
}

func (c *PostgresCartRepository) Delete(id int) error {
	_, err := c.pool.Exec("DELETE FROM carts WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (c *PostgresCartRepository) Update(model.Cart) error {
	return nil
}
