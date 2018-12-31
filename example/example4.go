package example

import (
	"time"
)

// TaskF - Function example
func TaskF() {

	taskG()

}

// TaskG - Function example
func taskG() {

	time.Sleep(1 * time.Second)

}
