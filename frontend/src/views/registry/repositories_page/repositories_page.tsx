import PageTitle from '@src/components/page_title';
import { join_url, route } from '@src/routes';
import {
  RegistryService,
  RepositoryListModel,
} from '@src/services/registry_service';
import { useConfiguration } from '@src/store/configuration';
import React, { useEffect, useMemo, useState } from 'react';
import { Link } from 'react-router-dom';

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

  return (
    <div>
      <PageTitle>Repositories</PageTitle>
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
