package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
	if releases.Count == 0 {
		return
	}
	io.WriteString(w, "<h1>Releases</h1>")
	io.WriteString(w, "<table><tr><th>Namespace</th><th>Name</th><th>Version</th></tr>")
	for _, r := range releases.GetReleases() {
		io.WriteString(w, "<tr>")
		io.WriteString(w, "<td>"+r.Namespace+"</td>")
		io.WriteString(w, "<td>"+r.Name+"</td>")
		io.WriteString(w, fmt.Sprintf("%s%d%s", "<td>", r.Version, "</td>"))
		io.WriteString(w, "</tr>")
	}
	io.WriteString(w, "</table>")
}

func main() {
	client := NewHelmClient(os.Getenv("TILLER_HOST"))

	http.HandleFunc("/", client.listReleases)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
