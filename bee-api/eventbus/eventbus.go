package eventbus

import (
	evbus "github.com/asaskevich/EventBus"
)

var EventBus = &eventBus{
	bus: evbus.New(),
}

type eventBus struct {
	bus evbus.Bus
}

func (eb *eventBus) Subscribe(event string, handler interface{}) error {
	return eb.bus.Subscribe(event, handler)
}

func (eb *eventBus) SubscribeAsync(event string, handler interface{}, transactional bool) error {
	return eb.bus.SubscribeAsync(event, handler, transactional)
}
