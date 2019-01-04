package shopping

// Price represents the amount a product costs
type Price struct {
	Value float64
}

// Add returns a Price with combine value of adder and this
func (p Price) Add(adder Price) Price {
	newValue := p.Value + adder.Value
	return Price{Value: newValue}
}

// Multiply returns a Price multiplied by the num amount
func (p Price) Multiply(num int) Price {
	return Price{Value: p.Value * float64(num)}
}

// Sum takes an array of prices and sums them uptPr
func Sum(prices []Price) Price {
	sum := float64(0)

	for _, price := range prices {
		sum += price.Value
	}

	return Price{Value: sum}
}
