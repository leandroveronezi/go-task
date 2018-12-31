package example

import (
	"time"
)

func C() bool {

	time.Sleep(1 * time.Second)

	return true

}

func B() error {

	time.Sleep(1 * time.Second)

	return nil

}

func A() {

	time.Sleep(1 * time.Second)

}
