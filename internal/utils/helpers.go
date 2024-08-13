package utils

import (
	"errors"
	"fmt"
)

// helper functions
func HandleDBError(m string, err error) error {
	return errors.New(fmt.Sprint(m, err))
}
