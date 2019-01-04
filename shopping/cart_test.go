package shopping

import (
	"testing"

	"github.com/alexlivenson/cart_v2/testutils"
)

func TestCart(t *testing.T) {

	makeSUT := func(id string, items []*Item) Cart {
		return Cart{ID: id, Items: items}
	}

	t.Run("calcuate empty cart", func(t *testing.T) {
		sut := makeSUT("1", []*Item{})
		want := CartDetail{
			items:      []ItemDetail{},
			totalPrice: Price{Value: 0.0},
		}

		testutils.AssertEquals(t, want, sut.Calculate())
	})

	t.Run("add single product to empty", func(t *testing.T) {
		price := Price{Value: 10.0}

		sut := makeSUT("1", []*Item{})
		sut.Add("a", price, 1)

		want := CartDetail{
			totalPrice: price,
			items: []ItemDetail{
				ItemDetail{
					productID: "a",
					price:     price,
					amount:    1,
				},
			},
		}

		testutils.AssertEquals(t, want, sut.Calculate())
	})

	t.Run("add two different products", func(t *testing.T) {
		price1 := Price{Value: 10.0}
		price2 := Price{Value: 20.0}

		sut := makeSUT("1", []*Item{})
		sut.Add("a", price1, 1)
		sut.Add("b", price2, 2)

		want := CartDetail{
			totalPrice: Price{Value: 50.0},
			items: []ItemDetail{
				ItemDetail{
					productID: "a",
					price:     price1,
					amount:    1,
				},
				ItemDetail{
					productID: "b",
					price:     price2,
					amount:    2,
				},
			},
		}

		testutils.AssertEquals(t, want, sut.Calculate())
	})

	t.Run("add same product increment amount only", func(t *testing.T) {
		sut := makeSUT("1", []*Item{})
		sut.Add("a", Price{10.0}, 1)
		sut.Add("a", Price{0.0}, 1)

		want := CartDetail{
			totalPrice: Price{20.0},
			items: []ItemDetail{
				ItemDetail{
					productID: "a",
					price:     Price{10},
					amount:    2,
				},
			},
		}

		testutils.AssertEquals(t, want, sut.Calculate())
	})

	t.Run("remove non existing item from empty cart", func(t *testing.T) {
		sut := makeSUT("1", []*Item{})
		err := sut.Remove("x")
		testutils.NotOk(t, err)
	})

	t.Run("remove non existing item from non-empty cart", func(t *testing.T) {
		sut := makeSUT("1", []*Item{})
		sut.Add("a", Price{10.0}, 1)

		err := sut.Remove("x")
		testutils.NotOk(t, err)
	})

	t.Run("remove product from cart", func(t *testing.T) {
		sut := makeSUT("1", []*Item{})
		sut.Add("a", Price{10.0}, 1)
		sut.Add("b", Price{20.0}, 2)

		sut.Remove("a")

		want := CartDetail{
			totalPrice: Price{40.0},
			items: []ItemDetail{
				ItemDetail{
					productID: "b",
					price:     Price{20},
					amount:    2,
				},
			},
		}

		testutils.AssertEquals(t, want, sut.Calculate())
	})

	t.Run("change amount on non-existing", func(t *testing.T) {
		sut := makeSUT("1", []*Item{})
		sut.Add("a", Price{10.0}, 1)

		err := sut.ChangeAmount("x", 2)

		testutils.NotOk(t, err)
	})

	t.Run("change amount", func(t *testing.T) {
		sut := makeSUT("1", []*Item{})
		sut.Add("a", Price{10.0}, 1)

		err := sut.ChangeAmount("a", 5)
		want := CartDetail{
			totalPrice: Price{50.0},
			items: []ItemDetail{
				ItemDetail{
					productID: "a",
					price:     Price{10},
					amount:    5,
				},
			},
		}

		testutils.Ok(t, err)
		testutils.AssertEquals(t, want, sut.Calculate())
	})
}
