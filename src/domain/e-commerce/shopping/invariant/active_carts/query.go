package active_carts

import "github.com/with-hindsight/chronicle/domain"

type Query interface {

	Exists(shopper domain.Identifier) (bool, error)
}
