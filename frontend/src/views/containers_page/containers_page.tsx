import PageTitle from "../../components/page_title";
import { useEffect, useState } from "react";
import ContainersFrame from "./containers_frame";
import { useContainersService } from "../../services/containers_service";
import TotalReport from "./total_report";
import IconContainers from "../../components/icon_containers";
import { ApiContainerListModel } from "@src/models/api_container_list_model";
import { useConfiguration } from "@src/store/configurationContext";

export default function ContainersPage() {
  const { base_url } = useConfiguration();
  const { get_containers } = useContainersService(base_url);

  const [containers, setContainers] = useState<ApiContainerListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  // const [filter, setFilter] = useState<FilterModel>({ dates: [] });

  const refresh = () => {
    setLoading(true);
    get_containers()
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
    console.log("containers page mounted");
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
      <PageTitle onRefresh={refresh} isRefresh={loading}>
        <IconContainers /> Containers
      </PageTitle>

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
