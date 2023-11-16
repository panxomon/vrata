package domain

//interface with the domain event on user entity
type Event interface {
	//return the name of the event
	Name() string
	//return the data of the event
	Data() interface{}
}

type EventHandler interface {
	//handle the event
	Handle(event Event) error
}

type event struct {
	name string
	data interface{}
}

type eventHandler struct {
	handlers map[string]EventHandler
}

type eventHandlerFunc func(event Event) error

func (h eventHandlerFunc) Handle(event Event) error {
	return h(event)
}

func NewEventHandler() *eventHandler {
	return &eventHandler{
		handlers: make(map[string]EventHandler),
	}
}
