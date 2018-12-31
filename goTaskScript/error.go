package goTaskScript

import (
	"errors"
	"path/filepath"
	"runtime"
	"strings"
)

type GoCallError struct {
	Err  error
	Line int
	File string
}

func ProcessReturn(Err error) *GoCallError {

	filename, fileline := FileLine(2)

	result := GoCallError{}

	result.Err = Err
	result.Line = fileline - 1
	result.File = filename

	return &result

}

func ProcessBool(Result bool) *GoCallError {

	filename, fileline := FileLine(2)

	result := GoCallError{}

	result.Err = nil
	if !Result {
		result.Err = errors.New("false returned")
	}

	result.Line = fileline - 1
	result.File = filename

	return &result

}

const packageName = "github.com/leandroveronezi/go-task/goTaskScript"

func FileLine(depth int) (string, int) {

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
