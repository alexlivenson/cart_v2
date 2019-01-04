package shopping

// CartError ...
type CartError string

// ErrProductNotInCart ...
const ErrProductNotInCart = CartError("Product not it cart")

func (e CartError) Error() string {
	return e.Error()
}

// CartDetail give summary of Cart
type CartDetail struct {
	items      []ItemDetail
	totalPrice Price
}

// Cart represent the holder for items
type Cart struct {
	ID    string
	Items []*Item
}

// Add takes a product and adds to cart
func (c *Cart) Add(productID string, unitPrice Price, amount int) error {
	if item, err := c.find(productID); err != nil {
		c.Items = append(c.Items, &Item{productID, amount, unitPrice})
	} else {
		return item.Add(amount)
	}
	return nil
}

// Calculate returns the CartDetail
func (c *Cart) Calculate() CartDetail {
	itemDetails := []ItemDetail{}
	prices := []Price{}

	for _, i := range c.Items {
		itemDetails = append(itemDetails, i.ToDetail())
		prices = append(prices, i.CalculatePrice())
	}

	totalPrice := Sum(prices)

	return CartDetail{
		items:      itemDetails,
		totalPrice: totalPrice,
	}
}

// Remove removes item from shopping cart
func (c *Cart) Remove(productID string) error {
	filteredItems := []*Item{}

	for _, i := range c.Items {
		if i.productID != productID {
			filteredItems = append(filteredItems, i)
		}
	}

	if len(filteredItems) == len(c.Items) {
		return ErrProductNotInCart
	}

	c.Items = filteredItems
	return nil
}

// ChangeAmount updates the amount of item
func (c *Cart) ChangeAmount(productID string, amount int) error {
	item, err := c.find(productID)
	if err != nil {
		return err
	}
	item.ChangeAmount(amount)
	return nil
}

func (c *Cart) find(productID string) (*Item, error) {
	for _, i := range c.Items {
		if i.productID == productID {
			return i, nil
		}
	}

	return nil, ErrProductNotInCart
}
