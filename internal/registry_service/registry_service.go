package registry_service

import "hdu/internal/registry_client"

type RegistryService struct {
	client *registry_client.RegistryClient
}

func NewRegistryService(client *registry_client.RegistryClient) *RegistryService {

	return &RegistryService{
		client: client,
	}
}

// func (r *RegistryService) GetRepositories() ([]RepositoryModel, error) {
//
// 	err := r.client.APIVersionCheck()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "wrong api version check")
// 	}
//
// 	var result []RepositoryModel
//
// 	catalog, err := r.client.GetCatalog()
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	for _, repo := range catalog.Repositories {
// 		uid := name2id(repo)
// 		service_repo := RepositoryModel{
// 			Name: repo,
// 			Uid:  uid,
// 		}
// 		result = append(result, service_repo)
// 	}
//
// 	return result, nil
// }
//
// func (r *RegistryService) GetRepository(uid string) (RepositoryModel, error) {
// 	var result RepositoryModel
//
// 	name, err := id2name(uid)
// 	if err != nil {
// 		return result, err
// 	}
//
// 	repo_model, err := r.client.GetRepositoryTags(name)
// 	if err != nil {
// 		return result, err
// 	}
//
// 	result.Uid = uid
// 	result.Name = repo_model.Name
// 	result.Tags = tags2models(repo_model.Tags)
//
// 	return result, nil
// }
//
// func (r *RegistryService) GetManifest(reposytory_uid string, tag_uid string) (ManifestModel, error) {
// 	var model ManifestModel
//
// 	repo_name, err := id2name(reposytory_uid)
// 	if err != nil {
// 		return model, err
// 	}
//
// 	tag_name, err := id2name(tag_uid)
// 	if err != nil {
// 		return model, err
// 	}
//
// 	manifest_result, err := r.client.GetManifestV2(repo_name, tag_name)
// 	if err != nil {
// 		return model, err
// 	}
//
// 	repo, err := r.client.GetRepositoryTags(repo_name)
//
// 	repo_model := newReposytoryModelFromClient(repo)
// 	tag_model := TagModel{Name: tag_name, Uid: tag_uid}
// 	model = makeManifestModel(manifest_result, repo_model, tag_model)
//
// 	return model, err
// }
//
// func (r *RegistryService) RemoveTag(repoUid string, digest string) error {
//
// 	repo_name, err := id2name(repoUid)
// 	if err != nil {
// 		return err
// 	}
//
// 	err = r.client.RemoveTag(repo_name, digest)
//
// 	return err
// }
