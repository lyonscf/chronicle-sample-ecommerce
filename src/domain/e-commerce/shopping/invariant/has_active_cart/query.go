package has_active_cart

import "github.com/with-hindsight/chronicle/src/domain/core"

type Query interface {

	IsActive(shopper_id core.Identifier) bool
}
