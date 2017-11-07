package infrastructure

import (
	"github.com/with-hindsight/chronicle/infrastructure/repository"
)

var AggregateRepository = repository.NewAggregateRepository(
	AggregateGenerator,
	AggregateStream,
	EventLog,
)

var SnapshotRepository = repository.NewProjectorSnapshotRepository()

var ProjectionRepository = repository.NewProjectorRepository(
	IdentifierGenerator,
	ContextMap,
	SnapshotRepository,
	ProjectionGenerator,
)
