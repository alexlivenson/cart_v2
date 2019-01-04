package main

import (
	"fmt"

	"github.com/alexlivenson/cart_v2/shopping"

	"github.com/alexlivenson/cart_v2/memory"
)

var (
	cartRepo shopping.CartRepository
)

func main() {
	fmt.Println("nice")
}

func assemble() {
	cartRepo = memory.NewCartRepository()
}
