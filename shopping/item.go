package shopping

// ItemError ...
type ItemError string

// ErrAmountMustBePositive ...
const ErrAmountMustBePositive = ItemError("Amount must be positive")

func (e ItemError) Error() string {
	return e.Error()
}

// ItemDetail is simple value object to represent an Item
type ItemDetail struct {
	productID string
	price     Price
	amount    int
}

// Item represents an product in a cart
type Item struct {
	productID string
	amount    int
	unitPrice Price
}

// NewItem returns item or error
func NewItem(productID string, unitPrice Price, amount int) (*Item, error) {
	if err := checkAmount(amount); err != nil {
		return nil, err
	}

	return &Item{
		productID: productID,
		unitPrice: unitPrice,
		amount:    amount,
	}, nil
}

// Add ...
func (i *Item) Add(amount int) error {
	if err := checkAmount(amount); err != nil {
		return err
	}

	i.amount += amount
	return nil
}

// ChangeAmount ...
func (i *Item) ChangeAmount(amount int) error {
	if err := checkAmount(amount); err != nil {
		return err
	}

	i.amount = amount
	return nil
}

// CalculatePrice ...
func (i *Item) CalculatePrice() Price {
	return i.unitPrice.Multiply(i.amount)
}

// ToDetail ...t
func (i *Item) ToDetail() ItemDetail {
	return ItemDetail{
		productID: i.productID,
		price:     i.unitPrice,
		amount:    i.amount,
	}
}

func checkAmount(amount int) error {
	if amount <= 0 {
		return ErrAmountMustBePositive
	}

	return nil
}
