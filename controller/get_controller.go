package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/gommon/log"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/usecase"
)

// GetController データの取得に関するhandlerを提供するController
type GetController struct {
	mu usecase.MacUseCase
	iu usecase.IPadUseCase
	wu usecase.WatchUseCase
}

// NewGetController GetControllerを初期化して返す
func NewGetController(mu usecase.MacUseCase, iu usecase.IPadUseCase, wu usecase.WatchUseCase) (*GetController, error) {
	if mu == nil {
		return nil, fmt.Errorf("mac usecase must not be nil")
	}
	if iu == nil {
		return nil, fmt.Errorf("ipad usecase must not be nil")
	}
	if wu == nil {
		return nil, fmt.Errorf("watch usecase must not be nil")
	}
	return &GetController{
		mu: mu,
		iu: iu,
		wu: wu,
	}, nil
}

// GetMacHandler macのGetリクエストの API Handler
func (gc *GetController) GetMacHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	req := model.MacRequestParam{}

	// parse request parameters
	err := r.ParseForm()
	if err != nil {
		log.Errorf("failed to parse request parameter [error][%w]", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	for k, v := range r.Form {
		switch k {
		case "name":
			req.Name = GetMacName(v[0])
		case "color":
			req.Color = GetColor(v[0])
		case "is_sold":
			req.IsSold = GetIsSold(v[0])
		case "max_amount":
			req.MaxAmount = GetAmount(v[0])
		case "min_amount":
			req.MinAmount = GetAmount(v[0])
		case "max_inch":
			req.MaxInch = GetInch(v[0])
		case "min_inch":
			req.MinInch = GetInch(v[0])
		case "max_storage":
			req.MaxStorage = GetStorage(v[0])
		case "min_storage":
			req.MinStorage = GetStorage(v[0])
		case "max_memory":
			req.MaxMemory = GetMemory(v[0])
		case "min_memory":
			req.MinMemory = GetMemory(v[0])
		}
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	macs, err := gc.mu.GetMacs(req)
	if err != nil {
		log.Errorf("failed to find mac information from db [error][%w]", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	obj, err := json.Marshal(macs)
	if err != nil {
		log.Errorf("failed to parse mac information to json [error][%w]", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(obj)
	if err != nil {
		log.Errorf("failed to write json [error][%w]", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetIPadHandler ipadのGetリクエストの API Handler
func (gc *GetController) GetIPadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	req := model.IPadRequestParam{}
	// parse request parameters
	err := r.ParseForm()
	if err != nil {
		log.Errorf("failed to parse resquest parameter [error][%w]", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	for k, v := range r.Form {
		switch k {
		case "name":
			req.Name = GetIPadName(v[0])
		case "color":
			req.Color = GetColor(v[0])
		case "is_sold":
			req.IsSold = GetIsSold(v[0])
		case "max_amount":
			req.MaxAmount = GetAmount(v[0])
		case "min_amount":
			req.MinAmount = GetAmount(v[0])
		case "max_inch":
			req.MaxInch = GetInch(v[0])
		case "min_inch":
			req.MinInch = GetInch(v[0])
		case "max_storage":
			req.MaxStorage = GetStorage(v[0])
		case "min_storage":
			req.MinStorage = GetStorage(v[0])
		}
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ipads, err := gc.iu.GetIPads(req)
	if err != nil {
		log.Errorf("failed to find ipad information from db [error][%w]", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	obj, err := json.Marshal(ipads)
	if err != nil {
		log.Errorf("failed to parse ipad information to json [error][%w]", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(obj)
	if err != nil {
		log.Errorf("failed to write json [error][%w]", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetWatchHandler apple watchのGetリクエストの API Handler
func (gc *GetController) GetWatchHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	req := model.WatchRequestParam{}
	// parse request parameters
	err := r.ParseForm()
	if err != nil {
		log.Errorf("failed to parse request parameter [error][%w]", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	for k, v := range r.Form {
		switch k {
		case "name":
			req.Name = GetWatchName(v[0])
		case "color":
			req.Color = GetColor(v[0])
		case "is_sold":
			req.IsSold = GetIsSold(v[0])
		case "max_amount":
			req.MaxAmount = GetAmount(v[0])
		case "min_amount":
			req.MinAmount = GetAmount(v[0])
		case "max_inch":
			req.MaxInch = GetInch(v[0])
		case "min_inch":
			req.MinInch = GetInch(v[0])
		case "max_storage":
			req.MaxStorage = GetStorage(v[0])
		case "min_storage":
			req.MinStorage = GetStorage(v[0])
		}
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	watches, err := gc.wu.GetWatches(req)
	if err != nil {
		log.Errorf("failed to find apple watch information from db [error][%w]", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	obj, err := json.Marshal(watches)
	if err != nil {
		log.Errorf("failed to parse apple watch information to json [error][%w]", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(obj)
	if err != nil {
		log.Errorf("failed to write json [error][%w]", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// HealthCheck ヘルスチェック用
func (gc *GetController) HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{\"message\": \"ok\"}"))
	if err != nil {
		log.Error(err)
	}
}

// -------------------------
// 以下は各メソッドで使う共通処理
// -------------------------

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
