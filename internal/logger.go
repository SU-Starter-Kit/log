package internal

import (
	"io"
	"log"

	"github.com/su-starter-kit/log/messages"
)

var defaultLogLevelPrefix LogLevelPrefix = LogLevelPrefix{
	debug: "[DEBUG]",
	info:  "[INFO]",
	warn:  "[WARN]",
	err:   "[ERROR]",
}

type Logger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger

	defaultCorrelationId string
}

type LogLevelPrefix struct {
	debug string
	info  string
	warn  string
	err   string
}

type LoggerConfiguration struct {
	Output io.Writer
	// LogFlags
	//
	// standard log package flags.
	//
	// Eg.: log.Ldate | log.ShortFile
	LogFlags int
	// CorrelationId
	//
	// Sets a standard correlation id to be used for the logger lifetime.
	//
	// If a correlation id is specified for the LogMessage, it will have pririty over the value
	// set to this variable.
	CorrelationId string
}

func NewLogger(config *LoggerConfiguration) *Logger {
	return &Logger{
		debugLogger:          log.New(config.Output, defaultLogLevelPrefix.debug, config.LogFlags),
		infoLogger:           log.New(config.Output, defaultLogLevelPrefix.info, config.LogFlags),
		warnLogger:           log.New(config.Output, defaultLogLevelPrefix.warn, config.LogFlags),
		errorLogger:          log.New(config.Output, defaultLogLevelPrefix.err, config.LogFlags),
		defaultCorrelationId: config.CorrelationId,
	}
}

func (l *Logger) Debug(v *messages.LogMessage) {
	commitLog(l.debugLogger, v, l.defaultCorrelationId)
}

func (l *Logger) Info(v *messages.LogMessage) {
	commitLog(l.infoLogger, v, l.defaultCorrelationId)
}

func (l *Logger) Warn(v *messages.LogMessage) {
	commitLog(l.warnLogger, v, l.defaultCorrelationId)
}

func (l *Logger) Error(v *messages.LogMessage) {
	commitLog(l.errorLogger, v, l.defaultCorrelationId)
}

func commitLog(logger *log.Logger, m *messages.LogMessage, defaultCorrelationId string) {
	resolveDefaultValues(defaultCorrelationId, m)
	logger.Println(m.AsJson())
}

// resolveDefaultValues.
//
// Note! This function changes values in LogMessages!
func resolveDefaultValues(defaultCoreId string, m *messages.LogMessage) {
	if len(m.CorrelationId) == 0 {
		m.CorrelationId = defaultCoreId
	}
}
