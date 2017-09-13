package has_active_cart

import (
	"github.com/with-hindsight/chronicle/src/domain/core"
)

type Projection interface {

	Create(cart core.Identifier)

	Checkout(cart core.Identifier)
}
