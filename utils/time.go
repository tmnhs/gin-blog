package utils

import "time"

const (
	TimeFormatSecond               = "2006-01-02 15:04:05"
	TimeFormatMinute               = "2006-01-02 15:04"
	TimeFormatDateV1               = "2006-01-02"
	TimeFormatDateV2               = "2006_01_02"
	Expiration       time.Duration = 24 * 60 * 60 * 1000 * 1000 * 1000 //过期时间
)
