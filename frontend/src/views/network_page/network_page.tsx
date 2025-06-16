import { useNavigate, useParams } from "react-router-dom";
import PageTitle from "../../components/page_title";
import { useEffect, useState } from "react";
import DetailsFrame from "./detailes_frame";
import { ApiFullNetworkModel, useNetworksService } from "../../services/networks_service";
import IconNetworks from "../../components/icon_networks";
import ContainersFrame from "./containers_frame";
import { route } from "@src/routes";
import { useConfiguration } from "@src/store/configurationContext";

export default function NetworkPage() {
  const { id } = useParams();
  const { base_url } = useConfiguration();
  const { get_network, remove_network } = useNetworksService(base_url);
  const navigate = useNavigate();

  const [network, setNetwork] = useState<ApiFullNetworkModel | null>(null);

  const refresh = (uid: string) => {
    get_network(uid)
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
    remove_network(uid)
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
