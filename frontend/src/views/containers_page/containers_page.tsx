import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import ContainersFrame from './containers_frame';
import ContainersService from '../../services/containers_service';
import TotalReport from './total_report';
import IconContainers from '../../components/icon_containers';
import { useConfiguration } from '@src/store/configuration';
import { ApiContainerListModel } from '@src/models/api_container_list_model';
import ButtonRefresh from '@src/components/button_refresh';

export default function ContainersPage() {
  const { configuration } = useConfiguration();
  const containers_service = useMemo(() => {
    return new ContainersService(configuration.base_url);
  }, []);

  const [containers, setContainers] = useState<ApiContainerListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  // const [filter, setFilter] = useState<FilterModel>({ dates: [] });

  const refresh = () => {
    setLoading(true);
    containers_service
      .get_containers()
      .then((containers) => {
        setContainers(containers);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    console.log('containers page mounted');
    refresh();
  }, []);

  // const remove_image = (id: string) => {
  //   console.log('remove', id);

  //   containers_service
  //     .remove_image(id)
  //     .then(() => {
  //       refresh();
  //     })
  //     .catch((err) => {
  //       console.log(err);
  //     });
  // };

  // const on_date = (date: string) => {
  //   let result = [];
  //   if (filter.dates.includes(date)) {
  //     result = filter.dates.filter((d) => d !== date);
  //   } else {
  //     result = [...filter.dates, date];
  //   }
  //   setFilter({ ...filter, dates: result });
  // };

  return (
    <div>
      <PageTitle>
        <IconContainers /> Containers
      </PageTitle>
      <div>
        <ButtonRefresh on_refresh={refresh} loading={loading} />
      </div>
      <ContainersFrame containers={containers} />
      <TotalReport total={containers.length} />
      {/* <FilterPanel filter={filter} on_date={on_date} />
      <div>
        <span>loading: {loading}</span>
        <button className='button' onClick={() => refresh()}>
          Refresh
        </button>
      </div>
      <ImagesTable
        images={images}
        filter={filter}
        on_remove={remove_image}
        on_date={on_date}
      />
       */}
    </div>
  );
}
