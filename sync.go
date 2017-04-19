package main

import (
	"log"
	"os"
	"strings"

	"github.com/ericchiang/k8s"
	"k8s.io/helm/pkg/repo"
)

func GetSynced(client *k8s.Client) {
	changed := false

	f, err := repo.LoadRepositoriesFile(repoFile)
	if err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			log.Fatalln(err)
		}
		f = repo.NewRepoFile()
		changed = true
	}

	stored, err := GetHelmRepos(client)
	if err != nil {
		log.Fatalln(err)
	}

	// remove any repositories that are not in the configmap
	for _, r := range f.Repositories {
		found := false
		for _, rs := range stored {
			if r.Name == rs.Name {
				found = true
			}
		}

		if !found {
			f.Remove(r.Name)
			changed = true
		}
	}

	// add new repos to the repositories file
	for _, r := range stored {
		if !f.Has(r.Name) {
			f.Add(&repo.Entry{
				Name:  r.Name,
				URL:   r.URL,
				Cache: r.Cache,
			})
			changed = true
		}
	}

	if changed {
		err = f.WriteFile(repoFile, os.FileMode(0644))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
