import { join_url } from '@src/routes';

export class RegistryService {
  private base_url: string;
  constructor(base_url: string) {
    this.base_url = base_url;
  }

  async get_repositories(): Promise<RepositoryListModel[]> {
    const response = await fetch(
      join_url(this.base_url, '/api/registry/repositories'),
    );

    const data = (await response.json()) as RepositoriesResponse;

    return data.repositories || [];
  }

  async get_repository(id: string): Promise<RepositoryModel> {
    const response = await fetch(
      join_url(this.base_url, '/api/registry/repository/' + id),
    );

    const data = (await response.json()) as RepositoryModel;

    return data;
  }

  async get_repository_tag(
    repository_id: string,
    tag: string,
  ): Promise<ReposytoryTagResponse> {
    const response = await fetch(
      join_url(
        this.base_url,
        `/api/registry/repository_tag/${repository_id}/${tag}`,
      ),
    );

    const data = (await response.json()) as ReposytoryTagResponse;

    return data;
  }

  async remove_tag(reposytory_id: string, tag: string): Promise<void> {
    await fetch(
      join_url(
        this.base_url,
        `/api/registry/repository/${reposytory_id}/${tag}`,
      ),
      {
        method: 'DELETE',
      },
    );
  }
}

// models ---------------------------------------------------------------------
interface RepositoriesResponse {
  repositories: RepositoryListModel[];
}

export interface RepositoryListModel {
  id: string;
  name: string;
}

export type RepositoryModel = {
  id: string;
  name: string;
  tags: string[];
};

export type ReposytoryTagResponse = {
  repository: RepositoryModel;
  tag_manifest: TagManifest;
};

export type TagManifest = {
  name: string;
  schema_version: number;
  media_type: string;
  config: DescriptorManifest;
  layers: DescriptorManifest[];
  total_size: number;
  content_digest: string;
};

type DescriptorManifest = {
  media_type: string;
  size: number;
  digest: string;
};
