package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/application"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/command"
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

}
