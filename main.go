package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ericchiang/k8s"
	"github.com/gorilla/mux"
)

var (
	K8sClient       *k8s.Client
	HELMUIConfigMap = "helmui"
)

func main() {
	helmClient := NewHelmClient(os.Getenv("TILLER_HOST"))
	var err error
	K8sClient, err = k8s.NewInClusterClient()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", helmClient.listReleases)
	r.HandleFunc("/repo", HelmRepoHandler).Methods("POST", "GET")
	http.Handle("/", r)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
