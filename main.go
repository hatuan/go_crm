package main

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/routers"
	"erpvietnam/crm/settings"
	"net/http"

	"erpvietnam/crm/middleware"

	"github.com/codegangsta/negroni"
)

func main() {
	router := routers.InitRoutes()
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), middleware.NewContext(), negroni.Wrap(router))
	log.WithFields(log.Fields{"address": settings.Settings.ListenURL}).Info("Running on address")

	//err := http.ListenAndServe(settings.Settings.ListenURL, n)
	err := http.ListenAndServeTLS(settings.Settings.ListenURL, settings.Settings.CertKeyPath, settings.Settings.PrivateKeyPath, n)
	if err != nil {
		log.Panic(err.Error())
	}
}
