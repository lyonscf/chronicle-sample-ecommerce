package entity

import (
	"github.com/with-hindsight/chronicle/src/domain"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/value"
)

type Product struct {

	Id domain.Identifier
	Quantity value.Quantity
}
