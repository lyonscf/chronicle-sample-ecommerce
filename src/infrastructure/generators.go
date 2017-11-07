package infrastructure

import (
	"github.com/with-hindsight/chronicle/infrastructure/generator"
	"github.com/with-hindsight/chronicle-identifier-uuid"
)

var IdentifierGenerator = uuid.NewGenerator()

var CommandGenerator = generator.NewCommandGenerator(ContextMap)
var EventGenerator = generator.NewEventGenerator(IdentifierGenerator)

var AggregateGenerator = generator.NewAggregateGenerator(
	EventGenerator,
	&StubChecker{},
)

var ProjectionGenerator = generator.NewProjectorGenerator(
	EventStream,
)

func Boot() {

	EventGenerator.SetContextMap(ContextMap)
}
