package memory

import (
	"testing"

	"github.com/alexlivenson/cart_v2/shopping"

	"github.com/alexlivenson/cart_v2/testutils"
)

func TestMemoryCartRepository(t *testing.T) {
	makeSUT := func() shopping.CartRepository {
		return NewCartRepository()
	}

	makeCart := func() shopping.Cart {
		want := shopping.Cart{
			ID: "a",
		}

		return want
	}

	t.Run("adds a cart to repo and retrieves it", func(t *testing.T) {
		// Given
		sut := makeSUT()
		cart := makeCart()

		// When adding to cart
		err := sut.Add(cart)

		// Then no error
		testutils.Ok(t, err)

		// When retrieving cart
		got, err := sut.Get(cart.ID)

		// Then no error
		testutils.Ok(t, err)
		testutils.AssertEquals(t, cart, *got)
	})

	t.Run("adds a cart to repo and removes it", func(t *testing.T) {
		// Given
		sut := makeSUT()
		cart := makeCart()

		// When adding to cart
		err := sut.Add(cart)

		// Then no error
		testutils.Ok(t, err)

		// When retrieving cart
		err = sut.Remove(cart.ID)

		// Then no error
		testutils.Ok(t, err)

		// When retrieving cart again, we get error
		_, err = sut.Get(cart.ID)

		// Then
		testutils.NotOk(t, err)
	})
}
