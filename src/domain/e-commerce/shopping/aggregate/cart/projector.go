package cart

import (
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
	"github.com/with-hindsight/chronicle/src/domain/context"
	"github.com/with-hindsight/chronicle/src/domain/context/aggregate"
)

type Projector struct {

	State *State
}

func (self *Projector) Handle(e *context.Event) error {

	switch event := e.Payload.(type) {

	case event.Created:

		self.State.IsCreated = true
		self.State.ShopperId = event.ShopperId

		return nil

	case event.ProductAdded:

		self.State.Products.Add(event.Product.Id)

		return nil

	case event.ProductRemoved:

		self.State.Products.Remove(event.ProductId)

		return nil

	case event.CheckedOut:

		self.State.IsCheckedOut = true

		return nil
	}

	return aggregate.ErrProjectorHandlerNotExists
}
