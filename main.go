package main

import (
	"log"
	"os"

	"k8s.io/helm/pkg/helm"
)

func main() {
	client := helm.NewClient(helm.Host(os.Getenv("TILLER_HOST")))
	client.Option()
	releases, err := client.ListReleases()
	if err != nil {
		log.Fatalf("failed to get helm client: %v", err)
	}
	log.Println("RELEASES:")
	for _, r := range releases.GetReleases() {
		log.Println(r.Name)
	}

	select {}
}
