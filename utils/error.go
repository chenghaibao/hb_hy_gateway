package utils

import "log"

func PanicError(err error) {
	// 判断结果
	if err != nil {
		// 抛出异常
		panic(err)
	}
}

func CheckError(err error) bool {
	// 判断结果
	if err != nil {
		// 打印错误
		log.Println(err)
		// 返回false
		return true
	}
	// 返回无错误
	return false
}
