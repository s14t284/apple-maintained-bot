package handler

import (
	"encoding/json"
	"net/http"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/service"

	"github.com/labstack/gommon/log"
)

// GetMacHandler macのGetリクエストの API Handler
func GetMacHandler(ms service.MacService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		macs, err := ms.Find(&req)
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
}
