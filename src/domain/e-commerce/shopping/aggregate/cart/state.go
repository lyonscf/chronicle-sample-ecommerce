package cart

import (
	"github.com/with-hindsight/chronicle/domain"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/sets"
)

type State struct {

	IsCreated bool
	IsCheckedOut bool

	ShopperId domain.Identifier
	Products sets.Set
}

func (self *State) Reset() {

	self.IsCreated = false
	self.IsCheckedOut = false

	self.ShopperId = nil
	self.Products = hashset.New()
}
