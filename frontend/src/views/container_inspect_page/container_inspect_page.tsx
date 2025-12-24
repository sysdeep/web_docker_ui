import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import PageTitle from "../../components/page_title";
import { useContainersService } from "../../services/containers_service";
import IconContainers from "../../components/icon_containers";
import { useConfiguration } from "@src/store/configurationContext";

export default function ContainerInspectPage() {
  const { id } = useParams();
  const { base_url } = useConfiguration();
  const { get_container_inspect } = useContainersService(base_url);

  const [container, setContainer] = useState<any | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = (uid: string) => {
    setLoading(true);
    get_container_inspect(uid)
      .then((data) => {
        setContainer(data);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    if (id) {
      refresh(id);
    }
  }, []);

  if (!id) {
    return <div>no id!!!</div>;
  }

  return (
    <div>
      <PageTitle onRefresh={() => refresh(id)} isRefresh={loading}>
        <IconContainers /> Container inspect
      </PageTitle>

      <pre>{JSON.stringify(container, null, 4)}</pre>
    </div>
  );
}
