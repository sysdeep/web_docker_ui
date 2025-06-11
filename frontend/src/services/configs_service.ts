import { join_url } from "@src/routes";

export function useConfigsServices(base_url: string) {
  async function get_configs(): Promise<ApiConfigListModel[]> {
    const response = await fetch(join_url(base_url, "/api/configs"));

    const data = (await response.json()) as ApiConfigsListModel;

    return data.configs || [];
  }

  async function get_config(id: string): Promise<ApiFullConfigModel> {
    const response = await fetch(join_url(base_url, "/api/configs/" + id));

    const data = (await response.json()) as ApiFullConfigModel;

    return data;
  }

  async function remove_config(id: string): Promise<void> {
    await fetch(join_url(base_url, "/api/configs/" + id), {
      method: "DELETE",
    });

    return;
  }

  return {
    get_config,
    get_configs,
    remove_config,
  };
}

// list models ----------------------------------------------------------------
export interface ApiConfigListModel {
  id: string;
  name: number;
  created: string;
  updated: string;
}

interface ApiConfigsListModel {
  configs: ApiConfigListModel[];
  total: number;
}

// config model ---------------------------------------------------------------
interface ApiConfigModel {
  id: string;
  name: string;
  created: string;
  updated: string;
  data_text: string;
}

export interface ApiFullConfigModel {
  config: ApiConfigModel;
}
