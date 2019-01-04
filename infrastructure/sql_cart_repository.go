package infrastructure

import (
	"database/sql"

	"github.com/alexlivenson/cart_v2/shopping"
)

// SQLCartRepository is concrete implementation of domain.CartRepository
type SQLCartRepository struct {
	db *sql.DB
}

// NewSQLCartRepository is constructor for SQLCartRepository
func NewSQLCartRepository(db *sql.DB) *SQLCartRepository {
	return &SQLCartRepository{db}
}

// Add ...
func (scr *SQLCartRepository) Add(cart shopping.Cart) error {
	return nil
}

// Get ...
func (scr *SQLCartRepository) Get(id string) (*shopping.Cart, error) {
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
func (scr *SQLCartRepository) Remove(id string) error {
	return nil
}
