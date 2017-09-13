package cart

import (
	"github.com/with-hindsight/chronicle/src/domain/core"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/sets"
)

type State struct {

	IsCreated bool
	IsCheckedOut bool

	ShopperId core.Identifier
	Products sets.Set
}

func (self *State) Reset() {

	self.IsCreated = false
	self.IsCheckedOut = false

	self.ShopperId = nil
	self.Products = hashset.New()
}
