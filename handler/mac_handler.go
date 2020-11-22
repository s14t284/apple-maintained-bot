package handler

import (
	"encoding/json"
	"net/http"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/utils"

	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/usecase/repository"
)

// GetMacHandler macのGetリクエストの API Handler
func GetMacHandler(mr repository.MacRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		req := model.MacRequestParam{}

		// parse request parameters
		err := r.ParseForm()
		if err != nil {
			log.Errorf("failed to parse resquest parameter [error][%w]", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		for k, v := range r.Form {
			switch k {
			case "name":
				req.Name = utils.GetMacName(v[0])
			case "color":
				req.Color = utils.GetColor(v[0])
			case "is_sold":
				req.IsSold = utils.GetIsSold(v[0])
			case "max_amount":
				req.MaxAmount = utils.GetAmount(v[0])
			case "min_amount":
				req.MinAmount = utils.GetAmount(v[0])
			case "max_inch":
				req.MaxInch = utils.GetInch(v[0])
			case "min_inch":
				req.MinInch = utils.GetInch(v[0])
			case "max_storage":
				req.MaxStorage = utils.GetStorage(v[0])
			case "min_storage":
				req.MinStorage = utils.GetStorage(v[0])
			case "max_memory":
				req.MaxMemory = utils.GetMemory(v[0])
			case "min_memory":
				req.MinMemory = utils.GetMemory(v[0])
			}
		}

		if r.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		macs, err := mr.FindMac(&req)
		if err != nil {
			log.Errorf("failed to find mac information from db [error][%w]", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		obj, err := json.Marshal(macs)
		if err != nil {
			log.Errorf("failed to parse mac infomation to json [error][%w]", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		_, err = w.Write(obj)
		if err != nil {
			log.Errorf("failed to write json [error][%w]", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
