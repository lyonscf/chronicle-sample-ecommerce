package infrastructure

import "github.com/with-hindsight/chronicle/infrastructure/repository"

var PlayerRepository = repository.NewPlayerRepository()

var AggregateRepository = repository.NewAggregateRepository(
	AggregateStream,
	EventLog,
)
