import { join_url } from "@src/routes";

export function useSecretsService(base_url: string) {
  const get_secrets = async (): Promise<ApiSecretListModel[]> => {
    const response = await fetch(join_url(base_url, "/api/secrets"));

    const data = (await response.json()) as ApiSecretsListModel;

    return data.secrets || [];
  };

  const get_secret = async (id: string): Promise<ApiFullSecretModel> => {
    const response = await fetch(join_url(base_url, "/api/secrets/" + id));

    const data = (await response.json()) as ApiFullSecretModel;

    return data;
  };

  const remove_secret = async (id: string): Promise<void> => {
    await fetch(join_url(base_url, "/api/secrets/" + id), {
      method: "DELETE",
    });
    return;
  };

  return {
    get_secret,
    get_secrets,
    remove_secret,
  };
}

// list models ----------------------------------------------------------------
export interface ApiSecretListModel {
  id: string;
  name: number;
  created: string;
  updated: string;
}

interface ApiSecretsListModel {
  secrets: ApiSecretListModel[];
  total: number;
}

// Secret model ---------------------------------------------------------------
interface ApiSecretModel {
  id: string;
  name: string;
  created: string;
  updated: string;
}

export interface ApiFullSecretModel {
  secret: ApiSecretModel;
}
