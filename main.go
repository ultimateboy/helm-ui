package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ericchiang/k8s"
	"github.com/ericchiang/k8s/api/v1"
	"github.com/gorilla/mux"

	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
	"k8s.io/helm/pkg/helm"
)

var (
	K8sClient       *k8s.Client
	HELMUIConfigMap = "helmui"
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

type HelmRepo struct {
	Name string
	URL  string
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

func AddHelmRepoHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newRepo HelmRepo
	err := decoder.Decode(&newRepo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer r.Body.Close()

	err = SaveHelmRepo(K8sClient, newRepo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		payload, err := json.MarshalIndent(newRepo, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		io.WriteString(w, string(payload))
	}

}

func HelmRepoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		AddHelmRepoHandler(w, r)
		return
	default:
		repos, err := GetHelmRepos(K8sClient)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		payload, err := json.MarshalIndent(repos, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		io.WriteString(w, string(payload))
	}

}

func main() {
	helmClient := NewHelmClient(os.Getenv("TILLER_HOST"))
	var err error
	K8sClient, err = k8s.NewInClusterClient()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", helmClient.listReleases)
	r.HandleFunc("/repo", HelmRepoHandler).Methods("POST", "GET")
	http.Handle("/", r)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
