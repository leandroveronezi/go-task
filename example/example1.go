package example

import (
	"time"
)

// TaskC - Function example
func TaskC() bool {

	time.Sleep(1 * time.Second)

	return true

}

// TaskB - Function example
func TaskB() error {

	time.Sleep(1 * time.Second)

	return nil

}

// TaskA - Function example
func TaskA() {

	time.Sleep(1 * time.Second)

}
