package http

import (
	"hb_hy_gateway/utils"
	"strings"
)

const (
	SERVICE = "family,kids"
)

func CheckServiceMiddle(service string) bool {
	limitSer := strings.Split(SERVICE, `,`)
	isArray := utils.InStringArray(service, limitSer)
	return isArray
}
