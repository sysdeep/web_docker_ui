import React, { useState, useEffect, useMemo } from 'react';
import PageTitle from '../../components/page_title';
import ImagesTable from './images_table';
import ImageListModel from '../../models/image_list_model';
import ImagesService from '../../services/images_service';
import FilterPanel from './filter_panel';
import FilterModel from './filter_model';
import TotalReport from './total_report';
import IconImages from '../../components/icon_images';
import { useConfiguration } from '@src/store/configuration';
import ButtonRefresh from '@src/components/button_refresh';

export default function ImagesPage() {
  const { configuration } = useConfiguration();
  const images_service = useMemo(() => {
    return new ImagesService(configuration.base_url);
  }, []);

  const [images, setImages] = useState<ImageListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [filter, setFilter] = useState<FilterModel>({ dates: [] });

  const refresh = () => {
    setLoading(true);
    images_service
      .get_images()
      .then((images: ImageListModel[]) => {
        setImages(images);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    console.log('page images mounted');
    refresh();
  }, []);

  const remove_image = (id: string) => {
    images_service
      .remove_image(id)
      .then(() => {
        refresh();
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const on_date = (date: string) => {
    let result = [];
    if (filter.dates.includes(date)) {
      result = filter.dates.filter((d) => d !== date);
    } else {
      result = [...filter.dates, date];
    }
    setFilter({ ...filter, dates: result });
  };

  return (
    <div>
      <PageTitle>
        <IconImages /> Images
      </PageTitle>
      <FilterPanel filter={filter} on_date={on_date} />
      <div>
        <ButtonRefresh on_refresh={refresh} loading={loading} />
      </div>
      <ImagesTable
        images={images}
        filter={filter}
        on_remove={remove_image}
        on_date={on_date}
      />
      <TotalReport total={images.length} />
    </div>
  );
}
