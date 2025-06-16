import { useNavigate, useParams } from "react-router-dom";
import PageTitle from "../../components/page_title";
import { useEffect, useState } from "react";
import DetailsFrame from "./detailes_frame";
import { ApiFullVolumeModel, useVolumesService } from "../../services/volumes_service";
import IconVolumes from "../../components/icon_volumes";
import { route } from "@src/routes";
import ContainersFrame from "./containers_frame";
import ButtonRemove from "@src/components/button_remove";
import { useConfiguration } from "@src/store/configurationContext";

export default function VolumePage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const { base_url } = useConfiguration();
  const { get_volume, remove_volume } = useVolumesService(base_url);

  const [volume, setVolume] = useState<ApiFullVolumeModel | null>(null);

  const refresh = (uid: string) => {
    get_volume(uid)
      .then((volume) => {
        setVolume(volume);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log("page volume mounted");
    if (id) {
      refresh(id);
    }
  }, []);

  const on_remove = () => {
    if (id) {
      remove_volume(id)
        .then(() => {
          navigate(route.volumes);
        })
        .catch((err) => {
          console.log(err);
        });
    }
  };

  if (!id) {
    return <div>no id!!!</div>;
  }

  const body = () => {
    if (volume) {
      return (
        <div>
          {/* actions */}
          <div className='pull-right'>{volume.containers.length === 0 && <ButtonRemove on_remove={on_remove} />}</div>

          <DetailsFrame volume={volume} />
          <ContainersFrame containers={volume.containers} />
        </div>
      );
    }

    return <p>no volume</p>;
  };

  return (
    <div>
      <PageTitle onRefresh={() => refresh(id)}>
        <IconVolumes />
        &nbsp; Volume: {id}
      </PageTitle>

      {body()}
    </div>
  );
}
