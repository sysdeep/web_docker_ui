import PageTitle from '@src/components/page_title';
import {
  RegistryService,
  RepositoryModel,
} from '@src/services/registry_service';
import { useConfiguration } from '@src/store/configuration';
import React, { useEffect, useMemo, useState } from 'react';
import { Link, Outlet, useNavigate, useParams } from 'react-router-dom';
import RepositoryFrame from './repository_frame';

export default function RepositoryPage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const { configuration } = useConfiguration();

  const registry_service = useMemo(() => {
    return new RegistryService(configuration.base_url);
  }, []);

  const [repository, setRepository] = useState<RepositoryModel | null>();
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = () => {
    setLoading(true);
    registry_service
      .get_repository(id)
      .then((repo) => {
        setRepository(repo);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    console.log('page repository mounted');
    refresh();
  }, []);

  const on_repository_remove = () => {
    console.log('repo remove page');
    // if (secret) {
    //   secret_service
    //     .remove_secret(id)
    //     .then(() => {
    //       console.log('remove ok');
    //       navigate(route.secrets);
    //     })
    //     .catch((err) => {
    //       console.log(err);
    //     });
    // }
  };

  const on_tag_remove = (tag: string) => {
    console.log('tag remove page: ' + tag);
    // if (secret) {
    //   secret_service
    //     .remove_secret(id)
    //     .then(() => {
    //       console.log('remove ok');
    //       navigate(route.secrets);
    //     })
    //     .catch((err) => {
    //       console.log(err);
    //     });
    // }
  };

  const page_name = repository ? repository.name : id;

  return (
    <div>
      <PageTitle>Repository: {page_name}</PageTitle>
      {repository && (
        <RepositoryFrame
          repository={repository}
          on_repository_remove={on_repository_remove}
          on_tag_remove={on_tag_remove}
        />
      )}
      {!repository && <div>error</div>}

      <div>
        <Outlet />
      </div>
    </div>
  );
}
