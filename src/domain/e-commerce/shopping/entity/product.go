package entity

import (
	"github.com/with-hindsight/chronicle/domain"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/value"
)

type Product struct {

	Id domain.Identifier `json:"id"`
	Quantity value.Quantity `json:"quantity"`
}

func (self Product) Check() error {

	return self.Quantity.Check()
}
