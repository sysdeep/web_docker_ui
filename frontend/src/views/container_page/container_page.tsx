import { useEffect, useMemo, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

import StatusFrame from "./status_frame";
import PageTitle from "../../components/page_title";
import ContainersService, { ApiContainerResponseModel } from "../../services/containers_service";
import DetailsFrame from "./details_frame";
import VolumesFrame from "./volumes_frame";
import NetworksFrame from "./networks_frame";
import IconContainers from "../../components/icon_containers";
import TopFrame from "./top_frame";
import { route } from "@src/routes";
import { strip_container_name } from "@src/utils/containers";
import { useConfiguration } from "@src/store/configurationContext";

export default function ContainerPage() {
  const { id } = useParams();
  const { base_url } = useConfiguration();
  const navigate = useNavigate();

  const containers_service = useMemo(() => {
    return new ContainersService(base_url);
  }, []);

  const [container, setContainer] = useState<ApiContainerResponseModel | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = (uid: string) => {
    setLoading(true);
    containers_service
      .get_container(uid)
      .then((data) => {
        setContainer(data);
        // refresh_top();
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    console.log("container page mounted");
    if (id) {
      refresh(id);
    }
  }, []);

  const on_action = (action: string) => {
    if (id) {
      containers_service.container_action(id, action).then(() => {
        if (action === "remove") {
          navigate(route.containers);
        } else {
          refresh(id);
        }
      });
    }
  };

  if (!id) {
    return <div>no id!!!</div>;
  }

  const main_body = () => {
    if (container === null) {
      return <div>not loaded</div>;
    }

    return (
      <>
        <StatusFrame container={container} on_action={on_action} />

        <div>
          <ul>
            <li>
              <a href='/container/{container.container.id}/logs'>Logs TODO</a>
            </li>
            <li>
              <a href='/container/{container.container.id}/inspect'>Inspect TODO</a>
            </li>
            <li>
              <a href='/container/{container.container.id}/stats'>Stats TODO</a>
            </li>
            <li>
              <a href='/container/{container.container.id}/console'>Console TODO</a>
            </li>
            <li>
              <a href='/container/{container.container.id}/attach'>Attach TODO</a>
            </li>
          </ul>
        </div>

        <TopFrame container={container} containers_service={containers_service} />
        {/* TODO: not ready */}
        {/* <StatsFrame container={container} containers_service={containers_service} /> */}

        <DetailsFrame container={container} />
        <NetworksFrame container={container} />
        <VolumesFrame container={container} />
      </>
    );
  };

  return (
    <div>
      <PageTitle onRefresh={() => refresh(id)} isRefresh={loading}>
        <IconContainers /> Container: {container ? strip_container_name(container.container.name) : "loading"}
      </PageTitle>

      {main_body()}
    </div>
  );
}
