package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	templateDir   = "/var/www/templates/"
	defaultLayout = "/var/www/templates/layout.html"
	repoFile      = "/var/www/repositories.yaml"
)

func main() {

	log.Printf("Starting Helm UI version %s...\n", os.Getenv("VERSION"))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverContext := NewServerContext(ctx, os.Getenv("TILLER_HOST"), "default", "helmui")

	serverContext.tmpls = map[string]*template.Template{}
	serverContext.tmpls["home.html"] = template.Must(template.ParseFiles(templateDir+"home.html", defaultLayout))

	GetSynced(serverContext)

	r := mux.NewRouter()
	r.HandleFunc("/", serverContext.HomeHandler)
	r.HandleFunc("/releases", serverContext.ListReleases)
	r.HandleFunc("/repos", serverContext.HelmRepoHandler).Methods("POST", "GET")
	http.Handle("/", r)

	port := os.Getenv("PORT")
	log.Printf("Serving on port %s...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
