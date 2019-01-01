package example

import (
	"time"
)

// TaskO - Function example
// group:dev,prod
func TaskO() {

	time.Sleep(1 * time.Second)

}

// TaskP - Function example
// group:dev
func TaskP() {

	time.Sleep(1 * time.Second)

}

// TaskQ - Function example
// group:prod
func TaskQ() {

	time.Sleep(1 * time.Second)

}
