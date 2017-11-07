package active_carts

import (
	"github.com/with-hindsight/chronicle/domain"
	"github.com/with-hindsight/chronicle/domain/projection"
)

type Invariant struct {

	Shopper domain.Identifier
}

func (self Invariant) Check(qry projection.Query) (bool, error) {

	query := qry.(Query)

	return query.Exists(self.Shopper)
}
