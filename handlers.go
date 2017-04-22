package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/ericchiang/k8s"
	"k8s.io/helm/pkg/helm"
)

type ServerContext struct {
	helmClient    *helm.Client
	k8sClient     *k8s.Client
	tmpls         map[string]*template.Template
	ctx           context.Context
	namespace     string
	configMapName string
}

func NewServerContext(ctx context.Context, host string, namespace string, configMapName string) *ServerContext {
	k8sClient, err := k8s.NewInClusterClient()
	if err != nil {
		log.Fatal(err)
	}
	return &ServerContext{
		helmClient:    helm.NewClient(helm.Host(host)),
		k8sClient:     k8sClient,
		ctx:           ctx,
		namespace:     namespace,
		configMapName: configMapName,
	}
}

func (c ServerContext) ListReleases(w http.ResponseWriter, r *http.Request) {
	releases, err := c.helmClient.ListReleases()
	if err != nil {
		log.Printf("failed to list releases: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

	w.Header().Set("Content-Type", "application/json")
	err = c.SaveHelmRepo(newRepo)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "POST":
		c.AddHelmRepoHandler(w, r)
		return
	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	default:
		repos, err := c.GetHelmRepos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
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
