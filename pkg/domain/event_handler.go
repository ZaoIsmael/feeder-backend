package domain

type EventHandler interface {
	Execute(e EventDomain)
	EventSubscriberName() string
}
