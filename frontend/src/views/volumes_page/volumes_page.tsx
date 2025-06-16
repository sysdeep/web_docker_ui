import { ApiVolumeListModel, useVolumesService } from "../../services/volumes_service";
import IconVolumes from "../../components/icon_volumes";
import PageTitle from "../../components/page_title";
import { useEffect, useState } from "react";
import VolumesTable from "./volumes_table";
import TotalReport from "./total_report";
import { useConfiguration } from "@src/store/configurationContext";

export default function VolumesPage() {
  const { base_url } = useConfiguration();
  const { get_volumes, remove_volume } = useVolumesService(base_url);

  const [volumes, setVolumes] = useState<ApiVolumeListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = () => {
    setLoading(true);
    get_volumes()
      .then((volumes: ApiVolumeListModel[]) => {
        setVolumes(volumes);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    console.log("page volumes mounted");
    refresh();
  }, []);

  const on_remove = (name: string) => {
    remove_volume(name)
      .then(() => {
        refresh();
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <div>
      <PageTitle onRefresh={refresh} isRefresh={loading}>
        <IconVolumes /> Volumes
      </PageTitle>

      {/* // TODO //{' '} */}
      {/* <div>
        //{' '}
        <a href='/volumes/actions/prune' class='button error'>
          // <i class='fa fa-trash-o' aria-hidden='true'></i>
          // Prune //{' '}
        </a>
        //{' '}
      </div> */}

      <VolumesTable volumes={volumes} on_remove={on_remove} />
      <TotalReport total={volumes.length} />
    </div>
  );
}
