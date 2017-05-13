package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gorilla/mux"
)

const (
	templateDir   = "/opt/helm-ui/templates"
	defaultLayout = "/opt/helm-ui/templates/layout.html"
	homeDir       = "/opt/helm-ui"
)

func syncChartRepos(serverContext *ServerContext) {
	for {
		GetSynced(serverContext)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	log.Printf("Starting Helm UI version %s...\n", os.Getenv("VERSION"))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverContext := NewServerContext(ctx, os.Getenv("TILLER_HOST"), "default", "helmui")

	serverContext.tmpls = map[string]*template.Template{}
	serverContext.tmpls["home.html"] = template.Must(template.ParseFiles(path.Join(templateDir, "home.html"), defaultLayout))

	// continually reconcile the local repo cache with configmap repo list
	go syncChartRepos(serverContext)

	r := mux.NewRouter()

	r.HandleFunc("/", serverContext.HomeHandler)
	r.HandleFunc("/releases", serverContext.ReleaseHandler).Methods("GET")
	r.HandleFunc("/releases/{release}", serverContext.ReleaseHandler).Methods("GET", "PATCH", "DELETE", "OPTIONS")
	r.HandleFunc("/releases/{release}/history", serverContext.ReleaseHistoryHandler).Methods("GET")
	r.HandleFunc("/releases/{release}/rollback/{revision}", serverContext.ReleaseRevertHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/releases/{release}/diff/{revision}", serverContext.ReleaseDiffHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/repos", serverContext.HelmRepoHandler).Methods("POST", "GET", "OPTIONS")
	r.HandleFunc("/repos/{repo}", serverContext.HelmRepoHandler).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/repos/{repo}/charts", serverContext.HelmRepoChartsHandler).Methods("GET")
	r.HandleFunc("/repos/{repo}/charts/{chart}/install", serverContext.HelmRepoChartInstallHandler).Methods("POST", "OPTIONS")

	http.Handle("/", r)

	port := os.Getenv("PORT")
	log.Printf("Serving on port %s...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
