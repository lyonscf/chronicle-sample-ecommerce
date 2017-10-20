package cart

import (
	"github.com/with-hindsight/chronicle/domain"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
)

type Projector struct {

	State *State
}

func (self *Projector) Handle(evt domain.Event) {

	switch event := evt.Payload().(type) {

	case event.Created:

		self.State.IsCreated = true
		self.State.ShopperId = event.ShopperId

	case event.ProductAdded:

		self.State.Products.Add(event.Product.Id)

	case event.ProductRemoved:

		self.State.Products.Remove(event.ProductId)

	case event.CheckedOut:

		self.State.IsCheckedOut = true
	}
}
