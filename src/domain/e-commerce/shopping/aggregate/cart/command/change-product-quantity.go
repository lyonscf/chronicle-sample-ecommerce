package command

import (
	"github.com/with-hindsight/chronicle/domain"
	"github.com/with-hindsight/chronicle/domain/aggregate"
	"github.com/with-hindsight/chronicle/domain/rule"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/value"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
)

type ChangeProductQuantity struct {

	ProductId domain.Identifier
	Quantity value.Quantity
}

func (self ChangeProductQuantity) Check(inv aggregate.Invariant) (bool, []error) {

	invariant := inv.(*cart.Invariant)

	return rule.Assert(
		rule.Not(invariant.CheckedOut()),
		rule.Is(invariant.ExistingProduct(self.ProductId)),
	)
}

func (self ChangeProductQuantity) Handle(projector aggregate.Projector, inv aggregate.Invariant) {

	projector.Apply(
		event.ProductQuantityChanged{
			ProductId: self.ProductId,
			Quantity: self.Quantity,
		},
	)
}
