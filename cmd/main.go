package main

import (
	"context"
	"fmt"
	"net/http"

	"template-validating-webhook/internal/utils"
	"template-validating-webhook/pkg/apis"

	"github.com/gorilla/mux"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	cert = utils.CertDir + "/tls.crt"
	key  = utils.CertDir + "/tls.key"
)

var log = logf.Log.WithName("template-validating-webhook")

func main() {
	log.Info("initializing server....")

	if err := utils.CreateCert(context.Background()); err != nil {
		fmt.Println(err, "failed to create cert")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", apis.CheckInstanceUpdatable).Methods("POST")

	http.Handle("/", r)

	if err := http.ListenAndServeTLS(":8443", cert, key, nil); err != nil {
		fmt.Println(err, "failed to initialize a server")
	}
	//TODO Add kube client and check instance
	//TODO whenever pod restart, update CABundle of ValidatingWebhookConfiguration CRD
}
