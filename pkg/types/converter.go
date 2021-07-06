package types

import (
	"goblog/pkg/logger"
	"strconv"
)

// Int64ToString 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StringToInt(idStr string) int {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogError(err)
	}
	return id
}

func Uint64ToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}
