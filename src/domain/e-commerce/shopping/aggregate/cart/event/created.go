package event

import (
	"github.com/with-hindsight/chronicle/src/domain"
)

type Created struct {

	ShopperId domain.Identifier
}
