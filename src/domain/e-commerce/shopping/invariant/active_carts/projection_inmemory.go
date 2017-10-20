package active_carts

import (
	"github.com/with-hindsight/chronicle/domain"
	"github.com/emirpasic/gods/sets"
	"github.com/emirpasic/gods/maps"
)

type InMemoryProjection struct {

	shoppers sets.Set
	carts maps.Map
}

func (self *InMemoryProjection) Exists(shopper domain.Identifier) bool {

	return self.shoppers.Contains(shopper)
}

func (self *InMemoryProjection) Add(cart domain.Identifier, shopper domain.Identifier) error {

	if self.Exists(shopper) {
		return ErrShopperHasActiveCart
	}

	self.carts.Put(
		cart,
		shopper,
	)

	self.shoppers.Add(shopper)

	return nil
}

func (self *InMemoryProjection) Remove(cart domain.Identifier) error {

	shopper, exists := self.carts.Get(cart)

	if !exists {
		return ErrCartDoesNotExists
	}

	self.shoppers.Remove(shopper)
	self.carts.Remove(cart)

	return nil
}

