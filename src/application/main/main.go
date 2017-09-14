package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/application"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/command"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
	"github.com/with-hindsight/chronicle/src/domain/aggregate"
)

func main() {

	events, err := application.Handle(
		application.IdentifierGenerator.New(),
		command.Create {
			ShopperId: application.IdentifierGenerator.New(),
		},
	)

	spew.Dump(events)
	spew.Dump(err)

	//handler := cart.Handlers[1]
	//spew.Dump(reflect.ValueOf(handler).Type().In(1).PkgPath())

	state := &cart.State{}

	p1 := aggregate.Projector{State: state}
	projector := cart.Projector{Projector: p1, State: state}

	spew.Dump(projector)

	event, err := application.EventGenerator.New(
		application.IdentifierGenerator.New(),
		application.IdentifierGenerator.New(),
		event.Created {
			ShopperId: application.IdentifierGenerator.New(),
		},
	)

	spew.Dump(event)

	projector.Apply(event)

	spew.Dump(projector)

}
