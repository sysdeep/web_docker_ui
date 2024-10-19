import PageTitle from '@src/components/page_title';
import {
  RegistryService,
  RepositoryModel,
  TagManifest,
} from '@src/services/registry_service';
import { useConfiguration } from '@src/store/configuration';
import React, { useEffect, useMemo, useState } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import DetailsFrame from './details_frame';
// import RepositoryFrame from './repository_frame';

export default function RepositoryTagPage() {
  const { id, tag } = useParams();
  // const navigate = useNavigate();
  const { configuration } = useConfiguration();

  const registry_service = useMemo(() => {
    return new RegistryService(configuration.base_url);
  }, []);

  const [manifest, setManifest] = useState<TagManifest | null>();
  const [repository, setReposytory] = useState<RepositoryModel | null>();
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = () => {
    setLoading(true);
    registry_service
      .get_repository_tag(id, tag)
      .then((repo) => {
        setManifest(repo.tag_manifest);
        setReposytory(repo.repository);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    console.log('page repository tag mounted');
    refresh();
  }, []);

  // const on_repository_remove = () => {
  //   console.log('repo remove page');
  //   // if (secret) {
  //   //   secret_service
  //   //     .remove_secret(id)
  //   //     .then(() => {
  //   //       console.log('remove ok');
  //   //       navigate(route.secrets);
  //   //     })
  //   //     .catch((err) => {
  //   //       console.log(err);
  //   //     });
  //   // }
  // };

  // const on_tag_remove = (tag: string) => {
  //   console.log('tag remove page: ' + tag);
  //   // if (secret) {
  //   //   secret_service
  //   //     .remove_secret(id)
  //   //     .then(() => {
  //   //       console.log('remove ok');
  //   //       navigate(route.secrets);
  //   //     })
  //   //     .catch((err) => {
  //   //       console.log(err);
  //   //     });
  //   // }
  // };

  // const page_name = repository ? repository.name : id;

  return (
    <div>
      <PageTitle>Repository tag: ----- </PageTitle>
      <div>{id}</div>
      <div>{tag}</div>

      {manifest && repository && (
        <DetailsFrame manifest={manifest} repository={repository} />
      )}
      {/* <PageTitle>Repository: {page_name}</PageTitle>
      {repository && (
        <RepositoryFrame
          repository={repository}
          on_repository_remove={on_repository_remove}
          on_tag_remove={on_tag_remove}
        />
      )}
      {!repository && <div>error</div>} */}
    </div>
  );
}
