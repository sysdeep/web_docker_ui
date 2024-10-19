import PageTitle from '@src/components/page_title';
import { join_url, route } from '@src/routes';
import {
  RegistryAction,
  RegistryService,
  RepositoryListModel,
} from '@src/services/registry_service';
import { useConfiguration } from '@src/store/configuration';
import React, { useEffect, useMemo, useState } from 'react';
import { Link } from 'react-router-dom';
import ActionsBar from './actions_bar';

export default function RepositoriesPage() {
  const { configuration } = useConfiguration();

  const registry_service = useMemo(() => {
    return new RegistryService(configuration.base_url);
  }, []);

  const [repositories, setRepositories] = useState<RepositoryListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = () => {
    setLoading(true);
    registry_service
      .get_repositories()
      .then((repos) => {
        setRepositories(repos);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    console.log('page repositories mounted');
    refresh();
  }, []);

  // actions ------------------------------------------------------------------
  const on_garbage = () => {
    registry_service.registry_action(RegistryAction.garbage);
  };
  const on_restart = () => {
    registry_service.registry_action(RegistryAction.restart);
  };

  return (
    <div>
      <PageTitle>Repositories</PageTitle>

      <ActionsBar on_garbage={on_garbage} on_restart={on_restart} />

      <div>
        <ul>
          {repositories.map((row, idx) => (
            <li key={idx}>
              <Link to={join_url(route.registry_repository, row.id)}>
                {row.name}
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}
