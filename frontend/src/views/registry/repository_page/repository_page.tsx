import PageTitle from "@src/components/page_title";
import { RegistryService, RepositoryModel } from "@src/services/registry_service";
import { useEffect, useMemo, useState } from "react";
import { Outlet, useParams } from "react-router-dom";
import RepositoryFrame from "./repository_frame";
import RepositoryNavFrame from "../components/reposytory_nav_frame";
import IconRegistry from "@src/components/icon_registry";
import { useConfiguration } from "@src/store/configurationContext";

export default function RepositoryPage() {
  const { id } = useParams();
  // const navigate = useNavigate();
  const { base_url } = useConfiguration();

  const registry_service = useMemo(() => {
    return new RegistryService(base_url);
  }, []);

  const [repository, setRepository] = useState<RepositoryModel | null>();
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = (uid: string) => {
    setLoading(true);
    registry_service
      .get_repository(uid)
      .then((repo) => {
        setRepository(repo);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    console.log("page repository mounted");
    if (id) {
      refresh(id);
    }
  }, []);

  const on_repository_remove = () => {
    console.log("repo remove page");
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
    console.log("tag remove page: " + tag);
    if (repository && id) {
      registry_service
        .remove_tag(repository.id, tag)
        .then(() => {
          refresh(id);
        })
        .catch((err) => {
          console.log(err);
        });
    }
  };

  // const page_name = repository ? repository.name : id;

  if (!id) {
    return <div>no id!!!</div>;
  }

  return (
    <div>
      <PageTitle onRefresh={() => refresh(id)} isRefresh={loading}>
        <IconRegistry /> Repository info
      </PageTitle>

      {repository && <RepositoryNavFrame repository={repository} />}

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
