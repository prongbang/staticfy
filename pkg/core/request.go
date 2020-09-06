package core

import "fmt"

func AllowedMethod(method string, reqMethod string) error {
	switch reqMethod {
	case method:
		return nil
	default:
		return fmt.Errorf("%s", "Method Not Allowed")
	}
}
