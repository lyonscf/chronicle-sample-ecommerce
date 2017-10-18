package infrastructure

import (
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/command"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart"
	"github.com/with-hindsight/chronicle/src/domain/context/aggregate"
	"github.com/with-hindsight/chronicle/src/domain/context"
	"github.com/with-hindsight/chronicle/src/domain/mapper"
	"github.com/with-hindsight/chronicle/src/domain"
	"reflect"
)

var CartAggregate = &aggregate.Archetype {
	EventGenerator: EventGenerator,
	Checker: &StubChecker{},
	AggregateContext: &context.AggregateContext{
		Id: domain.IdentifierString("f5e5d99a"),
		Name: "cart",
		Context: &context.Context{
			Id: domain.IdentifierString("958bbea6"),
			Name: "shopping",
			Players: []context.Player{},
			Domain: &domain.Context{
				Id: domain.IdentifierString("b9a185c3"),
				Name: "e-commerce",
			},
		},
	},
	ArchetypeContext: aggregate.ArchetypeContext{
		Invariant: reflect.TypeOf(cart.Invariant{}),
		ProjectorHandler: reflect.TypeOf(cart.Projector{}),
		State: reflect.TypeOf(cart.State{}),
	},
}

var Commands = []*context.ObjectContext{
	&context.ObjectContext{
		Id: domain.IdentifierString("5bc8ca6b"),
		Name: "create",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(command.Create{}),
		AggregateArchetype: CartAggregate,
	},
	&context.ObjectContext{
		Id: domain.IdentifierString("08d14870"),
		Name: "add-product",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(command.AddProduct{}),
		AggregateArchetype: CartAggregate,
	},
	&context.ObjectContext{
		Id: domain.IdentifierString("1e420e4e"),
		Name: "change-product-quantity",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(command.ChangeProductQuantity{}),
		AggregateArchetype: CartAggregate,
	},
	&context.ObjectContext{
		Id: domain.IdentifierString("c308783c"),
		Name: "remove-product",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(command.RemoveProduct{}),
		AggregateArchetype: CartAggregate,
	},
	&context.ObjectContext{
		Id: domain.IdentifierString("6bec2657"),
		Name: "checkout",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(command.Checkout{}),
		AggregateArchetype: CartAggregate,
	},
}

var Events = []*context.ObjectContext{
	&context.ObjectContext{
		Id: domain.IdentifierString("1edc201b"),
		Name: "created",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(event.Created{}),
		AggregateArchetype: CartAggregate,
	},
	&context.ObjectContext{
		Id: domain.IdentifierString("f1b88fd9"),
		Name: "empty",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(event.Empty{}),
		AggregateArchetype: CartAggregate,
	},
	&context.ObjectContext{
		Id: domain.IdentifierString("c952cb41"),
		Name: "full",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(event.Full{}),
		AggregateArchetype: CartAggregate,
	},
	&context.ObjectContext{
		Id: domain.IdentifierString("72784981"),
		Name: "product-added",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(event.ProductAdded{}),
		AggregateArchetype: CartAggregate,
	},
	&context.ObjectContext{
		Id: domain.IdentifierString("89a0688f"),
		Name: "product-quantity-changed",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(event.ProductQuantityChanged{}),
		AggregateArchetype: CartAggregate,
	},
	&context.ObjectContext{
		Id: domain.IdentifierString("17d8c65e"),
		Name: "product-removed",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(event.ProductRemoved{}),
		AggregateArchetype: CartAggregate,
	},
	&context.ObjectContext{
		Id: domain.IdentifierString("2481cb7f"),
		Name: "checked-out",
		Version: 1,
		ObjectArchetype: reflect.TypeOf(event.CheckedOut{}),
		AggregateArchetype: CartAggregate,
	},
}

var DomainMap = mapper.NewMap(Commands, Events)
