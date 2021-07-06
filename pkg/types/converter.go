package types

import (
	"goblog/pkg/logger"
	"strconv"
)

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
