package main

import (
	"github.com/with-hindsight/chronicle/src/application/bus"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/infrastructure"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/command"
	"github.com/davecgh/go-spew/spew"
)

var IdentifierGenerator = infrastructure.IdentifierGenerator

var Bus = bus.New(
	infrastructure.AggregateRepository,
	infrastructure.CommandGenerator,
)

func main() {

	aggregate_id := IdentifierGenerator.New()

	events, err := Bus.Dispatch(
		aggregate_id,
		command.Create {
			ShopperId: IdentifierGenerator.New(),
		},
	)

	spew.Dump(events)
	spew.Dump(err)
}

