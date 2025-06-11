import { useNavigate, useParams } from "react-router-dom";
import PageTitle from "../../components/page_title";
import { useEffect, useMemo, useState } from "react";
import DetailsFrame from "./detailes_frame";
import { ApiFullNetworkModel, NetworksService } from "../../services/networks_service";
import IconNetworks from "../../components/icon_networks";
import ContainersFrame from "./containers_frame";
import { useConfiguration } from "@src/store/configuration";
import { route } from "@src/routes";

export default function NetworkPage() {
  const { id } = useParams();
  const { configuration } = useConfiguration();
  const navigate = useNavigate();

  const network_service = useMemo(() => {
    return new NetworksService(configuration.base_url);
  }, []);

  const [network, setNetwork] = useState<ApiFullNetworkModel | null>(null);

  const refresh = (uid: string) => {
    network_service
      .get_network(uid)
      .then((network) => {
        setNetwork(network);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log("page network mounted");
    if (id) {
      refresh(id);
    }
  }, []);

  const on_remove = (uid: string) => {
    network_service
      .remove_network(uid)
      .then(() => {
        navigate(route.networks);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  if (!id) {
    return <div>no id!!!</div>;
  }

  const body = () => {
    if (network) {
      return (
        <div>
          <DetailsFrame network={network} on_remove={() => on_remove(id)} />
          <ContainersFrame containers={network.containers} />
        </div>
      );
    }

    return <p>no network</p>;
  };

  const page_title = network ? network.network.name : id;

  return (
    <div>
      <PageTitle onRefresh={() => refresh(id)}>
        <IconNetworks />
        &nbsp; Network: {page_title}
      </PageTitle>

      {body()}
    </div>
  );
}
