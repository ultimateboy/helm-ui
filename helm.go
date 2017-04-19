package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ericchiang/k8s"
	"github.com/ericchiang/k8s/api/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

type HelmRepo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetHelmRepos(client *k8s.Client) ([]HelmRepo, error) {
	var repos []HelmRepo
	configMap, err := client.CoreV1().GetConfigMap(context.Background(), HELMUIConfigMap, "default")
	if err != nil {
		return repos, err
	}

	for k, v := range configMap.Data {
		repo := HelmRepo{
			Name: k,
			URL:  v,
		}
		repos = append(repos, repo)
	}

	return repos, nil
}

func SaveHelmRepo(client *k8s.Client, r HelmRepo) error {
	// need to deal with the namespacing at some point
	namespace := "default"
	ctx := context.Background()
	// try to get teh config map and create it if it does not
	configMap, err := client.CoreV1().GetConfigMap(ctx, HELMUIConfigMap, namespace)
	if err != nil {

		if !strings.Contains(err.Error(), fmt.Sprintf(`configmaps "%s" not found`, HELMUIConfigMap)) {
			return err
		}

		cfgmap := &v1.ConfigMap{
			Metadata: &metav1.ObjectMeta{
				Name:      &HELMUIConfigMap,
				Namespace: &namespace,
			},
		}

		configMap, err = client.CoreV1().CreateConfigMap(ctx, cfgmap)
		if err != nil {
			return err
		}
	}

	// dont save the repo if one exists by that name already
	if val, ok := configMap.Data[r.Name]; ok {
		return fmt.Errorf("A helm repo with the name '%s' is already in the system pointing to: %s", r.Name, val)
	}

	// if this config map is brand new then the data map needs to be initalized
	if configMap.Data == nil {
		configMap.Data = map[string]string{}
	}
	// now save the new repo url
	configMap.Data[r.Name] = r.URL

	_, err = client.CoreV1().UpdateConfigMap(ctx, configMap)
	if err != nil {
		return err
	}

	return nil
}
