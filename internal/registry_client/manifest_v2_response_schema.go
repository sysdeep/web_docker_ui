package registry_client

// TODO: copyed from https://github.com/zebox/registry-admin/blob/master/app/registry/registry.go

// ManifestSchemaV2 is V2 format schema for docker image manifest file which contain information about docker image, such as layers, size, and digest
// https://docs.docker.com/registry/spec/manifest-v2-2/#image-manifest-field-descriptions
type manifestV2ResponseSchema struct {
	SchemaVersion     int                 `json:"schemaVersion"`
	MediaType         string              `json:"mediaType"`
	ConfigDescriptor  schema2Descriptor   `json:"config"`
	LayersDescriptors []schema2Descriptor `json:"layers"`

	// additional fields which not include in schema specification and need for this service only
	TotalSize     int64  `json:"total_size"`     // total compressed size of image data
	ContentDigest string `json:"content_digest"` // a main content digest using for delete image from registry
}

type schema2Descriptor struct {
	MediaType string `json:"mediaType"`
	Size      int64  `json:"size"`
	Digest    string `json:"digest"`
	// URLs      []string `json:"urls,omitempty"`
}

// calculateCompressedImageSize will iterate with image layers in fetched manifest file and append size of each layers to TotalSize field
func (m *manifestV2ResponseSchema) CalculateCompressedImageSize() int64 {

	result := m.TotalSize

	for _, v := range m.LayersDescriptors {
		result += v.Size
	}

	return result
}

/*
{
   "schemaVersion": 2,
   "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
   "config": {
      "mediaType": "application/vnd.docker.container.image.v1+json",
      "size": 13672,
      "digest": "sha256:5e63a2a08d4a737fe78d08715065794a15c82ac652360941f84cabf6ad5f6863"
   },
   "layers": [
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 2814446,
         "digest": "sha256:a0d0a0d46f8b52473982a3c466318f479767577551a53ffc9074c9fa7035982e"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 1707843,
         "digest": "sha256:153eea49496a46a69cf5f48803e9014824a6be1d3e04f9ee47cd3f395aba6d76"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 1260,
         "digest": "sha256:11efd0df1fcb3f56e825ef3449f199b261f44e0130a10cb77fcf78339cb88173"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 268,
         "digest": "sha256:b3f3214c344df86e90fe2ae24d643d9dbc5dcfd6f229f4ee58f7c60c6f0cc895"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 10392314,
         "digest": "sha256:11801e573c1ee31d53c6decdf5375edd55e8decaffec447068657a3233118e31"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 495,
         "digest": "sha256:831aca17040b231f3b58011e30b649d9f23aa0c6951694b2c36d798f92c5c6fc"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 14549660,
         "digest": "sha256:b73cdf572c888363485d521caa6a3b2e5469c28e19689396f4246a0b64e69fc6"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 2264,
         "digest": "sha256:f38af604656f710789428a34236a76f9c6ffb7c79b52c721539e2d7c4f3bbaba"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 17831,
         "digest": "sha256:b25f05d07d18cfe3b393279fc6f7461a7557e355c5351f4d6d8b50db32a4d66f"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 307,
         "digest": "sha256:4af324acc19043c0a4951ab4037f8e30e1655673112928e378333c70b92aaa14"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 1414,
         "digest": "sha256:94a864fb90df99395c5c35ec2a20edbbfbde9fbe3371dd5b3f297749c2f04bf1"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 3947823,
         "digest": "sha256:66840e788c2bf52f537e4ce90c7efda5d636842fbe8ce4304aa13a1c6eb301c5"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 1491,
         "digest": "sha256:6c85195b5f246c481f06577c9d906b7a8f1a0b3a43839dc16ad770101971725a"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 583042,
         "digest": "sha256:1f346ec26e1bc472672dc7b44f8364d0fad32bfc70bbb13e4101598d39a48e3d"
      },
      {
         "mediaType": "application/vnd.docker.image.rootfs.diff.tar.gzip",
         "size": 499,
         "digest": "sha256:b9d0c4d5910ad71d33274ad4a3404f16cdbbd9def34707305607e8d92897efc2"
      }
   ]
}

*/
