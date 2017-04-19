package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	HELMUIConfigMap = "helmui"
)

const (
	templateDir   = "/var/www/templates/"
	defaultLayout = "/var/www/templates/layout.html"
	repoFile      = "/var/www/repositories.yaml"
)

func main() {

	serverContext := NewServerContext(os.Getenv("TILLER_HOST"))
	serverContext.tmpls = map[string]*template.Template{}
	serverContext.tmpls["home.html"] = template.Must(template.ParseFiles(templateDir+"home.html", defaultLayout))

	GetSynced(serverContext.k8sClient)

	r := mux.NewRouter()
	r.HandleFunc("/", serverContext.HomeHandler)
	r.HandleFunc("/releases", serverContext.listReleases)
	r.HandleFunc("/repos", serverContext.HelmRepoHandler).Methods("POST", "GET")
	http.Handle("/", r)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
