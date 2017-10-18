package event

import (
	"github.com/with-hindsight/chronicle/domain"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/value"
)

type ProductQuantityChanged struct {

	ProductId domain.Identifier
	Quantity value.Quantity
}
