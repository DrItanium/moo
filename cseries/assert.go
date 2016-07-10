// assertion emulation through the error type
package cseries

import (
	"fmt"
	"runtime"
)

type AssertionError struct {
	Function string
	Message  string
}

func (this AssertionError) Error() string {
	return fmt.Sprintf("%s: %s", this.Function, this.Message)
}
