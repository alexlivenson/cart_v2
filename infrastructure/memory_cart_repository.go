package infrastructure

import (
	"errors"

	"github.com/alexlivenson/cart_v2/shopping"
)

// MemoryCartRepository is concrete implementation of domain.CartRepository
type MemoryCartRepository struct {
	cart map[string]shopping.Cart
}

// NewMemoryCartRepository returns new instance of MemoryCartRepository
func NewMemoryCartRepository() *MemoryCartRepository {
	return &MemoryCartRepository{cart: map[string]shopping.Cart{}}
}

// Add appends cart to in memory cart
func (mcr *MemoryCartRepository) Add(cart shopping.Cart) error {
	mcr.cart[cart.ID] = cart
	return nil
}

// Get retrieves in memory shopping cart
func (mcr *MemoryCartRepository) Get(id string) (*shopping.Cart, error) {
	cart, ok := mcr.cart[id]

	if !ok {
		return nil, errors.New("Cart does not exist in dictonary")
	}

	return &cart, nil
}

// Remove removes cart in memory shopping cart
func (mcr *MemoryCartRepository) Remove(id string) error {
	delete(mcr.cart, id)
	return nil
}
