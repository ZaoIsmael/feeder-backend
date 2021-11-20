package domain

type AggregateRoot struct {
	domainEvents []EventDomain
}

func (a AggregateRoot) RegisterEvent(e EventDomain) {
	a.domainEvents = append(a.domainEvents, e)
}

func (a AggregateRoot) Pull() []EventDomain {
	events := a.domainEvents

	a.domainEvents = nil
	return events
}
