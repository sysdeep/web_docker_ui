import { join_url } from "@src/routes";
import ImageListModel from "../models/image_list_model";
import { ApiFullImageModel, ApiImagesListModel } from "./images_service";

export default function useImagesService(base_url: string) {
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
