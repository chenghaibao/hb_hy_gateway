package utils

import (
	"math/rand"
	"time"
)

var traceIds = make(chan string, 1024)

// traceIdRand不支持并发调用，不要存在并发情况
var traceIdRand = rand.New(rand.NewSource(time.Now().UnixNano() + 111))

func init() {
	go func() {
		for {
			traceIds <- genTraceId()
		}
	}()
}

func genTraceId() string {
	const (
		chars         = "2345678abcdefhjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXY"
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)

	var cache = int64(0)
	var remain = 0
	var getNextRandom = func() uint8 {
		for {
			if remain == 0 || cache == 0 {
				cache, remain = traceIdRand.Int63(), letterIdxMax
			}

			idx := int(cache & letterIdxMask)
			cache >>= letterIdxBits
			remain--

			if idx < len(chars) {
				return chars[idx]
			}
		}
	}

	result := make([]byte, 16)
	for i := range result {
		result[i] = getNextRandom()
	}

	return string(result)
}

func GetTraceId() string {
	return <-traceIds
}
