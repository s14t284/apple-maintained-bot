package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/usecase/repository"
)

func GetIPadHandler(ir repository.IPadRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)

		// parse request parameters
		err := req.ParseForm()
		if err != nil {
			log.Errorf("failed to parse resquest parameter [error][%w]", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		if req.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ipads, err := ir.FindIPadAll()
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
