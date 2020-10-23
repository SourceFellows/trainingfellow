package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
	"training-fellow.de/registrierung"
)

var registrationCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "trainingfellow_registration_handler_total",
	Help: "The total number of requests",
})

var bindingErrorCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "trainingfellow_registration_handler_bindingerror_total",
	Help: "The total number of request with binding problems",
})

var registrationErrorCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "trainingfellow_registration_handler_registrationerror_total",
	Help: "The total number of request with registration errors",
})

var registrationSuccessFulCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "trainingfellow_registration_handler_registration_total",
	Help: "The total number of successful requests",
})

func NewRegistrationHandler(ser *registrierung.RegistrierungsService) gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Info("new Registration call reached server")
		registrationCounter.Add(1)

		registrierung := &registrierung.Registrierung{}
		err := c.Bind(registrierung)
		if err != nil {
			log.Errorf("Could not bind Registrierung: %v", err)
			bindingErrorCounter.Add(1)
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}

		err = ser.HandleNewRegistrierung(registrierung)
		if err != nil {
			log.Errorf("Could not handle Registrierung: %v", err)
			registrationErrorCounter.Add(1)
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		registrationSuccessFulCounter.Add(1)
		c.Writer.WriteHeader(http.StatusCreated)

	}
}

func NewUnconfirmedListHandler(ser *registrierung.RegistrierungsService) gin.HandlerFunc {
	return func(c *gin.Context) {

		registrations, err := ser.GetUnconfirmedRegistrierungen()
		if err != nil {
			log.Errorf("Could not find unconfirmed registrations: %v", err)
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, registrations)
	}
}

func NewConfirmationHandler(ser *registrierung.RegistrierungsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		registrierungsID := c.DefaultPostForm("registrierungsID", "")
		err := ser.ConfirmRegistration(registrierungsID)
		if err != nil {
			log.Errorf("Could not handle confirmation: %v", err)
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
	}
}
