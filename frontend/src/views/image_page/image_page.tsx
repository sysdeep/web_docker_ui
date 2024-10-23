import { useNavigate, useParams } from 'react-router-dom';
import IconImages from '../../components/icon_images';
import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import DetailsFrame from './detailes_frame';
import ContainersFrame from './containers_frame';
import HistoryFrame from './history_frame';
import ImagesService, { ApiFullImageModel } from '../../services/images_service';
import { route } from '@src/routes';
import { useConfiguration } from '@src/store/configuration';
import ButtonRefresh from '@src/components/button_refresh';

export default function ImagePage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const { configuration } = useConfiguration();

  const images_service = useMemo(() => {
    return new ImagesService(configuration.base_url);
  }, []);

  const [image, setImage] = useState<ApiFullImageModel | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = () => {
    setLoading(false);
    images_service
      .get_image(id)
      .then(setImage)
      .catch(console.log)
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    console.log('page image mounted');
    refresh();
  }, []);

  const on_remove = () => {
    images_service
      .remove_image(id)
      .then(() => {
        navigate(route.images);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const body = () => {
    if (image) {
      return (
        <div>
          <DetailsFrame image={image} on_remove={on_remove} />
          <ContainersFrame containers={image.containers} />
          <HistoryFrame image={image} />
        </div>
      );
    }

    return <p>no image</p>;
  };

  const [, image_hash] = id.split(':');
  const page_title = image_hash.slice(0, 12);

  return (
    <div>
      <PageTitle>
        <IconImages />
        &nbsp; Image: {page_title}
      </PageTitle>

      <div>
        <ButtonRefresh on_refresh={refresh} loading={loading} />
      </div>

      {body()}
    </div>
  );
}
