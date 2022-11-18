// This package provides standardized log
// functions and message for the company <company>
package logger

import "github.com/su-starter-kit/log/messages"

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
