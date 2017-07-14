package comm

import (
	"fmt"
)

//error type to string，方便重定向输出
func CustomerErr(err error) string {
	if err != nil {
		return fmt.Sprintf("%s", err.Error())
	}
	return ""
}
