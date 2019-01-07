package sql

import (
	"database/sql"

	"github.com/alexlivenson/cart_v2/shopping"
)

// cartRepository is concrete implementation of shopping.CartRepository
type cartRepository struct {
	db *sql.DB
}

// NewCartRepository is constructor for SQLCartRepository
func NewCartRepository(db *sql.DB) shopping.CartRepository {
	return &cartRepository{db}
}

// Add ...
func (cr *cartRepository) Add(cart shopping.Cart) error {
	return nil
}

// Get ...
func (cr *cartRepository) Get(id string) (*shopping.Cart, error) {
	row := cr.db.QueryRow(
		`SELECT
    			c.id ID, i.product_id productID, i.amount
			FROM
				cart c
			INNER JOIN
				c.item_id = i.id
    		WHERE
    			id = (?);`,
		id,
	)
	var cart shopping.Cart
	err := row.Scan(
		&cart.ID,
		&cart.Items,
	)
	if err != nil {
		return &cart, err
	}
	return &cart, nil
}

// Remove ...
func (cr *cartRepository) Remove(id string) error {
	return nil
}
