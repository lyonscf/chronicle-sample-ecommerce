package cart

import (
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/command"
	"github.com/with-hindsight/chronicle/src/domain/aggregate"
	"github.com/with-hindsight/chronicle/src/domain"
)

type Aggregate struct {

	aggregate.Aggregate
	invariant *Invariant
}

func (self *Aggregate) Handle(c domain.Command) error {

	switch command := c.Payload().(type) {

	case command.Create:

		if self.invariant.IsCreated() {
			return ErrIsCreated
		}

		self.Apply(
			event.Created{
				ShopperId: command.ShopperId,
			},
		)

		self.Apply(
			event.Empty{},
		)

		return nil

	case command.AddProduct:

		if !self.invariant.IsCreated() {
			return ErrIsNotCreated
		}

		if self.invariant.IsFull() {
			return ErrIsFull
		}

		if self.invariant.ProductExists(command.Product.Id) {
			return ErrProductExists
		}

		self.Apply(
			event.ProductAdded{
				command.Product,
			},
		)

		if self.invariant.IsFull() {

			self.Apply(
				event.Full{},
			)

			return nil
		}

		return nil

	case command.RemoveProduct:

		if !self.invariant.IsCreated() {
			return ErrIsNotCreated
		}

		if self.invariant.IsEmpty() {
			return ErrIsEmpty
		}

		if !self.invariant.ProductExists(command.ProductId) {
			return ErrProductNotExists
		}

		self.Apply(
			event.ProductRemoved{
				command.ProductId,
			},
		)

		if self.invariant.IsEmpty() {

		}

		return nil

	case command.ChangeProductQuantity:

		if !self.invariant.IsCreated() {
			return ErrIsNotCreated
		}

		if !self.invariant.ProductExists(command.ProductId) {
			return ErrProductNotExists
		}

		self.Apply(
			event.ProductQuantityChanged{
				command.ProductId,
				command.Quantity,
			},
		)

		return nil

	case command.Checkout:

		if !self.invariant.IsCreated() {
			return ErrIsNotCreated
		}

		self.Apply(
			event.CheckedOut{},
		)

		return nil

	}

	return aggregate.ErrAggregateHandlerNotExists
}
