import { ApiContainerListModel } from "@src/models/api_container_list_model";

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
