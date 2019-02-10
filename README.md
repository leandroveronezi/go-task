## go-task

[![Go Report Card](https://goreportcard.com/badge/github.com/leandroveronezi/go-task)](https://goreportcard.com/report/github.com/leandroveronezi/go-task)
[![GoDoc](https://godoc.org/github.com/leandroveronezi/go-task?status.png)](https://godoc.org/github.com/leandroveronezi/go-task)
![MIT Licensed](https://img.shields.io/github/license/leandroveronezi/go-task.svg)
![](https://img.shields.io/github/repo-size/leandroveronezi/go-task.svg)
[![](https://img.shields.io/badge/Require-go--terminal-blue.svg)](https://github.com/leandroveronezi/go-terminal)
[![](https://img.shields.io/badge/Require-go-1-11-blue.svg)](http://golang.com)


Go-Task is a very simple library that allows you to write simple "task" scripts in Go and run.

## First, we need $GOPATH/bin

```bash
export PATH=$PATH:$GOPATH/bin
```

## Install 

```bash
go install github.com/leandroveronezi/go-task
```

## Usage

```bash
go-task -f file.go
```

###### Parameters

```
    -f      File
    -silent Silent mode
    -k      Keep generated file
    -w      View generated source 
    -s      Sort orders of functions by name before run
    -c      Skip errors and continue
    -t      Target functions
    -g      Run function by group
```

###### Simple File

> An import between parenthesis is required.

> Run only Exported functions.

> No main function allowed

```go
package example

import (
	"time"
)

func TaskA() {
	time.Sleep(1 * time.Second)
}
```

###### Grouped

```go
package example

import (
	"time"
)

// group:dev,prod
func TaskA() {
	time.Sleep(1 * time.Second)
}

// group:dev
func TaskB() {
	time.Sleep(1 * time.Second)
}

// group:prod
func TaskC() {
	time.Sleep(1 * time.Second)
}
```

```bash
go-task -f file.go -g dev
```

## Examples

###### File example5.go

```go
package example

import (
	"github.com/pkg/errors"
	"time"
)

func TaskN() error {

	time.Sleep(1 * time.Second)

	return errors.New("An error")

}

func TaskM() bool {

	time.Sleep(1 * time.Second)

	return false

}

func TaskL() bool {

	time.Sleep(1 * time.Second)

	return true

}

func TaskK() error {

	time.Sleep(1 * time.Second)

	return nil

}

func TaskJ() {

	time.Sleep(1 * time.Second)

}

func TaskH() {

	taskI()

}

func taskI() {

	time.Sleep(1 * time.Second)

}
```

###### Run

```bash
go-task -f example5.go -c -s
```

###### Terminal

![](https://leandroveronezi.github.io/go-task/example/img/example5.png)


###### Run Functions by name

```bash
go-task -f example5.go -c -s -t TaskH,TaskL
```

![](https://leandroveronezi.github.io/go-task/example/img/example5_2.png)
