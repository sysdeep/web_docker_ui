import PageTitle from '@src/components/page_title';
import { RegistryService, RepositoryModel, TagManifest } from '@src/services/registry_service';
import { useConfiguration } from '@src/store/configuration';
import React, { useEffect, useMemo, useState } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import DetailsFrame from './details_frame';
import RepositoryNavFrame from '../components/reposytory_nav_frame';
import ButtonRemove from '@src/components/button_remove';
import { join_url, route } from '@src/routes';
import IconRegistry from '@src/components/icon_registry';
// import RepositoryFrame from './repository_frame';

export default function RepositoryTagPage() {
  const { id, tag } = useParams();
  const navigate = useNavigate();
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
  }, [tag]);

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

  const on_tag_remove = () => {
    if (id && tag) {
      registry_service
        .remove_tag(id, tag)
        .then(() => {
          navigate(join_url(route.registry_repository, id));
        })
        .catch((err) => {
          console.log(err);
        });
    }
  };

  let page_name = '';
  if (repository && manifest) {
    page_name = `${repository.name}:${manifest.name}`;
  }

  return (
    <div>
      <PageTitle>
        <IconRegistry /> Tag info
      </PageTitle>

      {repository && <RepositoryNavFrame repository={repository} />}

      {manifest && repository && <DetailsFrame manifest={manifest} repository={repository} />}

      {/* actions */}
      <div>
        <ButtonRemove on_remove={on_tag_remove} />
      </div>
    </div>
  );
}
