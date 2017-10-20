package active_carts

import "github.com/with-hindsight/chronicle/domain"

type Query interface {

	Exists(cart domain.Identifier, shopper domain.Identifier) bool
}
