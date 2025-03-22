import PageTitle from '@src/components/page_title';
import { join_url, route } from '@src/routes';
import { RegistryAction, RegistryService, RepositoryListModel } from '@src/services/registry_service';
import { useConfiguration } from '@src/store/configuration';
import React, { useEffect, useMemo, useState } from 'react';
import { Link } from 'react-router-dom';
import ActionsBar from './actions_bar';
import IconRegistry from '@src/components/icon_registry';
import TreeItem, { TreeItemData } from './tree_item';

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
      <PageTitle>
        <IconRegistry /> Repositories
      </PageTitle>

      <ActionsBar on_garbage={on_garbage} on_restart={on_restart} />

      <div>
        <ul>
          {repositories.map((row, idx) => (
            <li key={idx}>
              <Link to={join_url(route.registry_repository, row.id)}>{row.name}</Link>
            </li>
          ))}
        </ul>
      </div>

      <TreeItem root={to_tree(repositories)} />
    </div>
  );
}

function to_tree(items: RepositoryListModel[]): TreeItemData {
  let root: TreeItemData = {
    id: '',
    name: 'root',
    childrens: [],
  };

  for (let repo of items) {
    const names = repo.name.split('/');
    let root_node = root;
    for (let name_part of names) {
      let node = find_in_tree(root_node, name_part);
      if (!node) {
        let new_node: TreeItemData = { id: repo.id, name: name_part, childrens: [] };
        root_node.childrens.push(new_node);
        root_node = new_node;
      } else {
        root_node = node;
      }
    }
  }

  return root;
}

function find_in_tree(root: TreeItemData, name: string): TreeItemData | null {
  if (root.name == name) {
    return root;
  }

  for (let node of root.childrens) {
    let result = find_in_tree(node, name);
    if (result) {
      return result;
    }
  }

  return null;
}
