package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ericchiang/k8s"
	"github.com/gorilla/mux"
	"k8s.io/helm/pkg/downloader"
	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/helm/helmpath"
	"k8s.io/helm/pkg/repo"
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
	switch r.Method {
	case "POST":
		c.AddHelmRepoHandler(w, r)
		return
	default:
		repos, err := c.GetHelmRepos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(w).Encode(repos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (c ServerContext) HelmRepoChartsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	home := helmpath.Home(homeDir)

	cacheIndex, err := repo.LoadIndexFile(home.CacheIndex(vars["repo"]))
	if err != nil {
		log.Printf("failed to load cache index: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	cacheIndex.SortEntries()

	var cvs []*repo.ChartVersion
	for _, chartVersions := range cacheIndex.Entries {
		// for now we only care about the first version (the latest)
		cvs = append(cvs, chartVersions[0])
	}
	jsonData, err := json.Marshal(cvs)
	if err != nil {
		log.Printf("failed to json marshal: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("failed to write: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c ServerContext) HelmRepoChartInstallHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	home := helmpath.Home(homeDir)

	chartDownloader := downloader.ChartDownloader{
		HelmHome: home,
	}
	tarDest, _, err := chartDownloader.DownloadTo(fmt.Sprintf("%s/%s", vars["repo"], vars["chart"]), "", "")
	if err != nil {
		log.Printf("failed to resolve chart version: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	resp, err := c.helmClient.InstallRelease(tarDest, c.namespace, helm.ValueOverrides([]byte("")))
	if err != nil {
		log.Printf("failed to install release: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, err := json.Marshal(resp.Release)
	if err != nil {
		log.Printf("failed to json marshal: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("failed to write: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c ServerContext) HomeHandler(w http.ResponseWriter, r *http.Request) {
	c.tmpls["home.html"].ExecuteTemplate(w, "base", map[string]interface{}{
		"message": "hello!",
	})
}
