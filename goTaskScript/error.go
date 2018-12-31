package goTaskScript

import (
	"errors"
	"path/filepath"
	"runtime"
	"strings"
)

/*
Task Return representation
*/
type GoTaskReturn struct {
	Err  error
	Line int
	File string
}

/*
Process error return
*/
func ProcessReturn(Err error) *GoTaskReturn {

	filename, fileline := fileLine(2)

	result := GoTaskReturn{}

	result.Err = Err
	result.Line = fileline - 1
	result.File = filename

	return &result

}

/*
Process boolean return
*/
func ProcessBool(Result bool) *GoTaskReturn {

	filename, fileline := fileLine(2)

	result := GoTaskReturn{}

	result.Err = nil
	if !Result {
		result.Err = errors.New("false returned")
	}

	result.Line = fileline - 1
	result.File = filename

	return &result

}

const packageName = "github.com/leandroveronezi/go-task/goTaskScript"

func fileLine(depth int) (string, int) {

	for i := depth; ; i++ {

		_, file, line, ok := runtime.Caller(i)

		if !ok {
			break
		}

		if strings.Contains(file, packageName) {
			continue
		}

		aux, _ := filepath.Abs(file)
		return aux, line
	}

	return "", 0

}
