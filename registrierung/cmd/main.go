package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"training-fellow.de/registrierung"
	"training-fellow.de/registrierung/http/rest"
	"training-fellow.de/registrierung/mongodb"
	"training-fellow.de/registrierung/nats"
)

type configuration struct {
	mongoURL        string
	mongoDatabase   string
	mongoCollection string
	natsURL         string
	sslCertFileName string
	sslKeyFileName  string
}

func readConfig() configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error with config file: %v", err))
	}
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	return configuration{
		mongoURL:        viper.GetString("mongo.url"),
		mongoDatabase:   viper.GetString("mongo.database"),
		mongoCollection: viper.GetString("mongo.collection"),
		natsURL:         viper.GetString("nats.url"),
		sslCertFileName: viper.GetString("server.ssl.certFile"),
		sslKeyFileName:  viper.GetString("server.ssl.keyFile"),
	}
}

func main() {

	config := readConfig()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/registrierung", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registrierung.tmpl", gin.H{
			"title":  "Schulungsregistrierung",
			"action": "/registrierung",
		})
	})

	p := ginprometheus.NewPrometheus("gin")
	p.SetListenAddress(":9900")
	p.Use(router)

	notifier := nats.NewNotifier(config.natsURL)
	repository := mongodb.NewRepo(config.mongoURL, config.mongoDatabase, config.mongoCollection)
	service := &registrierung.RegistrierungsService{Notifier: notifier, Repository: repository}
	router.POST("/registrierung", rest.NewRegistrationHandler(service))
	router.POST("/confirmations", rest.NewConfirmationHandler(service))
	router.GET("/confirmations", rest.NewUnconfirmedListHandler(service))

	router.RunTLS(":8443", config.sslCertFileName, config.sslKeyFileName)

}
