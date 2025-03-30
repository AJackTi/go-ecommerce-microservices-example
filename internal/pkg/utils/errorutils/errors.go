package errorutils

import "fmt"

// ErrorsWithStack returns a string contains errors messages in the stack with its stack trace levels for given error
func ErrorsWithStack(err error) string {
	res := fmt.Sprintf("%+v\n", err)
	return res
}
