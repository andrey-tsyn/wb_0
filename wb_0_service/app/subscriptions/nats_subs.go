package subscriptions

import (
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

func CreateClient(connectionUrl string, options ...nats.Option) (*nats.Conn, error) {
	log.Debugf("NATS Connection url: %s; Options: %+v", connectionUrl, options)
	nc, err := nats.Connect(connectionUrl, options...)
	if err != nil {
		log.Error("Can't connect to nats server:")
		log.Error(err.Error())
		return nil, err
	}
	return nc, nil
}
