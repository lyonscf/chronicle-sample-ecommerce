package command

import (
	"github.com/with-hindsight/chronicle/domain/aggregate"
	"github.com/with-hindsight/chronicle/domain/rule"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/entity"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
)

type AddProduct struct {

	Product entity.Product `json:"product"`
}

func (self AddProduct) Check(inv aggregate.Invariant) (bool, []error) {

	invariant := inv.(*cart.Invariant)

	return rule.Assert(
		rule.Is(invariant.Created()),
		rule.Not(invariant.ExistingProduct(self.Product.Id)),
		rule.Not(invariant.Full()),
	)
}

func (self AddProduct) Handle(projector aggregate.Projector, inv aggregate.Invariant) {

	invariant := inv.(*cart.Invariant)

	projector.Apply(
		event.ProductAdded{
			self.Product,
		},
	)

	if rule.Check(invariant.Full()) {

		projector.Apply(
			event.Full{},
		)
	}
}
