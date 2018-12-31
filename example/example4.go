package example

import (
	"time"
)

func TaskF() {

	taskG()

}

func taskG() {

	time.Sleep(1 * time.Second)

}
