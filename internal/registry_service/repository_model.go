package registry_service

// import "reg_ui/internal/registry_client"
//
// type RepositoryModel struct {
// 	Name string     `json:"name"`
// 	Uid  string     `json:"uid"`
// 	Tags []TagModel `json:"tags"`
// }
//
// type TagModel struct {
// 	Name string `json:"name"`
// 	Uid  string `json:"uid"`
// }
//
// // TODO: move to..
// func tags2models(tags []string) []TagModel {
// 	var result []TagModel
// 	for _, name := range tags {
// 		uid := name2id(name)
//
// 		result = append(result, TagModel{Name: name, Uid: uid})
// 	}
//
// 	return result
// }
//
// func newReposytoryModelFromClient(model *registry_client.RepositoryModel) RepositoryModel {
// 	return RepositoryModel{
// 		Name: model.Name,
// 		Uid:  name2id(model.Name),
// 		Tags: tags2models(model.Tags),
// 	}
// }
//
