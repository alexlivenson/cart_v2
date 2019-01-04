package memory

import (
	"errors"

	"github.com/alexlivenson/cart_v2/shopping"
)

// cartRepository is concrete implementation of domain.CartRepository
type cartRepository struct {
	cart map[string]shopping.Cart
}

// NewCartRepository returns new instance of MemoryCartRepository
func NewCartRepository() shopping.CartRepository {
	return &cartRepository{cart: map[string]shopping.Cart{}}
}

// Add appends cart to in memory cart
func (cr *cartRepository) Add(cart shopping.Cart) error {
	cr.cart[cart.ID] = cart
	return nil
}

// Get retrieves in memory shopping cart
func (cr *cartRepository) Get(id string) (*shopping.Cart, error) {
	cart, ok := cr.cart[id]

	if !ok {
		return nil, errors.New("Cart does not exist in dictonary")
	}

	return &cart, nil
}

// Remove removes cart in memory shopping cart
func (cr *cartRepository) Remove(id string) error {
	delete(cr.cart, id)
	return nil
}
