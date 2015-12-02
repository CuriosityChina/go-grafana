package grafana

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"regexp"
)

var (
	globalLogger *log.Logger = log.New(ioutil.Discard, "[debug] ", log.LstdFlags|log.Lshortfile)
	globalDebug  bool
)

func SetLogger(logOutput io.Writer) {
	globalLogger = log.New(logOutput, "[debug] ", log.LstdFlags|log.Lshortfile)
}

func SetDebug(debug bool) {
	globalDebug = debug
}

func debugln(args ...interface{}) {
	if globalDebug && globalLogger != nil {
		globalLogger.Println(args...)
	}
}

func debugf(message string, args ...interface{}) {
	if globalDebug && globalLogger != nil {
		globalLogger.Printf(message+"\n", args...)
	}
}

var oneLogLineRegex = regexp.MustCompile(`(?m)^\s*`)

// oneLogLine removes indentation at the beginning of each line and
// escapes new line characters.
func oneLogLine(in []byte) []byte {
	return bytes.Replace(oneLogLineRegex.ReplaceAll(in, nil), []byte("\n"), []byte("\\n "), -1)
}
