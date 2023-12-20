package log

type logMsg = string

// General log message
const (
	InfoConnectionCreated logMsg = "Database connection established!"
	DebugConnectionRetry  logMsg = "Retry in %d seconds..."
)

// NATS-related log message
const (
	ErrNatsConnect  logMsg = "Error connecting to NATS - %s"
	ErrNatsMicroAdd logMsg = "Error adding NATS service - %s"
)

// MySQL-ralated log message
const (
	ErrMySqlConnect      logMsg = "Error connecting to MySQL - %s"
	ErrMySqlUnknwonTable logMsg = "Error: unknwon table %s"
)
