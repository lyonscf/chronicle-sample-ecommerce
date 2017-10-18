package infrastructure

import (
	"github.com/with-hindsight/chronicle-storage-inmemory"
)

var EventStore = memory.NewStore()
var EventLog = memory.NewLog(EventStore)

var AggregateStream = memory.NewAggregateStream(
	EventGenerator,
	EventStore,
)

