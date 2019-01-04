package sql

import (
	"database/sql"

	"github.com/alexlivenson/cart_v2/shopping"
)

// cartRepository is concrete implementation of domain.CartRepository
type cartRepository struct {
	db *sql.DB
}

// NewCartRepository is constructor for SQLCartRepository
func NewCartRepository(db *sql.DB) shopping.CartRepository {
	return &cartRepository{db}
}

// Add ...
func (scr *cartRepository) Add(cart shopping.Cart) error {
	return nil
}

// Get ...
func (scr *cartRepository) Get(id string) (*shopping.Cart, error) {
	row := scr.db.QueryRow(
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
func (scr *cartRepository) Remove(id string) error {
	return nil
}
