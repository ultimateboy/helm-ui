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
	for _, r := range releases.GetReleases() {
		io.WriteString(w, r.Name)
	}
}

func main() {
	client := NewHelmClient(os.Getenv("TILLER_HOST"))
	http.HandleFunc("/", client.listReleases)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
