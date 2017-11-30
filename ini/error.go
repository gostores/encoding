package ini

import (
	"fmt"
)

type ErrDelimiterNotFound struct {
	Line string
}

func IsErrDelimiterNotFound(err error) bool {
	_, ok := err.(ErrDelimiterNotFound)
	return ok
}

func (err ErrDelimiterNotFound) Error() string {
	return fmt.Sprintf("key-value delimiter not found: %s", err.Line)
}
