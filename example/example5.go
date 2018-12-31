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
