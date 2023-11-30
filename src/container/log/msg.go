package log

type logMsg = string

// NATS-related log message
const (
	InfoNatsMicroCreated logMsg = "NATS micro-service added successfully!"
	ErrNatsConnect       logMsg = "Error connecting to NATS - %s"
	ErrNatsMicroAdd      logMsg = "Error adding NATS service - %s"
)

// MySQL-ralated log message
const (
	ErrMySqlConnect      logMsg = "Error connecting to MySQL - %s"
	ErrMySqlUnknwonTable logMsg = "Error: unknwon table %s"
)
