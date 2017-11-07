package workflow

import (
	"github.com/with-hindsight/chronicle/domain"
	"github.com/with-hindsight/chronicle/application/bus"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/command"
)

type Sample struct {

	bus *bus.Bus
}

func (self *Sample) Handle(evt domain.Event) error {

	switch event := evt.Payload().(type) {

	case event.Created:

		self.bus.Dispatch(
			evt.Identifier(),
			command.Create{
				ShopperId: event.ShopperId,
			},
		)


	}

	return nil
}
