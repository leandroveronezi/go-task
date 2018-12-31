package example

import (
	"time"
)

func TaskC() bool {

	time.Sleep(1 * time.Second)

	return true

}

func TaskB() error {

	time.Sleep(1 * time.Second)

	return nil

}

func TaskA() {

	time.Sleep(1 * time.Second)

}
