package memory

import (
	"github.com/with-hindsight/chronicle/domain"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/invariant/active_carts"
)

type Projection struct {

	store *Store
}

func (self *Projection) Reset() error {

	return self.store.Reset()
}

func (self *Projection) Add(cart domain.Identifier, shopper domain.Identifier) error {

	if self.store.shoppers.Contains(shopper) {
		return active_carts.ErrShopperHasActiveCart
	}

	self.store.carts.Put(
		cart,
		shopper,
	)

	self.store.shoppers.Add(shopper)

	return nil
}

func (self *Projection) Remove(cart domain.Identifier) error {

	shopper, exists := self.store.carts.Get(cart)

	if !exists {
		return active_carts.ErrCartDoesNotExist
	}

	self.store.shoppers.Remove(shopper)
	self.store.carts.Remove(cart)

	return nil
}

func NewProjection(store *Store) *Projection {

	return &Projection{
		store: store,
	}
}
