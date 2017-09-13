package has_active_cart

import (
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/aggregate/cart/event"
	"github.com/with-hindsight/chronicle/src/domain"
)

var Handlers = []interface{}{

	func (projection Projection, event event.Created, e domain.Event) {

		projection.Create(e.AggregateId())
	},

	func (projection Projection, event event.CheckedOut, e domain.Event) {

		projection.Checkout(e.AggregateId())
	},
}
