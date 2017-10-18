package command

import (
	"github.com/with-hindsight/chronicle/domain"
	"github.com/with-hindsight/chronicle/domain/aggregate"
	"github.com/with-hindsight/chronicle/domain/rule"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
)

type Create struct {

	ShopperId domain.Identifier `json:"shopper_id"`
}

func (self Create) Check(inv aggregate.Invariant) (bool, []error) {

	invariant := inv.(*cart.Invariant)

	return rule.Assert(
		rule.Not(invariant.Created()),
		rule.Has(invariant.ActiveCart()),
	)
}

func (self Create) Handle(projector aggregate.Projector, inv aggregate.Invariant) {

	projector.Apply(
		event.Created{
			ShopperId: self.ShopperId,
		},
		event.Empty{

		},
	)
}
