package registry_client

/*
https://metanit.com/go/tutorial/9.6.php

*/
import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// scheme version of manifest file
	// for details about scheme version goto https://docs.docker.com/registry/spec/manifest-v2-2/
	manifestSchemeV2 = "application/vnd.docker.distribution.manifest.v2+json"

	//  It uniquely identifies content by taking a collision-resistant hash of the bytes.
	contentDigestHeader = "docker-content-digest"
)

var errDisabled = errors.New("registry client is disabled")

type RegistryClient struct {
	address string
	enabled bool
	client  *http.Client
}

// create client
// if address = "" - disabled
func NewRegistryClient(address string) *RegistryClient {

	enabled := len(address) > 0

	if !enabled {
		slog.Info("Registry client is disabled")
	}

	// for https ignoring
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr, Timeout: 5 * time.Second}

	return &RegistryClient{
		address: address,
		enabled: enabled,
		client:  &client,
	}
}

func (c *RegistryClient) IsEnabled() bool {
	return c.enabled
}

// TODO: check!!!
func (c *RegistryClient) APIVersionCheck() error {
	if !c.enabled {
		return errDisabled
	}

	url, _ := url.JoinPath(c.address, "v2")

	// TODO: вынести в отдельный метод - слишком много повторений
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", manifestSchemeV2)
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// GetCatalog
func (c *RegistryClient) GetCatalog(n int) (Catalog, error) {
	if !c.enabled {
		return Catalog{}, errDisabled
	}

	slog.Info("GetCatalog",
		"n", n)

	// make endpoint address
	url := c.make_url(fmt.Sprintf("/v2/_catalog?n=%d", n))
	slog.Debug("Get catalog url: " + url)

	// fetch
	body, err := c.make_get(url)
	if err != nil {
		return Catalog{}, err
	}

	// parse
	result := catalogResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Catalog{}, err
	}

	repos := []RepositoryListModel{}
	for _, row := range result.Repositories {
		repos = append(repos, newRepositoryListModel(row))
	}

	return Catalog{
		Repositories: repos,
	}, nil

}

// GetRepository
func (c *RegistryClient) GetRepository(id string) (RepositoryModel, error) {
	if !c.enabled {
		return RepositoryModel{}, errDisabled
	}

	// convert id to name
	image_name, err := id2name(id)
	if err != nil {
		return RepositoryModel{}, err
	}

	// send request
	url := c.make_url(fmt.Sprintf("/v2/%s/tags/list", image_name))

	body, err := c.make_get(url)
	if err != nil {
		return RepositoryModel{}, err
	}

	result := repositoryResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return RepositoryModel{}, err
	}

	tags_result := []string{}
	if result.Tags != nil {
		tags_result = result.Tags
	}

	return newRepositoryModel(result.Name, tags_result), nil

}

// GetManivestV2
func (c *RegistryClient) GetManivestV2(id string, tag_name string) (ManifestV2, error) {
	if !c.enabled {
		return ManifestV2{}, errDisabled
	}

	// convert id to name
	image_name, err := id2name(id)
	if err != nil {
		return ManifestV2{}, err
	}

	url := c.make_url(fmt.Sprintf("/v2/%s/manifests/%s", image_name, tag_name))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ManifestV2{}, err
	}
	req.Header.Add("Accept", manifestSchemeV2)
	res, err := c.client.Do(req)

	if err != nil {
		return ManifestV2{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ManifestV2{}, err
	}

	schema := manifestV2ResponseSchema{}
	err = json.Unmarshal(body, &schema)
	if err != nil {
		return ManifestV2{}, err
	}

	content_digest := res.Header.Get(contentDigestHeader)

	var layers_descriptors []Descriptor

	for _, ld := range schema.LayersDescriptors {
		layers_descriptors = append(layers_descriptors, c.cd_from_response(ld))
	}

	return ManifestV2{
		SchemaVersion:     schema.SchemaVersion,
		MediaType:         schema.MediaType,
		ConfigDescriptor:  c.cd_from_response(schema.ConfigDescriptor),
		LayersDescriptors: layers_descriptors,
		TotalSize:         schema.CalculateCompressedImageSize(),
		ContentDigest:     content_digest,
	}, nil

}

// DeleteTag will delete the manifest identified by name and reference. Note that a manifest can only be deleted by digest.
// A digest can be fetched from manifest get response header 'docker-content-digest'
// после удаления необходимо выполнить чистку
// docker exec -it registry bin/registry garbage-collect  /etc/docker/registry/config.yml
func (c *RegistryClient) RemoveManifest(id string, digest string) error {
	if !c.enabled {
		return errDisabled
	}

	// convert id to name
	image_name, err := id2name(id)
	if err != nil {
		return err
	}
	// curl -v --silent -H "Accept: application/vnd.docker.distribution.manifest.v2+json" \
	// -X DELETE http://127.0.0.1:5000/v2/ubuntu/manifests/sha256:7cc0576c7c0ec2384de5cbf245f41567e922aab1b075f3e8ad565f508032df17
	slog.Info("RemoveManifest", image_name, digest)

	url := c.make_url(fmt.Sprintf("/v2/%s/manifests/%s", image_name, digest))
	// fmt.Println("Client - remove: ", image_name, "/", digest)
	// url, _ := url.JoinPath(c.address, "v2", reposytoryName, "/manifests/", digest)
	// c.logger.Debug("sending request: " + url)
	//
	fmt.Println("-----------------------------------------")
	fmt.Println(url)
	fmt.Println("-----------------------------------------")

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", manifestSchemeV2)
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println("Delete body result ==================================================")
	fmt.Println("Status: ", res.Status)
	fmt.Println("StatusCode: ", res.StatusCode)
	fmt.Println("Body:", string(body))
	fmt.Println("=====================================================================")
	return nil

}

// private --------------------------------------------------------------------
func (c *RegistryClient) make_get(url string) ([]byte, error) {

	resp, err := c.client.Get(url)
	if err != nil {

		fmt.Println(err)
		return make([]byte, 0), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body, err
}

func (c *RegistryClient) make_url(part string) string {

	start := strings.TrimSuffix(c.address, "/")
	end := strings.TrimPrefix(part, "/")

	return start + "/" + end
}

func (c *RegistryClient) cd_from_response(data schema2Descriptor) Descriptor {
	return Descriptor{
		MediaType: data.MediaType,
		Size:      data.Size,
		Digest:    data.Digest,
	}
}

// NOTE: идея определять транспорт, но и без этого работает
// func makeHttpClient(address string) *http.Client{
// 	if strings.HasPrefix(address, "https:"){
//
// 	}
//
// }

// http models ----------------------------------------------------------------
type catalogResponse struct {
	Repositories []string `json:"repositories"`
}

type repositoryResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

/*

class RegistryClient(ClientInterface):
    _timeout = 5
    _manifest_scheme_v2_header = "application/vnd.docker.distribution.manifest.v2+json"
    _content_digest_header = "docker-content-digest"

    def __init__(self, params: RegistryClientParams):
        self._params = params
        urllib3.disable_warnings()



    def remove_manifest(self, image_name: str, digest: str):
        url = self._make_url(f"/v2/{image_name}/manifests/{digest}")
        headers = {
            'Accept': self._manifest_scheme_v2_header
        }
        resp = requests.delete(url, verify=False, timeout=self._timeout, headers=headers)
        resp.raise_for_status()

    def _make_url(self, part: str) -> str:
        start = self._params.registry_address
        if start.endswith('/'):
            start = start[:-1]

        end = part
        if end.startswith('/'):
            end = end[1:]

        return start + '/' + end



*/
