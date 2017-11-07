package main

import (
	"github.com/with-hindsight/chronicle/application/bus"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/infrastructure"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/command"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/entity"
	"github.com/davecgh/go-spew/spew"
)

var IdentifierGenerator = infrastructure.IdentifierGenerator

var Bus = bus.New(
	infrastructure.IdentifierGenerator,
	infrastructure.AggregateRepository,
	infrastructure.CommandGenerator,
)

func main() {

	var request *bus.Request
	var response *bus.Response

	infrastructure.Boot()

	////////////////
	////////////////
	////////////////

	request = bus.NewRequest(
		IdentifierGenerator.New(),
		command.Create {
			ShopperId: IdentifierGenerator.New(),
		},
	)

	for i := 0 ; i < 2000 ; i++ {

		request.Commands = append(request.Commands,
			command.AddProduct {
				Product: entity.Product{
					Id: IdentifierGenerator.New(),
					Quantity: 100,
				},
			},
		)
	}

	response = Bus.Dispatch(request).Wait()

	spew.Dump(response.Error)

	response = Bus.Dispatch(
		bus.NewRequest(
			response.AggregateId(),
			command.AddProduct {
				Product: entity.Product{
					Id: IdentifierGenerator.New(),
					Quantity: 5,
				},
			},
		),
	).Wait()

	spew.Dump(response.Error)

	response = Bus.Dispatch(
		bus.NewRequest(
			response.AggregateId(),
			command.Checkout {},
		),
	).Wait()

	spew.Dump(response.Error)

	spew.Dump(infrastructure.ProjectionRepository.Snapshots())
}
