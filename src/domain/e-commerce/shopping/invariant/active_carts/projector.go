package active_carts

import (
	"github.com/with-hindsight/chronicle/src/domain/projection"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
)

type Projector struct {

	Projection Projection
}

func (self *Projector) Handle(evt projection.Event) error {

	switch evt.Payload().(type) {

	case event.Created:

		return self.Projection.Add(evt.AggregateId())

	case event.CheckedOut:

		return self.Projection.Remove(evt.AggregateId())
	}

	return projection.ErrHandlerNotExists
}
