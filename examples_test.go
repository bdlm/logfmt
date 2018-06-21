package logfmt_test

import (
	"os"

	logfmt "github.com/bdlm/logfmt"
	log "github.com/sirupsen/logrus"
)

func ExampleFormat_text() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logfmt.TextFormat{})

	log.Infof("This is a log message")

	log.WithFields(log.Fields{
		"field1": "value 1",
		"field2": "value 2",
	}).Infof("This is a log message with additional fields")
}

func ExampleFormat_json() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logfmt.JSONFormat{})

	log.Infof("This is a log message")

	log.WithFields(log.Fields{
		"field1": "value 1",
		"field2": "value 2",
	}).Infof("This is a log message with additional fields")
}
