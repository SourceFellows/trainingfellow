package main

import (
	"context"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"training-fellow.de/vorbereitung"

	"github.com/nats-io/nats.go"
)

func connectNATS(ctx context.Context, url string) (*nats.Conn, error) {
	var err error
	for {
		select {
		case <-ctx.Done():
			return nil, err
		default:
			var nc *nats.Conn
			nc, err = nats.Connect(url)
			if err == nil {
				return nc, err
			}
		}
	}

}

func main() {

	url := nats.DefaultURL
	if v, ok := os.LookupEnv("NATS_URL"); ok {
		url = v
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	nc, err := connectNATS(ctx, url)
	if err != nil {
		log.Fatalf("Could not connect to server because of %v", err)
	}
	defer nc.Close()
	log.Println("successfully connected to NATS server")

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	sub, err := c.Subscribe("traingfellow.registrierung.neu", func(registrierung *vorbereitung.Registrierung) {
		log.Info("neue Registrierung", registrierung)
	})

	if err != nil {
		log.Fatalf("Could not subscribe because of %v", err)
	}

	log.Info("Wait for Registrierung")
	wg.Wait()

	sub.Unsubscribe()
	sub.Drain()

}
