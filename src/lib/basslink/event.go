package basslink

type Event struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorEvent struct {
	Error error
	Event
}

type InfoEvent struct {
	Event
}

type DebugEvent struct {
	Event
}

func NewEvent(message string, data interface{}) Event {
	return Event{message, data}
}

func NewErrorEvent(err error, message string, data interface{}) *ErrorEvent {
	return &ErrorEvent{
		Error: err,
		Event: NewEvent(message, data),
	}
}

func NewInfoEvent(message string, data interface{}) *InfoEvent {
	return &InfoEvent{
		Event: NewEvent(message, data),
	}
}

func NewDebugEvent(message string, data interface{}) *DebugEvent {
	return &DebugEvent{
		Event: NewEvent(message, data),
	}
}
