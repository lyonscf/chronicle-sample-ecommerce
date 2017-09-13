package has_active_cart

import "github.com/with-hindsight/chronicle/src/domain/core"

type Invariant struct {

	ShopperId core.Identifier
}

func (self *Invariant) Satisfier(query Query) bool {

	return query.IsActive(self.ShopperId) == true
}
