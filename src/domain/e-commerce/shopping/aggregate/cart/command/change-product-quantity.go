package command

import (
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/value"
	"github.com/with-hindsight/chronicle/src/domain/core"
)

type ChangeProductQuantity struct {

	ProductId core.Identifier
	Quantity value.Quantity
}
