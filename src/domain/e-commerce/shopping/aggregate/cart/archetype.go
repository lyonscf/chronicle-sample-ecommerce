package cart

import (
	"github.com/with-hindsight/chronicle/src/domain/context"
	"github.com/with-hindsight/chronicle/src/domain/context/aggregate"
	"github.com/with-hindsight/chronicle/src/domain/context/aggregate/projector"
	"github.com/with-hindsight/chronicle/src/domain"
)

type Archetype struct {

	event_generator *context.EventGenerator

	type_id domain.Identifier

	projector *Projector

	aggregate *Aggregate

	state *State
	invariant *Invariant
}

func (self *Archetype) FromSnapshot(snapshot context.AggregateSnapshot) context.Aggregate {

	agg_snapshot := snapshot.(*aggregate.Snapshot)
	state := agg_snapshot.State.(*State)

	invariant := *self.invariant
	invariant.state = state

	agg := *self.aggregate
	agg.Invariant = invariant

	projector_handler := *self.projector

	aggregate := aggregate.NewAggregate(
		&agg,
		projector.NewProjector(
			self.event_generator,
			&projector_handler,
			agg_snapshot,
		),
	)

	return aggregate
}

func (self *Archetype) New(identifier domain.Identifier) context.Aggregate {

	state := *self.state

	snapshot := aggregate.NewSnapshot(
		identifier,
		self.type_id,
		&state,
	)

	return self.FromSnapshot(snapshot)
}

func NewArchetype(event_generator *context.EventGenerator, type_id domain.Identifier) context.AggregateArchetype {

	state := &State{}
	state.Reset()

	return &Archetype{

		event_generator: event_generator,

		type_id: type_id,

		projector: &Projector{},
		aggregate: &Aggregate{},

		state: state,
		invariant: &Invariant{},
	}
}
