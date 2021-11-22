package domain

type EventBus interface {
	Publish(de []EventDomain)
	AddHandlers(eh []EventHandler)
}
