import { useEffect, useMemo, useState } from "react";
import ClientFrame from "./client_frame";
import ServerFrame from "./server_frame";
import SwarmFrame from "./swarm_frame";
import InfoService, { ApiInfoModel } from "../../services/info_service";
import { useConfiguration } from "@src/store/configuration";

export default function HomePage() {
  const { configuration } = useConfiguration();
  const info_service = useMemo(() => {
    return new InfoService(configuration.base_url);
  }, []);

  const [info, setInfo] = useState<ApiInfoModel | null>(null);
  const [error, setError] = useState<string | null>(null);

  const refresh = () => {
    info_service
      .get_info()
      .then((info) => {
        setInfo(info);
      })
      .catch((err) => {
        setError(String(err));
      });
  };

  useEffect(() => {
    refresh();
  }, []);

  const draw_body = () => {
    if (error) {
      return <div>{error}</div>;
    }

    if (info) {
      return (
        <div className='row'>
          <div className='col'>
            <ClientFrame info={info} />
          </div>
          <div className='col'>
            <ServerFrame info={info} />
            <SwarmFrame info={info} />
          </div>
        </div>
      );
    } else {
      return <p>no data</p>;
    }
  };

  return <div>{draw_body()}</div>;
}
