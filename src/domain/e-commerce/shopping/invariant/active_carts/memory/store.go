package memory

import (
	"github.com/emirpasic/gods/sets"
	"github.com/emirpasic/gods/maps"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/maps/hashmap"
)

var SampleStore = NewStore()

type Store struct {

	shoppers sets.Set
	carts maps.Map
}

func (self *Store) Reset() error {

	self.shoppers = hashset.New()
	self.carts = hashmap.New()

	return nil
}

func NewStore() *Store {

	return &Store{
		shoppers: hashset.New(),
		carts: hashmap.New(),
	}
}
