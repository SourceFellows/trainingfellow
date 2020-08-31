package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{
		"Firstname":             {"Peter"},
		"Lastname":              {"Lustig"},
		"DatenschutzAkzeptiert": {"true"},
		"Datum":                 {"2020-10-12"},
		"Schulungscode":         {"GO.EINF"},
		"Email":                 {"info@source-fellows.com"},
	}

	//disable TLS check for local tests
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true
	http.DefaultClient.Transport = &http.Transport{TLSClientConfig: tlsConfig}

	res, err := http.PostForm("https://localhost:8443/registrierung", data)
	if err != nil {
		panic(err)
	}
	log.Println(res.Status)
}
