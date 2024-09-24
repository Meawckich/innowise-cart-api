package repository

import (
	"cart-api/internal/pkg/common/model"
	"errors"

	"github.com/jmoiron/sqlx"
)

const IdAfterInsertMock int = 10

type ICartRepository interface {
	GetById(id int) (*model.Cart, error)
	GetAll() ([]model.Cart, error)
	Create() (model.Cart, error)
	Delete(id int) error
}

type PostgresCartRepository struct {
	pool *sqlx.DB
}

func NewPostgresCartRepository(dbPool *sqlx.DB) *PostgresCartRepository {
	return &PostgresCartRepository{
		pool: dbPool,
	}
}

func (c *PostgresCartRepository) GetById(id int) (*model.Cart, error) {
	existsCartIdRow := c.pool.QueryRowx("SELECT id FROM carts WHERE id = $1 LIMIT 1", id)
	var cartId int

	if err := existsCartIdRow.Scan(&cartId); err != nil {
		return &model.Cart{}, err
	}

	if cartId == 0 {
		return &model.Cart{}, errors.New("cannot find cart by given id")
	}

	cart := model.Cart{}

	itemRows, err := c.pool.Queryx("SELECT * FROM items where cart_id = $1", cartId)

	if err != nil {
		return &model.Cart{}, err
	}

	defer itemRows.Close()

	var items []model.CartItem
	for itemRows.Next() {
		var item model.CartItem

		if err := itemRows.StructScan(&item); err != nil {
			return &model.Cart{}, err
		}

		items = append(items, item)
	}

	cart.Id = id

	if len(items) == 0 {
		cart.Items = make([]model.CartItem, 0)
	} else {
		cart.Items = items
	}

	return &cart, nil
}

func (c *PostgresCartRepository) GetAll() ([]model.Cart, error) {
	var count int
	if err := c.pool.Get(&count, "SELECT count(id) FROM carts;"); err != nil {
		return []model.Cart{}, err
	}

	rows, err := c.pool.Queryx("Select id from carts")

	if err != nil {
		return []model.Cart{}, err
	}

	defer rows.Close()

	carts := make([]model.Cart, 0, count)

	//remove append
	for rows.Next() {
		cart := &model.Cart{}

		if err := rows.StructScan(cart); err != nil {
			return []model.Cart{}, err
		}

		itemRows, err := c.pool.Queryx("SELECT id, product, quantity, cart_id FROM items where cart_id = $1", cart.Id)
		if err != nil {
			break
		}

		defer itemRows.Close()

		var items []model.CartItem

		for itemRows.Next() {

			var item model.CartItem
			if err := itemRows.StructScan(&item); err != nil {
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
	tx := c.pool.MustBegin()
	tx.MustExec("INSERT INTO carts DEFAULT VALUES")
	if err := tx.Commit(); err != nil {
		return model.Cart{}, err
	}

	var createdId int
	if err := c.pool.QueryRowx("SELECT id FROM carts ORDER BY id DESC LIMIT 1").Scan(&createdId); err != nil {
		return model.Cart{}, err
	}
	items := make([]model.CartItem, 0)

	return model.Cart{Id: createdId, Items: items}, nil
}

func (c *PostgresCartRepository) Delete(id int) error {
	tx := c.pool.MustBegin()
	tx.MustExec("DELETE FROM carts WHERE id = $1", id)
	err := tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (c *PostgresCartRepository) Update(model.Cart) error {
	return nil
}
