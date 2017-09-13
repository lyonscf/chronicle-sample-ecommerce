package event

import "github.com/with-hindsight/chronicle/src/domain/core"
import "github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/value"

type ProductQuantityChanged struct {

	ProductId core.Identifier
	Quantity value.Quantity
}
