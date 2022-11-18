package logger

import (
	"io"
	"log"
	"os"

	"github.com/su-starter-kit/log/internal"
)

type loggerBuilder struct {
	loggerConfig *internal.LoggerConfiguration
}

// WithLogFlags
//
// flags int - golangs log package flags
//
// Eg.: log.Ldate | log.ShortFile
//
// Default: log.Ldate | log.ShortFile
func (lb *loggerBuilder) WithLogFlags(flags int) *loggerBuilder {
	lb.loggerConfig.LogFlags = flags
	return lb
}

// WithOutput
//
// out io.Writer - golangs io.Writer
//
// Eg.: file | os.Stdout | os.Stderr
//
// Default: os.Stdout
func (lb *loggerBuilder) WithOutput(out io.Writer) *loggerBuilder {
	lb.loggerConfig.Output = out
	return lb
}

// WithCorrelationId
//
// correlationId string
//
// Sets default correlation id for the logs logged by the logger constructed by this builder.
//
// NOTE: If CorrelationId is specified in the log message, it will have priority over CorrelationId
// configured by this method.
func (lb *loggerBuilder) WithCorrelationId(correlationId string) *loggerBuilder {
	lb.loggerConfig.CorrelationId = correlationId
	return lb
}

func (lb *loggerBuilder) Build() CompanyLogger {
	return internal.NewLogger(lb.loggerConfig)
}

// NewBuilder
//
// Defaults:
//
// Output: os.Stdout
//
// LogFlags: os.Stdout | os.Stderr
func NewBuilder() *loggerBuilder {
	return &loggerBuilder{
		loggerConfig: &internal.LoggerConfiguration{
			Output:   os.Stdout,
			LogFlags: log.Ldate | log.Lshortfile,
		},
	}
}
