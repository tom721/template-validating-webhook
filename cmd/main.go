package main

import (
	"net/http"

	"template-validating-webhook/pkg/apis"

	"github.com/gorilla/mux"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var log = logf.Log.WithName("template-validating-webhook")

func main() {
	logf.SetLogger(zap.Logger(true))
	log.Info("initializing server....")

	// cert := "/etc/webhook/certs/cert.pem"
	// key := "/etc/webhook/certs/key.pem"
	// listenOn := "0.0.0.0:8443"

	r := mux.NewRouter()
	r.HandleFunc("/", apis.CheckInstanceUpdatable).Methods("POST")

	http.Handle("/", r)
	// TODO) Convert to HTTPS
	if err := http.ListenAndServe(":8443", nil); err != nil {
		log.Error(err, "failed to initialize a server")
	}
}
