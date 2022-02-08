package logging

const (
	// environment variables
	encodingEnvVar  = "KONF_LOG_ENCODING"
	encodingDefault = consoleEncoding
	levelEnvVar     = "KONF_LOG_LEVEL"
	levelDefault    = infoLevel

	// encodings
	consoleEncoding = "console"
	jsonEncoding    = "json"

	// levels
	errorLevel = "error"
	warnLevel  = "warn"
	infoLevel  = "info"
	debugLevel = "debug"
)

var (
	availableEncodings = []string{consoleEncoding, jsonEncoding}
	availableLevels    = []string{debugLevel, infoLevel, warnLevel, errorLevel}
)
