package event

import "github.com/lyonscf/chronicle-sample-ecommerce/src/domain/e-commerce/shopping/entity"

type ProductAdded struct {

	Product entity.Product `json:"product"`
}
