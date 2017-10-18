package active_carts

import "github.com/with-hindsight/chronicle/domain"

type Invariant struct {

	Cart domain.Identifier
}

func (self Invariant) Check(query Query) (bool, error) {

	return query.Exists(self.Cart), nil
}
