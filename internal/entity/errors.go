package entity

import "fmt"

var (
	ErrorUserNotFound  = fmt.Errorf("user not found")
	ErrorInvalidParams = fmt.Errorf("invalid parameters")
	ErrorInternal      = fmt.Errorf("internal error")
)
