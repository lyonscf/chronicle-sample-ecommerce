package command

import (
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/value"
	"github.com/with-hindsight/chronicle/src/domain"
)

type ChangeProductQuantity struct {

	ProductId domain.Identifier
	Quantity value.Quantity
}
