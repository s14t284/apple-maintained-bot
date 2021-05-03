package handler

import (
	"encoding/json"
	"net/http"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
	"github.com/s14t284/apple-maitained-bot/utils"

	"github.com/labstack/gommon/log"
)

// GetIPadHandler ipadのGetリクエストの API Handler
func GetIPadHandler(ir database.IPadRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
				req.Name = utils.GetIPadName(v[0])
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
			}
		}

		if r.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ipads, err := ir.FindIPad(&req)
		if err != nil {
			log.Errorf("failed to find ipad information from db [error][%w]", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		obj, err := json.Marshal(ipads)
		if err != nil {
			log.Errorf("failed to parse ipad infomation to json [error][%w]", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		_, err = w.Write(obj)
		if err != nil {
			log.Errorf("failed to write json [error][%w]", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
