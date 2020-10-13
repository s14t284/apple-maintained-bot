package utils

import (
	"fmt"
	"time"
)

// GetReleaseYearAndMonth 商品が売られ始めた年月を表現するための時間を返却
func GetReleaseYearAndMonth(year int, month int) time.Time {
	timeZone, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Println(year)
	fmt.Println(time.Month(month))
	return time.Date(year, time.Month(month), 1, 9, 0, 0, 0, timeZone)
}
