package shopping

// CartRepository repo layer for Cart
type CartRepository interface {
	Add(cart Cart) error
	Get(id string) (*Cart, error)
	Remove(id string) error
}
