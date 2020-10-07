package utils

import "time"

// GetReleaseYearAndMonth 商品が売られ始めた年月を表現するための時間を返却
func GetReleaseYearAndMonth(year int, month int) time.Time {
	timeZone, _ := time.LoadLocation("Asia/Tokyo")
	return time.Date(year, time.Month(month), 1, 9, 0, 0, 0, timeZone)
}
