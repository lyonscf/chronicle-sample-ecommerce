package command

import (
	"github.com/with-hindsight/chronicle/domain"
	"github.com/with-hindsight/chronicle/domain/aggregate"
	"github.com/with-hindsight/chronicle/domain/rule"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
)

type RemoveProduct struct {

	ProductId domain.Identifier
}

func (self RemoveProduct) Check(inv aggregate.Invariant) (bool, []error) {

	invariant := inv.(*cart.Invariant)

	return rule.Assert(
		rule.Not(invariant.CheckedOut()),
		rule.Is(invariant.ExistingProduct(self.ProductId)),
	)
}

func (self RemoveProduct) Handle(projector aggregate.Projector, inv aggregate.Invariant) {

	invariant := inv.(*cart.Invariant)

	projector.Apply(
		event.ProductRemoved{
			self.ProductId,
		},
	)

	if rule.Check(invariant.Empty()) {

		projector.Apply(
			event.Empty{},
		)
	}
}
