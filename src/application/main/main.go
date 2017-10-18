package main

import (
	"github.com/with-hindsight/chronicle/application"
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

	infrastructure.Boot()

	////////////////
	////////////////
	////////////////

	aggregate_id := IdentifierGenerator.New()

	commands := []application.Command{
		command.Create {
			ShopperId: IdentifierGenerator.New(),
		},
	}

	for i := 0 ; i < 2000 ; i++ {

		commands = append(
			commands,
			command.AddProduct {
				Product: entity.Product{
					Id: IdentifierGenerator.New(),
					Quantity: 100,
				},
			},
		)
	}

	_,_,err := Bus.DispatchMany(
		aggregate_id,
		commands,
	)

	spew.Dump(err)

	_,_,err = Bus.DispatchMany(
		aggregate_id,
		[]application.Command{
			command.AddProduct {
				Product: entity.Product{
					Id: IdentifierGenerator.New(),
					Quantity: 5,
				},
			},
		},
	)

	spew.Dump(err)

	_,_,err = Bus.DispatchMany(
		aggregate_id,
		[]application.Command{
			command.Checkout {},
		},
	)

	spew.Dump(err)

	/*
	_,_,err = Bus.DispatchMany(
		aggregate_id,
		[]application.Command{
			command.Checkout {},
		},
	)

	spew.Dump(err)

	*/

	//spew.Dump(infrastructure.AggregateRepository.InstanceSnapshots())
	//spew.Dump(infrastructure.EventLog)

	/*
	fmt.Println("Commands")
	result, err := json.MarshalIndent(commands, ""," ")
	fmt.Println(string(result))

	fmt.Println("Events")
	result, err = json.MarshalIndent(events, ""," ")
	fmt.Println(string(result))
	*/
}
