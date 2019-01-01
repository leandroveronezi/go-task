package example

import (
	"time"
)

// TaskO - Function example
// GROUP:DEV,PROD
func TaskO() {

	time.Sleep(1 * time.Second)

}

// TaskP - Function example
// GROUP:DEV
func TaskP() {

	time.Sleep(1 * time.Second)

}

// TaskQ - Function example
// GROUP:PROD
func TaskQ() {

	time.Sleep(1 * time.Second)

}
