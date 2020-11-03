package escape

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// Escape contains methods related
// to regex escaping for the escape
// package
type Escape interface {
	EscString() (string, error)
	EscPipe() (string, error)
}

// EscString creates a reader and creates
// an escaped regex string from input
func EscString(s string) (string, error) {

	data := strings.NewReader(s)
	output, err := ioutil.ReadAll(data)
	if err != nil {
		log.Panic(err)
	}

	result := regexp.QuoteMeta(string(output))

	return result, err
}

// EscPipe creates a reader and creates
// an escaped regex string from input
func EscPipe(p *os.File) (string, error) {

	output, err := ioutil.ReadAll(p)
	if err != nil {
		log.Panic(err)
	}

	result := regexp.QuoteMeta(string(output))

	return result, err
}
