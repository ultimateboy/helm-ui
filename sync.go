package main

import (
	"fmt"
	"log"
	"strings"

	"k8s.io/helm/pkg/repo"
)

func GetSynced() {
	f, err := repo.LoadRepositoriesFile(repoFile)
	if err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			log.Fatalln(err)
		}
		f = repo.NewRepoFile()
	}
	fmt.Printf("%+v\n", f)
}
