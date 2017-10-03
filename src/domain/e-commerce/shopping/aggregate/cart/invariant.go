package cart

import (
	"errors"
	"github.com/with-hindsight/chronicle/src/domain"
)

var (
	ErrIsCreated = errors.New("The cart is already created.")
	ErrIsNotCreated = errors.New("")
	ErrIsCheckedOut = errors.New("")
	ErrIsNotCheckedOut = errors.New("")
	ErrIsEmpty = errors.New("")
	ErrIsFull = errors.New("")
	ErrProductExists = errors.New("")
	ErrProductNotExists = errors.New("")
)

type Invariant struct {

	state *State
}

func (self *Invariant) IsCreated() bool {

	return self.state.IsCreated == true
}

func (self *Invariant) IsEmpty() bool {

	return self.state.Products.Empty()
}

func (self *Invariant) IsFull() bool {

	return self.state.Products.Size() >= 10
}

func (self *Invariant) ProductExists(id domain.Identifier) bool {

	return self.state.Products.Contains(id)
}
