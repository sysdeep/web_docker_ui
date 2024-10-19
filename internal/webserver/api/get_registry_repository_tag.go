package api

import (
	"hdu/internal/registry_client"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler
func (h *Api) GetRegistryRepositoryTag(c echo.Context) error {

	id := c.Param("id")
	tag := c.Param("tag")

	manifest, err := h.registry_client.GetManivestV2(id, tag)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	repo, err := h.registry_client.GetRepository(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newRegistryRepositoryTagResponse(manifest, repo, tag))
}

type registryRepositoryTagResponse struct {
	TagManifest tagManifest                `json:"tag_manifest"`
	Repository  registryRepositoryResponse `json:"repository"`
}

func newRegistryRepositoryTagResponse(manifest registry_client.ManifestV2,
	repository registry_client.RepositoryModel, tag_name string) registryRepositoryTagResponse {
	return registryRepositoryTagResponse{
		TagManifest: newTagManifest(manifest, tag_name),
		Repository:  newRegistryRepositoryResponse(repository),
	}
}

type tagManifest struct {
	Name              string               `json:"name"`
	SchemaVersion     int                  `json:"schema_version"`
	MediaType         string               `json:"media_type"`
	ConfigDescriptor  descriptorManifest   `json:"config"`
	LayersDescriptors []descriptorManifest `json:"layers"`

	// additional fields which not include in schema specification and need for this service only
	TotalSize     int64  `json:"total_size"`     // total compressed size of image data
	ContentDigest string `json:"content_digest"` // a main content digest using for delete image from registry

}

func newTagManifest(manifest registry_client.ManifestV2, tag_name string) tagManifest {
	layers := []descriptorManifest{}
	for _, client_layer := range manifest.LayersDescriptors {
		layers = append(layers, newDescriptorManifest(client_layer))
	}

	return tagManifest{
		Name:              tag_name,
		SchemaVersion:     manifest.SchemaVersion,
		MediaType:         manifest.MediaType,
		ConfigDescriptor:  newDescriptorManifest(manifest.ConfigDescriptor),
		LayersDescriptors: layers,

		TotalSize:     manifest.TotalSize,
		ContentDigest: manifest.ContentDigest,
	}
}

type descriptorManifest struct {
	MediaType string `json:"media_type"`
	Size      int64  `json:"size"`
	Digest    string `json:"digest"`
	// URLs      []string `json:"urls,omitempty"`
}

func newDescriptorManifest(descriptor registry_client.Descriptor) descriptorManifest {
	return descriptorManifest{
		MediaType: descriptor.MediaType,
		Size:      descriptor.Size,
		Digest:    descriptor.Digest,
	}
}

/*
{
    "SchemaVersion": 2,
    "MediaType": "application/vnd.docker.distribution.manifest.v2+json",
    "ConfigDescriptor": {
        "MediaType": "application/vnd.docker.container.image.v1+json",
        "Size": 11125,
        "Digest": "sha256:43ea642414f671627f38df3015d4eed142a444e27787b61b2d64adaa182f0d64"
    },
    "LayersDescriptors": [
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 67756585,
            "Digest": "sha256:976991e750b3e9c205474b08f51a0fab2141dbc50265607d725c27e6d2435e6a"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 303,
            "Digest": "sha256:7220018e814bafe7f173573f0092e00037fc8767e8db4400830ea66766472fc7"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 265,
            "Digest": "sha256:857187d7f3739113493cd78615ce185f9202ddd6f7dc316c1289775beda77389"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 1611,
            "Digest": "sha256:3167c69ccd9391431be7d60c5711a95524e11857bef8aeab9e546803f808e25e"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 211,
            "Digest": "sha256:a1c2c0c89889775690848c0e42d34d7efa53140b30fe11abff89a3c4355aaf4a"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 10750772,
            "Digest": "sha256:662b7ed40f765d4b857ee98cfdd3e9f5efc98c51f1b225e837ca8603e3371891"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 860979,
            "Digest": "sha256:157082bbda4913f681794cd199b9d3cb93adc3e7a493a555090f7365754a5c2e"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 106022454,
            "Digest": "sha256:654ff63eb7dab9fdc7f02ad1680af08ac2eb24414d509f4a706608d1769099c1"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 32,
            "Digest": "sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 100015518,
            "Digest": "sha256:23b4303e97e893b474ab7034d8f082a687c9ebf0a4a79d0ac71e14eed5fbf71e"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 36684155,
            "Digest": "sha256:20633f6ad99bfe0551085518ec1d5741f12acc6bf836a677bfb1e81911c43e8c"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 215,
            "Digest": "sha256:09d4e98cf886c8da3c0210811c80b49c614ea2a699a7d9cbd4d56584266a591a"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 292,
            "Digest": "sha256:45b2d0d37112a9c3f229d4cbcddf300e322fef54a9f3240ab7541d8c87da8744"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 164,
            "Digest": "sha256:1c42899766d85c49ff6412b3cd2d2f7a74c2517ecdf2ad6232de75e201da164f"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 5157020,
            "Digest": "sha256:187832a7d16ba17df6213ea86a35a7be729453dc77c31db722f1c74006f28abc"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 907,
            "Digest": "sha256:46ce4b659e1f2d339415f41b089d81614944c6c1062719b86d84b2248b743a53"
        },
        {
            "MediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
            "Size": 388,
            "Digest": "sha256:e6681bd1aff924f3c33a413335649f5ca336e11f49084d5786cfab637a709ca3"
        }
    ],
    "TotalSize": 327251871,
    "ContentDigest": "sha256:c2b2bc6403328f909db5a71d8c5deb4d2dc5c6fdbce584bfb54306a3169b6a5c"
}


*/
