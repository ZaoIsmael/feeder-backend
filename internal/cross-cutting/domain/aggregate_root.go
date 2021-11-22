package domain

type AggregateRoot struct {
	DomainEvents []EventDomain
}

func (a *AggregateRoot) RegisterEvent(e EventDomain) {
	a.DomainEvents = append(a.DomainEvents, e)
}

func (a *AggregateRoot) Pull() []EventDomain {
	events := make([]EventDomain, len(a.DomainEvents))
	copy(events, a.DomainEvents)

	a.DomainEvents = nil

	return events
}
