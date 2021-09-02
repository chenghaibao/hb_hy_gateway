package http

import "fmt"

type GatewayError struct {
	code     uint64
	errorMsg string
}

func NewGatewayError(code uint64, errorMsg string) *GatewayError {
	return &GatewayError{code: code, errorMsg: errorMsg}
}

func (g *GatewayError) Error() string {
	return fmt.Sprintf("[%d] %s", g.code, g.errorMsg)
}

func (g *GatewayError) ErrorMsg() string {
	return g.errorMsg
}

func (g *GatewayError) Code() uint64 {
	return g.code
}

