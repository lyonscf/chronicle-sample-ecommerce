package entity

import (
	"github.com/with-hindsight/chronicle/src/domain/core"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/value"
)

type Product struct {

	Id core.Identifier
	Quantity value.Quantity
}
