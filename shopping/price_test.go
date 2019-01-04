package shopping

import (
	"fmt"
	"testing"

	"github.com/alexlivenson/cart_v2/testutils"
)

func TestPrice(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		a := Price{Value: 10.0}
		b := Price{Value: 0.5}

		got := a.Add(b)
		want := Price{Value: 10.5}

		testutils.AssertEquals(t, want, got)
	})

	t.Run("Multiply", func(t *testing.T) {
		a := Price{Value: 10.0}

		got := a.Multiply(2)
		want := Price{Value: 20}

		testutils.AssertEquals(t, want, got)
	})
}
func ExamplePrice_Add() {
	a := Price{Value: 10.0}
	b := Price{Value: 0.5}

	fmt.Println(a.Add(b).Value)
	// Output: 10.5
}

func ExamplePrice_Multiply() {
	a := Price{Value: 10.0}

	fmt.Println(a.Multiply(2).Value)
	// Output: 20
}

func TestSum(t *testing.T) {
	prices := []Price{
		Price{Value: 9.0},
		Price{Value: 0.7},
		Price{Value: 0.3},
	}

	got := Sum(prices)
	want := Price{Value: 10}

	testutils.AssertEquals(t, got, want)
}
