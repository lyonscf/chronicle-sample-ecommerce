package cart

import (
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
	"github.com/with-hindsight/chronicle/src/domain/context"
	"github.com/with-hindsight/chronicle/src/domain/context/aggregate"
)

type Projector struct {}

func (self *Projector) Handle(e *context.Event, s aggregate.State) error {

	state := s.(*State)

	switch event := e.Payload.(type) {

	case event.Created:

		state.IsCreated = true
		state.ShopperId = event.ShopperId

		return nil

	case event.ProductAdded:

		state.Products.Add(event.Product.Id)

		return nil

	case event.ProductRemoved:

		state.Products.Remove(event.ProductId)

		return nil

	case event.CheckedOut:

		state.IsCheckedOut = true

		return nil
	}

	return aggregate.ErrProjectorHandlerNotExists
}
