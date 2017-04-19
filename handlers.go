package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/ericchiang/k8s"
	"k8s.io/helm/pkg/helm"
)

type ServerContext struct {
	helmClient *helm.Client
	k8sClient  *k8s.Client
	tmpls      map[string]*template.Template
}

func NewServerContext(host string) *ServerContext {
	k8sClient, err := k8s.NewInClusterClient()
	if err != nil {
		log.Fatal(err)
	}
	return &ServerContext{
		helmClient: helm.NewClient(helm.Host(host)),
		k8sClient:  k8sClient,
	}
}

func (c ServerContext) listReleases(w http.ResponseWriter, r *http.Request) {
	releases, err := c.helmClient.ListReleases()
	if err != nil {
		log.Printf("failed to list releases: %v", err)
		return
	}
	err = json.NewEncoder(w).Encode(releases.GetReleases())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c ServerContext) AddHelmRepoHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newRepo HelmRepo
	err := decoder.Decode(&newRepo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer r.Body.Close()

	err = SaveHelmRepo(c.k8sClient, newRepo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		err = json.NewEncoder(w).Encode(newRepo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func (c ServerContext) HelmRepoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		c.AddHelmRepoHandler(w, r)
		return
	default:
		repos, err := GetHelmRepos(c.k8sClient)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(w).Encode(repos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func (c ServerContext) HomeHandler(w http.ResponseWriter, r *http.Request) {
	c.tmpls["home.html"].ExecuteTemplate(w, "base", map[string]interface{}{
		"message": "hello!",
	})
}
