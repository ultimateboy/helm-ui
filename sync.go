package main

import (
	"log"
	"os"
	"strings"

	"k8s.io/helm/pkg/helm/helmpath"
	"k8s.io/helm/pkg/repo"
)

func GetSynced(ctx *ServerContext) {
	changed := false

	home := helmpath.Home(homeDir)

	f, err := repo.LoadRepositoriesFile(home.RepositoryFile())
	if err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			log.Println(err)
		}
		f = repo.NewRepoFile()
		changed = true
	}

	stored, err := ctx.GetHelmRepos()
	if err != nil {
		log.Println(err)
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
			if _, err := os.Stat(home.CacheIndex(r.Name)); err == nil {
				err = os.Remove(home.CacheIndex(r.Name))
				if err != nil {
					log.Println(err)
				}
			}
			changed = true
		}
	}

	// add new repos to the repositories file
	for _, r := range stored {
		if !f.Has(r.Name) {
			cacheIndex := home.CacheIndex(r.Name)
			newRepo := &repo.Entry{
				Name:  r.Name,
				URL:   r.URL,
				Cache: cacheIndex,
			}
			chartRepo, err := repo.NewChartRepository(newRepo)
			if err != nil {
				log.Println(err)
			}
			if err := chartRepo.DownloadIndexFile(home.Cache()); err != nil {
				log.Printf("Looks like %q is not a valid chart repository or cannot be reached: %s", r.URL, err.Error())
			}

			f.Update(newRepo)
			changed = true
		}
	}

	if changed {
		err = f.WriteFile(home.RepositoryFile(), os.FileMode(0644))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
