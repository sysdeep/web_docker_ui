package registry_client

// type RegistryClient interface {
// 	GetCatalog(n int) (Catalog, error)
// 	GetRepository(image_name string) (Repository, error)
// 	GetManivestV2(image_name string, tag_name string) (ManifestV2, error)
// 	RemoveManifest(image_name string, digest string) error
// }

type Catalog struct {
	Repositories []RepositoryListModel
}

type RepositoryListModel struct {
	ID   string
	Name string
}

func newRepositoryListModel(name string) RepositoryListModel {
	return RepositoryListModel{
		ID:   name2id(name),
		Name: name,
	}
}

type RepositoryModel struct {
	ID   string
	Name string
	Tags []string
}

func newRepositoryModel(name string, tags []string) RepositoryModel {
	return RepositoryModel{
		ID:   name2id(name),
		Name: name,
		Tags: tags,
	}
}

type ManifestV2 struct {
	SchemaVersion     int
	MediaType         string
	ConfigDescriptor  Descriptor
	LayersDescriptors []Descriptor
	TotalSize         int64
	ContentDigest     string
}

type Descriptor struct {
	MediaType string
	Size      int64
	Digest    string
}

// @staticmethod
// def from_response(data: dict) -> 'Descriptor':
//     return Descriptor(media_type=data.get('mediaType'), size=data.get('size'), digest=data.get('digest'))
