package active_carts

import "github.com/with-hindsight/chronicle/domain"

type Projection interface {

	Add(cart domain.Identifier) error

	Remove(cart domain.Identifier) error
}
