package logger

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// LogInit() initializes logger, setting log messages format and output destination. Please specify log level if needed.
func LogInit(logLvl ...string) {
	// Setting logging format to JSON
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Setting Log level if specified
	if len(logLvl) > 0 {
		switch logLvl[0] {
		case "Panic":
			log.SetLevel(log.PanicLevel)
		case "Fatal":
			log.SetLevel(log.FatalLevel)
		case "Error":
			log.SetLevel(log.ErrorLevel)
		case "Warn":
			log.SetLevel(log.WarnLevel)
		case "Info":
			log.SetLevel(log.InfoLevel)
		case "Debug":
			log.SetLevel(log.DebugLevel)
		case "Trace":
			log.SetLevel(log.TraceLevel)
		}
	}

}

// LogErr() func delivers error to logs
func LogErr(errMsg error) {
	log.Error(errMsg)
}

// LogFatal() func delivers fatal error to logs
func LogFatal(errMsg error) {
	log.Fatal(errMsg)
}

// LogMsg() func delivers message and its type to logs.
// Accepteble log types: Info (default), Warn, Panic, Trace, Debug.
func LogMsg(msg string, logType string) (err error) {
	switch logType {
	case "Info":
		log.Info(msg)
	case "Warn":
		log.Warn(msg)
	case "Panic":
		log.Panic(msg)
	case "Trace":
		log.Trace(msg)
	case "Debug":
		log.Debug(msg)
	case "":
		log.Info(msg)
		// returning error if type didn't specified
		err = fmt.Errorf("message type didn'n specifed. Using default 'Info' type")
	default:
		log.Info(msg)
		// returning error if type couldn't be recognised
		err = fmt.Errorf("unknown message type %q. Using default 'Info' type", logType)
	}

	return

}

// LogMsgWithFields() func allows to log your message with additional fields in JSON format.
// In order to add fields to message, provide a map value with type [string]string
// (in format [JSON-key]JSON-value)
func LogMsgWithFields(msg string, logType string, fields map[string]string) (err error) {
	var params *log.Entry

	if len(fields) > 0 {
		isFirstRecord := true

		for key, val := range fields {
			if isFirstRecord {
				// Creating new log.Entry if it is the first fields value
				params = log.WithField(key, val)
				isFirstRecord = false
			} else {
				// Using excisting log.Entry if it isn't the first fields value
				params = params.WithField(key, val)
			}
		}
	} else {
		// Delivering a hint to log if fields value is empty
		params = log.WithField("Hint", "Try to use LogMsg() func when not using additional fields")
	}

	switch logType {
	case "Info":
		params.Info(msg)
	case "Warn":
		params.Warn(msg)
	case "Panic":
		params.Panic(msg)
	case "Trace":
		params.Trace(msg)
	case "Debug":
		params.Debug(msg)
	case "":
		params.Info(msg)
		// returning error if type didn't specified
		err = fmt.Errorf("message type didn'n specifed. Using default 'Info' type")
	default:
		params.Info(msg)
		// returning error if type couldn't be recognised
		err = fmt.Errorf("unknown message type %q. Using default 'Info' type", logType)
	}

	return
}
