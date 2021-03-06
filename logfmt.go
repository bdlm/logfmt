/*
Package logfmt defines a custom log format in the form of:
[ISO-8601-date] [level] [hostname] [caller] [message] [additional fields]
	log.SetFormatter(&logfmt.textFormat{})
	log.SetFormatter(&logfmt.jsonFormat{})
*/
package logfmt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
	"text/template"

	"github.com/bdlm/log"
)

/*
RFC3339Milli defines an RFC3339 date format with miliseconds
*/
const RFC3339Milli = "2006-01-02T15:04:05.000Z07:00"

/*
JSONFormat extends log.JSONFormatter.
*/
type JSONFormat struct {
	*log.JSONFormatter
}

/*
TextFormat extends log.TextFormatter.
*/
type TextFormat struct {
	*log.TextFormatter
}

type logData struct {
	Timestamp string            `json:"time"`
	Level     string            `json:"level"`
	Hostname  string            `json:"host"`
	Message   string            `json:"msg"`
	Data      map[string]string `json:"data"`
	Caller    string            `json:"caller"`
}

/*
Format is a custom log format method
*/
func (l *JSONFormat) Format(entry *log.Entry) ([]byte, error) {
	data := getData(entry)
	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal log data as JSON: %s", err.Error())
	}
	return append(serialized, '\n'), nil
}

/*
Format is a custom log format method
*/
func (l *TextFormat) Format(entry *log.Entry) ([]byte, error) {
	var logLine *bytes.Buffer

	if entry.Buffer != nil {
		logLine = entry.Buffer
	} else {
		logLine = &bytes.Buffer{}
	}

	data := getData(entry)
	textTemplate.Execute(logLine, data)
	logLine.WriteByte('\n')
	return logLine.Bytes(), nil
}

var textTemplate = template.Must(
	template.New("log").Parse(`time="{{.Timestamp}}" host="{{.Hostname}}" level="{{.Level}}" msg="{{.Message}}" {{range $k, $v := .Data}}{{$k}}="{{$v}}" {{end}}caller="{{.Caller}}"`),
)

func getCaller() string {
	caller := ""
	a := 0
	for {
		if pc, file, line, ok := runtime.Caller(a + 2); ok {
			if !strings.Contains(strings.ToLower(file), "github.com/sirupsen/logrus") && !strings.Contains(strings.ToLower(file), "github.com/bdlm/logfmt") {
				caller = fmt.Sprintf("%s:%d %s", path.Base(file), line, runtime.FuncForPC(pc).Name())
				break
			}
		} else {
			break
		}
		a++
	}
	return caller
}

/*
getData is a helper function that extracts log data from the logrus
entry.
*/
func getData(entry *log.Entry) *logData {
	data := &logData{
		Caller:    getCaller(),
		Data:      make(map[string]string),
		Hostname:  os.Getenv("HOSTNAME"),
		Level:     entry.Level.String(),
		Message:   entry.Message,
		Timestamp: entry.Time.Format(RFC3339Milli),
	}

	keys := make([]string, 0)
	for k := range entry.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for k, v := range entry.Data {
		data.Data[k] = fmt.Sprintf("%v", v)
	}

	return data
}
