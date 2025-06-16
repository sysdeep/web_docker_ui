import { ApiContainerListModel } from "@src/models/api_container_list_model";
import ImageListModel from "@src/models/image_list_model";
import { join_url } from "@src/routes";

export function useImagesService(base_url: string) {
  async function get_images(): Promise<ImageListModel[]> {
    const response = await fetch(join_url(base_url, "/api/images"));

    const data = (await response.json()) as ApiImagesListModel;

    const images = data.images || [];
    const dataset = images.map((model) => {
      const dmodel: ImageListModel = {
        id: model.id,
        created: model.created,
        tags: model.tags,
        size: model.size,
      };
      return dmodel;
    });

    return dataset;
  }

  async function get_image(id: string): Promise<ApiFullImageModel> {
    const response = (await fetch(join_url(base_url, "/api/images/" + id)).then((data) =>
      data.json(),
    )) as ApiFullImageModel;

    return response;
  }

  async function remove_image(id: string): Promise<void> {
    await fetch(join_url(base_url, "/api/images/" + id), {
      method: "DELETE",
    });

    return;
  }

  return { get_image, get_images, remove_image };
}

// images list ----------------------------------------------------------------
interface ApiImageListModel {
  containers: number;
  created: string;
  id: string;
  tags: string[];
  size: number;
}

export type ApiImagesListModel = {
  images: ApiImageListModel[];
  total: number;
};

// image models ---------------------------------------------------------------
interface ApiImageModel {
  id: string;
  repo_tags: string[];
  parent: string;
  comment: string;
  created: string;
  size: number;
}

interface ApiImageHistoryModel {
  created: string;
  created_by: string;
  id: string;
  size: number;
  tags: string[];
}

export interface ApiFullImageModel {
  image: ApiImageModel;
  history: ApiImageHistoryModel[];
  containers: ApiContainerListModel[];
}
