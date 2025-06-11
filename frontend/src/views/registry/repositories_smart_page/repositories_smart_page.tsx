import PageTitle from "@src/components/page_title";
import { join_url, route } from "@src/routes";
import { RegistryAction, RegistryService, RepositoryModel } from "@src/services/registry_service";
import { useConfiguration } from "@src/store/configuration";
import { useEffect, useMemo, useState } from "react";
import { Link } from "react-router-dom";
import ActionsBar from "../components/actions_bar";
import IconRegistry from "@src/components/icon_registry";

export default function RepositoriesSmartPage() {
  const { configuration } = useConfiguration();

  const registry_service = useMemo(() => {
    return new RegistryService(configuration.base_url);
  }, []);

  const [repositories, setRepositories] = useState<RepositoryModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [showEmpty, setShowEmpty] = useState<boolean>(false);

  const refresh = () => {
    setLoading(true);
    registry_service
      .get_repositories_smart()
      .then((repos) => {
        setRepositories(repos);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    console.log("page repositories mounted");
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
      <PageTitle onRefresh={refresh} isRefresh={loading}>
        <IconRegistry /> Repositories
      </PageTitle>

      <ActionsBar on_garbage={on_garbage} on_restart={on_restart} />

      {/* page settings */}
      <div>
        <div className='form-check'>
          <input
            className='form-check-input'
            type='checkbox'
            onChange={() => setShowEmpty(!showEmpty)}
            value={showEmpty ? "1" : "0"}
            id='flexCheckDefault'
          />
          <label className='form-check-label' htmlFor='flexCheckDefault'>
            отображать пустые
          </label>
        </div>
      </div>

      <hr />

      <div>
        <ul>
          {repositories
            .filter((row) => {
              if (showEmpty) return true;
              return row.tags.length > 0;
            })
            .map((row, idx) => (
              <li key={idx}>
                <div>
                  <Link to={join_url(route.registry_repository, row.id)}>{row.name}</Link>
                </div>
                <ul>
                  {row.tags.map((tag, idt) => (
                    <li key={idt}>{tag}</li>
                  ))}
                </ul>
              </li>
            ))}
        </ul>
      </div>

      {/* NOTE: not used here */}
      {/* <TreeItem root={to_tree(repositories)} /> */}
    </div>
  );
}

// function to_tree(items: RepositoryListModel[]): TreeItemData {
//   let root: TreeItemData = {
//     id: "",
//     name: "root",
//     childrens: [],
//   };

//   for (let repo of items) {
//     const names = repo.name.split("/");
//     let root_node = root;
//     for (let name_part of names) {
//       let node = find_in_tree(root_node, name_part);
//       if (!node) {
//         let new_node: TreeItemData = { id: repo.id, name: name_part, childrens: [] };
//         root_node.childrens.push(new_node);
//         root_node = new_node;
//       } else {
//         root_node = node;
//       }
//     }
//   }

//   return root;
// }

// function find_in_tree(root: TreeItemData, name: string): TreeItemData | null {
//   if (root.name == name) {
//     return root;
//   }

//   for (let node of root.childrens) {
//     let result = find_in_tree(node, name);
//     if (result) {
//       return result;
//     }
//   }

//   return null;
// }
