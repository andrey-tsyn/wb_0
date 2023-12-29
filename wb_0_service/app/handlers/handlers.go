package handlers

import (
	"encoding/json"
	"github.com/andrey-tsyn/wb_0/app/models"
	"github.com/andrey-tsyn/wb_0/app/services"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

func AddOrder(orderService *services.OrderStorageService) http.HandlerFunc {
	validation := validator.New()

	return func(w http.ResponseWriter, r *http.Request) {
		setDefaultHeaders(w)
		body, err := io.ReadAll(r.Body)

		if err != nil {
			badRequest(w, err.Error())
		}

		order := models.Order{}

		err = json.Unmarshal(body, &order)

		if err != nil {
			badRequest(w, err.Error())
			return
		}

		err = validation.Struct(order)

		if err != nil {
			badRequest(w, err.Error())
			return
		}

		err = orderService.AddOrder(order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetOrderById(orderService *services.OrderStorageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setDefaultHeaders(w)
		id := r.URL.Query().Get("id")
		if len(id) == 0 {
			badRequest(w, "Query param 'id' must be set.")
			return
		}

		order, err := orderService.GetOrderById(id)
		if err != nil {
			badRequest(w, err.Error())
			return
		}

		jsonData, err := json.Marshal(order)

		if err != nil {
			badRequest(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonData)
	}
}

func badRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(msg))
}

func setDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
