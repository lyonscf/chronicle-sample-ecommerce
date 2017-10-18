package command

import (
	"github.com/with-hindsight/chronicle/domain/aggregate"
	"github.com/with-hindsight/chronicle/domain/rule"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
)

type Checkout struct {}

func (self Checkout) Check(inv aggregate.Invariant) (bool, []error) {

	invariant := inv.(*cart.Invariant)

	return rule.Assert(
		rule.Not(invariant.CheckedOut()),
	)
}

func (self Checkout) Handle(projector aggregate.Projector, inv aggregate.Invariant) {

	projector.Apply(
		event.CheckedOut{},
	)
}
