import { useNavigate, useParams } from "react-router-dom";
import PageTitle from "../../components/page_title";
import { useEffect, useMemo, useState } from "react";
import DetailsFrame from "./detailes_frame";
import VolumesService, { ApiFullVolumeModel } from "../../services/volumes_service";
import IconVolumes from "../../components/icon_volumes";
import { route } from "@src/routes";
import { useConfiguration } from "@src/store/configuration";
import ContainersFrame from "./containers_frame";
import ButtonRemove from "@src/components/button_remove";

export default function VolumePage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const { configuration } = useConfiguration();

  const volume_service = useMemo(() => {
    return new VolumesService(configuration.base_url);
  }, []);

  const [volume, setVolume] = useState<ApiFullVolumeModel | null>(null);

  const refresh = (uid: string) => {
    volume_service
      .get_volume(uid)
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
      volume_service
        .remove_volume(id)
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
