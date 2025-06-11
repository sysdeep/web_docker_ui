import PageTitle from "@src/components/page_title";
import { join_url, route } from "@src/routes";
import { RegistryAction, RegistryService, RepositoryListModel } from "@src/services/registry_service";
import { useEffect, useMemo, useState } from "react";
import { Link } from "react-router-dom";
import ActionsBar from "../components/actions_bar";
import IconRegistry from "@src/components/icon_registry";
import { useConfiguration } from "@src/store/configurationContext";

export default function RepositoriesPage() {
  const { base_url } = useConfiguration();

  const registry_service = useMemo(() => {
    return new RegistryService(base_url);
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
        <IconRegistry /> Catalog
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
    </div>
  );
}
