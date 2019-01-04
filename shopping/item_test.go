package shopping

import (
	"testing"

	"github.com/alexlivenson/cart_v2/testutils"
)

func TestItem(t *testing.T) {

	productID := "x"
	price := Price{Value: 0.5}
	amount := 2

	makeSUT := func() *Item {
		item, _ := NewItem(productID, price, amount)
		return item
	}

	t.Run("ToDetail", func(t *testing.T) {
		sut := makeSUT()

		want := ItemDetail{
			productID: productID,
			price:     price,
			amount:    amount,
		}

		got := sut.ToDetail()

		testutils.AssertEquals(t, want, got)
	})

	t.Run("Add", func(t *testing.T) {
		sut := makeSUT()
		sut.Add(5)

		want := ItemDetail{
			productID: productID,
			price:     price,
			amount:    7,
		}

		testutils.AssertEquals(t, want, sut.ToDetail())
	})

	t.Run("cannot Add negative numbers", func(t *testing.T) {
		sut := makeSUT()
		err := sut.Add(-2)

		testutils.NotOk(t, err)
	})

	t.Run("cannot Add 0", func(t *testing.T) {
		sut := makeSUT()
		err := sut.Add(0)

		testutils.NotOk(t, err)
	})

	t.Run("ChangeAmount", func(t *testing.T) {
		sut := makeSUT()
		sut.ChangeAmount(1)

		want := ItemDetail{
			productID: productID,
			price:     price,
			amount:    1,
		}

		testutils.AssertEquals(t, want, sut.ToDetail())
	})

	t.Run("cannot ChangeAmount to negative", func(t *testing.T) {
		sut := makeSUT()
		err := sut.ChangeAmount(-1)

		testutils.NotOk(t, err)
	})

	t.Run("cannot ChangeAmount to 0", func(t *testing.T) {
		sut := makeSUT()
		err := sut.ChangeAmount(0)

		testutils.NotOk(t, err)
	})

	t.Run("initial amount cannot be negative", func(t *testing.T) {
		_, err := NewItem(productID, price, -1)
		testutils.NotOk(t, err)
	})

	t.Run("initial amount cannot be negative", func(t *testing.T) {
		_, err := NewItem(productID, price, -1)
		testutils.NotOk(t, err)
	})

	t.Run("CalculateTotalPrice", func(t *testing.T) {
		sut, _ := NewItem(productID, Price{Value: 4}, 5)

		got := sut.CalculatePrice()
		want := Price{Value: 20}

		testutils.AssertEquals(t, want, got)
	})
}
