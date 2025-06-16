import { join_url } from "@src/routes";

export function useNetworksService(base_url: string) {
  const get_networks = async (): Promise<ApiNetworkListModel[]> => {
    const response = await fetch(join_url(base_url, "/api/networks"));

    const data = (await response.json()) as ApiNetworksListModel;

    return data.networks || [];
  };

  const get_network = async (id: string): Promise<ApiFullNetworkModel> => {
    const response = await fetch(join_url(base_url, "/api/networks/" + id));

    const data = (await response.json()) as ApiFullNetworkModel;

    return data;
  };

  const remove_network = async (id: string): Promise<void> => {
    await fetch(join_url(base_url, "/api/networks/" + id), {
      method: "DELETE",
    });

    return;
  };

  return {
    get_network,
    get_networks,
    remove_network,
  };
}

// list models ----------------------------------------------------------------
export interface ApiNetworkListModel {
  id: string;
  name: number;
  created: string;
  driver: string;
}

interface ApiNetworksListModel {
  networks: ApiNetworkListModel[];
  total: number;
}
// network model --------------------------------------------------------------
interface ApiNetworkModel {
  name: string;
  id: string;
  created: string;
  scope: string;
  driver: string;
  internal: boolean;
  attachable: boolean;
  ingress: boolean;
}

export interface ApiNetworkContainerModel {
  id: string;
  name: string;
  endpoint_id: string;
  mac_address: string;
  ip_v4_address: string;
  ip_v6_address: string;
}

export interface ApiFullNetworkModel {
  network: ApiNetworkModel;
  containers: ApiNetworkContainerModel[];
}
