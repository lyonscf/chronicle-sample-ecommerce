package cart

import (
	"errors"
	"github.com/with-hindsight/chronicle/src/domain"
	"github.com/with-hindsight/chronicle/src/domain/context/aggregate"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/invariant/active_carts"
	"github.com/with-hindsight/chronicle/domain/rule"
)

var (
	ErrIsCreated = errors.New("The cart is created.")
	ErrIsNotCreated = errors.New("The cart is not created.")

	ErrIsCheckedOut = errors.New("The cart is checked out.")
	ErrIsNotCheckedOut = errors.New("This cart is not checked out.")

	ErrIsEmpty = errors.New("This cart is empty.")
	ErrIsNotEmpty = errors.New("This cart is not empty.")

	ErrIsFull = errors.New("This cart is full.")
	ErrIsNotFull = errors.New("This cart is not full.")

	ErrProductExists = errors.New("The product exists in this cart.")
	ErrProductNotExists = errors.New("The product does not exist in this cart.")
)

type Invariant struct {

	Checker aggregate.Checker
	Snapshot *aggregate.Snapshot
}

func (self *Invariant) Created() rule.Rule {

	state := self.Snapshot.State.(*State)

	return rule.NewRule(
		state.IsCreated == true,
		ErrIsCreated,
		ErrIsNotCreated,
	)
}

func (self *Invariant) CheckedOut() rule.Rule {

	state := self.Snapshot.State.(*State)

	return rule.NewRule(
		state.IsCheckedOut == true,
		ErrIsCheckedOut,
		ErrIsNotCheckedOut,
	)
}

func (self *Invariant) Empty() rule.Rule {

	state := self.Snapshot.State.(*State)

	return rule.NewRule(
		state.Products.Empty(),
		ErrIsEmpty,
		ErrIsNotEmpty,
	)
}

func (self *Invariant) Full() rule.Rule {

	state := self.Snapshot.State.(*State)

	return rule.NewRule(
		state.Products.Size() >= 10000,
		ErrIsFull,
		ErrIsNotFull,
	)
}

func (self *Invariant) ExistingProduct(id domain.Identifier) rule.Rule {

	state := self.Snapshot.State.(*State)

	return rule.NewRule(
		state.Products.Contains(id),
		ErrProductExists,
		ErrProductNotExists,
	)
}

func (self *Invariant) ActiveCart() rule.Rule {

	return self.Checker.Check(
		active_carts.Invariant{
			Cart: self.Snapshot.Id,
		},
	)
}
