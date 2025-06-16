import { useNavigate, useParams } from "react-router-dom";
import IconImages from "../../components/icon_images";
import PageTitle from "../../components/page_title";
import { useEffect, useState } from "react";
import DetailsFrame from "./detailes_frame";
import ContainersFrame from "./containers_frame";
import HistoryFrame from "./history_frame";
import { ApiFullImageModel, useImagesService } from "../../services/images_service";
import { route } from "@src/routes";
import { useConfiguration } from "@src/store/configurationContext";

export default function ImagePage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const { base_url } = useConfiguration();
  const { get_image, remove_image } = useImagesService(base_url);

  const [image, setImage] = useState<ApiFullImageModel | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = (uid: string) => {
    setLoading(false);
    get_image(uid)
      .then(setImage)
      .catch(console.log)
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    console.log("page image mounted");
    if (id) {
      refresh(id);
    }
  }, []);

  const on_remove = (uid: string) => {
    remove_image(uid)
      .then(() => {
        navigate(route.images);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  if (!id) {
    return <div>no id!!!</div>;
  }

  const body = () => {
    if (id && image) {
      return (
        <div>
          <DetailsFrame image={image} on_remove={() => on_remove(id)} />
          <ContainersFrame containers={image.containers} />
          <HistoryFrame image={image} />
        </div>
      );
    }

    return <p>no image</p>;
  };

  const [, image_hash] = id.split(":");
  const page_title = image_hash.slice(0, 12);

  return (
    <div>
      <PageTitle onRefresh={() => refresh(id)} isRefresh={loading}>
        <IconImages />
        &nbsp; Image: {page_title}
      </PageTitle>

      {body()}
    </div>
  );
}
