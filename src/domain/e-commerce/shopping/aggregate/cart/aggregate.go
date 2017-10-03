package cart

import (
	"github.com/with-hindsight/chronicle/src/domain/context/aggregate"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/command"
)

type Aggregate struct {

	//Checker aggregate.Checker
	Invariant Invariant
}

func (self *Aggregate) Handle(c aggregate.Command, projector aggregate.Projector) error {

	switch command := c.Payload.(type) {

	case command.Create:

		if self.Invariant.IsCreated() {
			return ErrIsCreated
		}

		/*if self.Checker.Check(
			has_active_cart.Invariant{
				command.ShopperId,
			},
		) {
			return has_active_cart.ErrHasActiveCart
		}*/

		projector.Apply(
			event.Created{
				ShopperId: command.ShopperId,
			},
		)

		projector.Apply(
			event.Empty{},
		)

		return nil

	case command.AddProduct:

		if !self.Invariant.IsCreated() {
			return ErrIsNotCreated
		}

		if self.Invariant.IsFull() {
			return ErrIsFull
		}

		if self.Invariant.ProductExists(command.Product.Id) {
			return ErrProductExists
		}

		projector.Apply(
			event.ProductAdded{
				command.Product,
			},
		)

		if self.Invariant.IsFull() {

			projector.Apply(
				event.Full{},
			)

			return nil
		}

		return nil

	case command.RemoveProduct:

		if !self.Invariant.IsCreated() {
			return ErrIsNotCreated
		}

		if self.Invariant.IsEmpty() {
			return ErrIsEmpty
		}

		if !self.Invariant.ProductExists(command.ProductId) {
			return ErrProductNotExists
		}

		projector.Apply(
			event.ProductRemoved{
				command.ProductId,
			},
		)

		if self.Invariant.IsEmpty() {

		}

		return nil

	case command.ChangeProductQuantity:

		if !self.Invariant.IsCreated() {
			return ErrIsNotCreated
		}

		if !self.Invariant.ProductExists(command.ProductId) {
			return ErrProductNotExists
		}

		projector.Apply(
			event.ProductQuantityChanged{
				command.ProductId,
				command.Quantity,
			},
		)

		return nil

	case command.Checkout:

		if !self.Invariant.IsCreated() {
			return ErrIsNotCreated
		}

		projector.Apply(
			event.CheckedOut{},
		)

		return nil

	}

	return aggregate.ErrAggregateHandlerNotExists
}
