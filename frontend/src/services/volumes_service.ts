import { ApiContainerListModel } from "@src/models/api_container_list_model";
import { join_url } from "@src/routes";

export function useVolumesService(base_url: string) {
  const get_volumes = async (): Promise<ApiVolumeListModel[]> => {
    const response = await fetch(join_url(base_url, "/api/volumes"));

    const data = (await response.json()) as ApiVolumesListModel;

    return data.volumes || [];
  };

  const get_volume = async (id: string): Promise<ApiFullVolumeModel> => {
    const response = await fetch(join_url(base_url, "/api/volumes/" + id));

    const data = (await response.json()) as ApiFullVolumeModel;

    return data;
  };

  const remove_volume = async (name: string): Promise<void> => {
    await fetch(join_url(base_url, "/api/volumes/" + name), {
      method: "DELETE",
    });
    return;
  };

  return {
    get_volume,
    get_volumes,
    remove_volume,
  };
}

// list models ----------------------------------------------------------------
export interface ApiVolumeListModel {
  created: string;
  name: string;
  stack_name: string;
  mount_point: string;
  driver: string;
  used: boolean;
}

interface ApiVolumesListModel {
  volumes: ApiVolumeListModel[];
  total: number;
}

// volume model ---------------------------------------------------------------
export interface ApiFullVolumeModel {
  volume: ApiVolumeModel;
  containers: ApiContainerListModel[];
}

interface ApiVolumeModel {
  name: string;
  created: string;
  driver: string;
  mount_point: string;
}
