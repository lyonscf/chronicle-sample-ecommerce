package active_carts

import (
	"github.com/with-hindsight/chronicle/domain/projection"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
)

type Projector struct {

	Projection Projection
}

func (self *Projector) Handle(evt projection.Event) error {

	switch event := evt.Payload().(type) {

	case event.Created:

		return self.Projection.Add(
			evt.AggregateId(),
			event.ShopperId,
		)

	case event.CheckedOut:

		return self.Projection.Remove(evt.AggregateId())
	}

	return projection.ErrHandlerNotExists
}
