package event

import (
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/value"
	"github.com/with-hindsight/chronicle/src/domain"
)

type ProductQuantityChanged struct {

	ProductId domain.Identifier
	Quantity value.Quantity
}
