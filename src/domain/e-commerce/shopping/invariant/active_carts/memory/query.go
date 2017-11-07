package memory

import "github.com/with-hindsight/chronicle/domain"

type Query struct {

	store *Store
}

func (self *Query) Exists(shopper domain.Identifier) (bool, error) {

	return self.store.shoppers.Contains(shopper), nil
}

func NewQuery(store *Store) *Query {

	return &Query{
		store: store,
	}
}
