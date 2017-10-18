package infrastructure

import (
	"github.com/with-hindsight/chronicle/infrastructure"
	"github.com/with-hindsight/chronicle-identifier-uuid"
)

var IdentifierGenerator = uuid.NewGenerator()

var CommandGenerator = infrastructure.NewCommandGenerator(DomainMap)
var EventGenerator = infrastructure.NewEventGenerator(IdentifierGenerator)

func Boot() {

	EventGenerator.SetContextMap(DomainMap)
}
