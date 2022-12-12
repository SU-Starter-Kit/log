// This package provides standardized log
// functions and message for the company <company>
package logger

import (
	"io"
	"log"
	"os"

	"github.com/su-starter-kit/log/internal"
	"github.com/su-starter-kit/log/messages"
)

// CompanyLogger
//
// Provides the logger interface to standardize
// logs for company <company name>
type CompanyLogger interface {
	Debug(*messages.LogMessage)
	Info(*messages.LogMessage)
	Warn(*messages.LogMessage)
	Error(*messages.LogMessage)
}

// loggerOption
//
// Provides interface for building logger
type loggerOption func(lc *internal.LoggerConfiguration) error

// New
//
// Defaults:
//
// Output: os.Stdout
//
// LogFlags: os.Stdout | os.Stderr
func New(opts ...loggerOption) (CompanyLogger, error) {
	// Sets default configuration
	loggerConfig := &internal.LoggerConfiguration{
		Output:   os.Stdout,
		LogFlags: log.Ldate | log.Lshortfile,
	}

	for _, opt := range opts {
		if err := opt(loggerConfig); err != nil {
			return nil, err
		}
	}

	return internal.NewLogger(loggerConfig), nil
}

// WithLogFlags
//
// flags int - golangs log package flags
//
// Eg.: log.Ldate | log.ShortFile
//
// Default: log.Ldate | log.ShortFile
func WithLogFlags(flags int) loggerOption {
	return func(lc *internal.LoggerConfiguration) error {
		lc.LogFlags = flags
		return nil
	}
}

// WithOutput
//
// out io.Writer - golangs io.Writer
//
// Eg.: file | os.Stdout | os.Stderr
//
// Default: os.Stdout
func WithOutput(out io.Writer) loggerOption {
	return func(lc *internal.LoggerConfiguration) error {
		lc.Output = out
		return nil
	}
}

// WithCorrelationId
//
// correlationId string
//
// Sets default correlation id for the logs logged by the logger constructed by this builder.
//
// NOTE: If CorrelationId is specified in the log message, it will have priority over CorrelationId
// configured by this method.
func WithCorrelationid(correlationId string) loggerOption {
	return func(lc *internal.LoggerConfiguration) error {
		lc.CorrelationId = correlationId
		return nil
	}
}
