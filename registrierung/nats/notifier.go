package nats

import (
	log "github.com/sirupsen/logrus"

	"github.com/nats-io/nats.go"
	"training-fellow.de/registrierung"
)

func NewNotifier(url string) *notifier {
	return &notifier{url}
}

type notifier struct {
	url string
}

func (nn *notifier) InformAboutNewRegistrierung(registrierung *registrierung.Registrierung) error {

	notifierLogger := log.WithField("Registrierung", registrierung)
	notifierLogger.Info("Inform about new Registrierung")

	nc, err := nats.Connect(nn.url)
	if err != nil {
		notifierLogger.WithError(err).Error("Could not connect to server: ")
		return err
	}
	defer nc.Close()
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	return c.Publish("traingfellow.registrierung.neu", registrierung)

}
