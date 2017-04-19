package main

import (
	"fmt"
	"strings"

	"github.com/ericchiang/k8s/api/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

type HelmRepo struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Cache string `json:"cache"`
}

func (s *ServerContext) GetHelmRepos() ([]HelmRepo, error) {
	var repos []HelmRepo
	configMap, err := s.k8sClient.CoreV1().GetConfigMap(s.ctx, s.configMapName, "default")
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

func (s *ServerContext) SaveHelmRepo(r HelmRepo) error {
	// need to deal with the namespacing at some point
	namespace := "default"
	// try to get teh config map and create it if it does not
	configMap, err := s.k8sClient.CoreV1().GetConfigMap(s.ctx, s.configMapName, namespace)
	if err != nil {

		if !strings.Contains(err.Error(), fmt.Sprintf(`configmaps "%s" not found`, s.configMapName)) {
			return err
		}

		cfgmap := &v1.ConfigMap{
			Metadata: &metav1.ObjectMeta{
				Name:      &s.configMapName,
				Namespace: &namespace,
			},
		}

		configMap, err = s.k8sClient.CoreV1().CreateConfigMap(s.ctx, cfgmap)
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

	_, err = s.k8sClient.CoreV1().UpdateConfigMap(s.ctx, configMap)
	if err != nil {
		return err
	}

	return nil
}
