package example

import (
	"errors"
	"time"
)

func E() error {

	time.Sleep(1 * time.Second)

	return errors.New("An error")

}
