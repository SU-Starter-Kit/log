// Provides examples on how to use this test library.
package log_test

import (
	"fmt"

	"github.com/su-starter-kit/log/logger"
	"github.com/su-starter-kit/log/messages"
)

func ExampleNewBuilder_Logger_WithoutCorrlationId() {
	loggr, _ := logger.New(
		logger.WithLogFlags(0), // removes log flags to assert output
	)

	loggr.Debug(messages.New("").WithCorrelationId("2323424").WithMessage("this is a log").Message()) // defining message with `WithMessage`
	loggr.Debug(messages.New("this is a log").WithCorrelationId("2323424").Message())

	// Output:
	// [DEBUG]{"correlation_id":"2323424","message":"this is a log"}
	// [DEBUG]{"correlation_id":"2323424","message":"this is a log"}
}

func ExampleNewBuilder_Logger_WithCorrelationId() {
	loggr, _ := logger.New(
		logger.WithLogFlags(0),              // removes log flags to assert output
		logger.WithCorrelationid("2323424"), // setting correlation id so we expect the same result as Example above
	)

	loggr.Debug(messages.New("").WithMessage("this is a log").Message())

	// Output:
	// [DEBUG]{"correlation_id":"2323424","message":"this is a log"}
}

func ExampleLogMessage_With_Tags() {
	loggr, _ := logger.New(
		logger.WithLogFlags(0),              // removes log flags to assert output
		logger.WithCorrelationid("2323424"), // setting correlation id
	)

	loggr.Debug(messages.New("this is a log").WithTag("key", "val").Message())

	// Output:
	// [DEBUG]{"correlation_id":"2323424","message":"this is a log","tags":{"key":"val"}}
}

func ExampleLogMessage_WithShortAlias() {

	loggr, _ := logger.New(
		logger.WithLogFlags(0),              // removes log flags to assert output
		logger.WithCorrelationid("2323424"), // setting correlation id
	)

	loggr.Warn(&messages.M{Message: "this is a log"})

	// Output:
	// [WARN]{"correlation_id":"2323424","message":"this is a log"}
}

func ExampleLogError() {
	loggr, _ := logger.New(
		logger.WithLogFlags(0),              // removes log flags to assert output
		logger.WithCorrelationid("2323424"), // setting correlation id
	)

	loggr.Error(messages.New("this is a log").WithError(fmt.Errorf("this is an error")).Message())

	// Output:
	// [ERROR]{"correlation_id":"2323424","message":"this is a log","error":"this is an error"}
}
