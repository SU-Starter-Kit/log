package messages

import "encoding/json"

// LogMessage
// provides basic structure for company log message.
type LogMessage struct {
	CorrelationId string            `json:"correlation_id"`
	Message       string            `json:"message,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
	Err           string            `json:"error,omitempty"`
}

func (lm *LogMessage) AsJson() string {
	b, _ := json.Marshal(lm)
	return string(b)
}

type logMessageOption func(lm *LogMessage)

// New
//
// # Returns a new LogMessageBuilder
//
// @param logMessage can be changed later using `WithMessage` method.
func New(logMessage string, opts ...logMessageOption) *LogMessage {
	// Sets default log message
	msg := &LogMessage{
		Message: logMessage,
		Tags:    make(map[string]string),
	}

	for _, opt := range opts {
		opt(msg)
	}

	return msg
}

func WithCorrelationId(correlationId string) logMessageOption {
	return func(msg *LogMessage) {
		msg.CorrelationId = correlationId
	}
}

func WithTag(key, value string) logMessageOption {
	return func(msg *LogMessage) {
		msg.Tags[key] = value
	}
}

func WithError(err error) logMessageOption {
	return func(msg *LogMessage) {
		msg.Err = err.Error()
	}
}
