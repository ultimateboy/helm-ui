package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"k8s.io/helm/pkg/helm"
)

type HelmClient struct {
	c *helm.Client
}

func NewHelmClient(host string) *HelmClient {
	return &HelmClient{
		c: helm.NewClient(helm.Host(host)),
	}
}

func (c HelmClient) listReleases(w http.ResponseWriter, r *http.Request) {
	releases, err := c.c.ListReleases()
	if err != nil {
		log.Printf("failed to list releases: %v", err)
		return
	}
	for _, r := range releases.GetReleases() {
		io.WriteString(w, r.Name)
	}
}

type HelmRepo struct {
	Name string
	URL  string
}

func AddHelmRepoHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newRepo HelmRepo
	err := decoder.Decode(&newRepo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer r.Body.Close()

	io.WriteString(w, fmt.Sprintf("%+v\n", newRepo))

}

func main() {
	client := NewHelmClient(os.Getenv("TILLER_HOST"))
	r := mux.NewRouter()
	r.HandleFunc("/", client.listReleases)
	r.HandleFunc("/repo/add", AddHelmRepoHandler).Methods("POST")
	http.Handle("/", r)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
