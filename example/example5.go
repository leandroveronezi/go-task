package example

import (
	"github.com/pkg/errors"
	"time"
)

/*
TaskN - Function example
*/
func TaskN() error {

	time.Sleep(1 * time.Second)

	return errors.New("An error")

}

/*
TaskM - Function example
*/
func TaskM() bool {

	time.Sleep(1 * time.Second)

	return false

}

/*
TaskL - Function example
*/
func TaskL() bool {

	time.Sleep(1 * time.Second)

	return true

}

/*
TaskK - Function example
*/
func TaskK() error {

	time.Sleep(1 * time.Second)

	return nil

}

/*
TaskJ - Function example
*/
func TaskJ() {

	time.Sleep(1 * time.Second)

}

/*
TaskH - Function example
*/
func TaskH() {

	taskI()

}

/*
TaskI - Function example
*/
func taskI() {

	time.Sleep(1 * time.Second)

}
