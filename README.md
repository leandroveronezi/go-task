## go-task

[![Go Report Card](https://goreportcard.com/badge/github.com/leandroveronezi/go-task)](https://goreportcard.com/report/github.com/leandroveronezi/go-task)
[![GoDoc](https://godoc.org/github.com/leandroveronezi/go-task?status.png)](https://godoc.org/github.com/leandroveronezi/go-task)
![MIT Licensed](https://img.shields.io/github/license/leandroveronezi/go-task.svg)
![](https://img.shields.io/github/repo-size/leandroveronezi/go-task.svg)

Golang Task Runner

Run all exported functions in golang script

## Install

```bash
go install github.com/leandroveronezi/go-task
```

## Examples

### File example5.go

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

### Run

```bash
go-task -f example5.go -c -s
```

### Terminal

![](https://leandroveronezi.github.io/go-task/example/img/example5.png)


###### Run Functions by name

```bash
go-task -f example5.go -c -s -t TaskH,TaskL
```

![](https://leandroveronezi.github.io/go-task/example/img/example5_1.png)
