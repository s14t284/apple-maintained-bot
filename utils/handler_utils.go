package utils

import (
	"strconv"
	"strings"

	"github.com/labstack/gommon/log"
)

// GetColor リクエストパラメータから色を取得
func GetColor(str string) string {
	switch str {
	case "ゴールド", "シルバー", "スペースグレイ":
		return str
	default:
		return ""
	}
}

// GetInch リクエストパラメータからインチ数を取得
func GetInch(inch string) float64 {
	val, err := strconv.ParseFloat(inch, 64)
	if err != nil {
		log.Warnf("failed to parse inch size from request parameter [error][%w]", err)
		return 0.0
	}
	return val
}

// GetAmount リクエストパラメータから金額を取得
func GetAmount(amount string) int {
	val, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Warnf("failed to parse amount from request parameter [error][%w]", err)
		return 0
	}
	return int(val)
}

// GetStorage リクエストパラメータから金額を取得
func GetStorage(storage string) int {
	val, err := strconv.ParseInt(storage, 10, 64)
	if err != nil {
		log.Warnf("failed to storage from request parameter [error][%w]", err)
		return 0
	}
	return int(val)
}

// GetIsSold リクエストパラメータから売り切れているかを取得
// isSoldで絞り込みをするのはrepositoryで行うので、文字列をそのまま返す
func GetIsSold(str string) string {
	return str
}

// GetMemory リクエストパラメータからメモリを取得
func GetMemory(storage string) int {
	val, err := strconv.ParseInt(storage, 10, 64)
	if err != nil {
		log.Warnf("failed to memory from request parameter [error][%w]", err)
		return 0
	}
	return int(val)
}

// GetMacName リクエストパラメータからMacの名前を取得
func GetMacName(str string) string {
	switch strings.ToLower(str) {
	case "pro":
		return "MacBook Pro"
	case "air":
		return "MacBook Air"
	case "macbook":
		return "MacBook"
	case "mini":
		return "Mac mini"
	case "macpro":
		return "Mac Pro"
	default:
		return ""
	}
}

// GetIPadName リクエストパラメータからIPadの名前を取得
func GetIPadName(str string) string {
	switch strings.ToLower(str) {
	case "mini":
		return "IPad mini"
	case "pro":
		return "IPad Pro"
	case "unmarked":
		return "IPad"
	default:
		return ""
	}
}

// GetWatchName リクエストパラメータからapple watchの名前を取得
func GetWatchName(str string) string {
	return strings.ToTitle(str)
}
