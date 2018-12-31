package example

import (
	"errors"
	"time"
)

func TaskE() error {

	time.Sleep(1 * time.Second)

	return errors.New("An error")

}
