package event

import (
	"github.com/with-hindsight/chronicle/src/domain"
)

type ProductRemoved struct {

	ProductId domain.Identifier
}
