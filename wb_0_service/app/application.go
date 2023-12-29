package app

import (
	"encoding/json"
	"fmt"
	"github.com/andrey-tsyn/wb_0/app/handlers"
	"github.com/andrey-tsyn/wb_0/app/models"
	"github.com/andrey-tsyn/wb_0/app/services"
	"github.com/andrey-tsyn/wb_0/app/subscriptions"
	"github.com/andrey-tsyn/wb_0/configuration"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var orderService *services.OrderStorageService = nil

func Start(config configuration.Config, db *sqlx.DB) {
	// Service initializing
	orderService = services.NewOrderStorageService(db)
	err := orderService.RestoreCache()
	if err != nil {
		log.Warnf("Restoring cache from database failed. Error: %s", err.Error())
	}

	// NATS
	natsClient, err := subscriptions.CreateClient(config.NatsUrl)
	if err != nil {
		log.Errorf("Can't connect to nats. Reason: %s", err.Error())
	} else {
		log.Info("NATS client started.")
	}
	defer natsClient.Close()

	subscribeToSubjects(natsClient)

	// Server
	router := mux.NewRouter()

	router.HandleFunc("/getOrder", handlers.GetOrderById(orderService)).Methods(http.MethodGet)
	router.HandleFunc("/addOrder", handlers.AddOrder(orderService)).Methods(http.MethodPost)

	log.Infof("Listening for client connections on 127.0.0.1:%s", config.Port)
	_ = http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router)
}

// TODO: Подумать, куда переместить, независимо от сервисов(по типу OrderStorageService и тд)
func subscribeToSubjects(nc *nats.Conn) {
	const orderAddSubject string = "order.add"

	_, err := nc.Subscribe(orderAddSubject, func(msg *nats.Msg) {
		validation := validator.New()
		order := models.Order{}
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Debugf("Can't parse json. Message data:\n%s", string(msg.Data))
			return
		}

		err = validation.Struct(&order)
		if err != nil {
			log.Debugf(
				"Struct validation failed.\nSubject: '%s'\nMessage data: '%s'\nError: %s",
				msg.Subject,
				string(msg.Data),
				err.Error(),
			)
			return
		}

		err = orderService.AddOrder(order)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
	})
	if err != nil {
		log.Warnf("Can't subscribe to subject 'order.add'. Error: %s", err.Error())
	} else {
		log.Debugf("NATS client subscribied to '%s' subject", orderAddSubject)
	}
}
