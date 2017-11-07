package infrastructure

import (
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/command"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/invariant/active_carts"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/invariant/active_carts/memory"
	"github.com/with-hindsight/chronicle/infrastructure"
	"github.com/with-hindsight/chronicle/domain/context_map"
	"reflect"
)

var CartAggregate = &context_map.Aggregate{
	Id: infrastructure.IdentifierString("f5e5d99a"),
	Name: "cart",
	Context: &context_map.Context{
		Id: infrastructure.IdentifierString("958bbea6"),
		Name: "shopping",
		Domain: &context_map.Domain{
			Id: infrastructure.IdentifierString("b9a185c3"),
			Name: "e-commerce",
		},
		Projectors: []*context_map.Projector{
			&context_map.Projector{
				Id: infrastructure.IdentifierString("1631ad2a"),
				Handler: reflect.TypeOf(active_carts.Projector{}),
				Invariant: reflect.TypeOf(active_carts.Invariant{}),
				Projection: memory.NewProjection(memory.SampleStore),
				Query: memory.NewProjection(memory.SampleStore),
			},
		},
	},
	Archetype: &context_map.Archetype{
		Invariant: reflect.TypeOf(cart.Invariant{}),
		ProjectorHandler: reflect.TypeOf(cart.Projector{}),
		State: reflect.TypeOf(cart.State{}),
	},
}

var ContextMap = context_map.New(
	[]*context_map.Object{
		&context_map.Object{
			Id: infrastructure.IdentifierString("5bc8ca6b"),
			Name: "create",
			Version: 1,
			Type: reflect.TypeOf(command.Create{}),
			Aggregate: CartAggregate,
		},
		&context_map.Object{
			Id: infrastructure.IdentifierString("08d14870"),
			Name: "add-product",
			Version: 1,
			Type: reflect.TypeOf(command.AddProduct{}),
			Aggregate: CartAggregate,
		},
		&context_map.Object{
			Id: infrastructure.IdentifierString("1e420e4e"),
			Name: "change-product-quantity",
			Version: 1,
			Type: reflect.TypeOf(command.ChangeProductQuantity{}),
			Aggregate: CartAggregate,
		},
		&context_map.Object{
			Id: infrastructure.IdentifierString("c308783c"),
			Name: "remove-product",
			Version: 1,
			Type: reflect.TypeOf(command.RemoveProduct{}),
			Aggregate: CartAggregate,
		},
		&context_map.Object{
			Id: infrastructure.IdentifierString("6bec2657"),
			Name: "checkout",
			Version: 1,
			Type: reflect.TypeOf(command.Checkout{}),
			Aggregate: CartAggregate,
		},
		&context_map.Object{
			Id: infrastructure.IdentifierString("1edc201b"),
			Name: "created",
			Version: 1,
			Type: reflect.TypeOf(event.Created{}),
			Aggregate: CartAggregate,
		},
		&context_map.Object{
			Id: infrastructure.IdentifierString("f1b88fd9"),
			Name: "empty",
			Version: 1,
			Type: reflect.TypeOf(event.Empty{}),
			Aggregate: CartAggregate,
		},
		&context_map.Object{
			Id: infrastructure.IdentifierString("c952cb41"),
			Name: "full",
			Version: 1,
			Type: reflect.TypeOf(event.Full{}),
			Aggregate: CartAggregate,
		},
		&context_map.Object{
			Id: infrastructure.IdentifierString("72784981"),
			Name: "product-added",
			Version: 1,
			Type: reflect.TypeOf(event.ProductAdded{}),
			Aggregate: CartAggregate,
		},
		&context_map.Object{
			Id: infrastructure.IdentifierString("89a0688f"),
			Name: "product-quantity-changed",
			Version: 1,
			Type: reflect.TypeOf(event.ProductQuantityChanged{}),
			Aggregate: CartAggregate,
		},
		&context_map.Object{
			Id: infrastructure.IdentifierString("17d8c65e"),
			Name: "product-removed",
			Version: 1,
			Type: reflect.TypeOf(event.ProductRemoved{}),
			Aggregate: CartAggregate,
		},
		&context_map.Object{
			Id: infrastructure.IdentifierString("2481cb7f"),
			Name: "checked-out",
			Version: 1,
			Type: reflect.TypeOf(event.CheckedOut{}),
			Aggregate: CartAggregate,
		},
	},
)
