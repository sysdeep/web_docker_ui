import { useEffect, useState } from "react";
import ClientFrame from "./client_frame";
import ServerFrame from "./server_frame";
import SwarmFrame from "./swarm_frame";
import { ApiInfoModel, useInfoService } from "../../services/info_service";
import { useConfiguration } from "@src/store/configurationContext";

export default function HomePage() {
  const { base_url } = useConfiguration();
  const { get_info } = useInfoService(base_url);

  const [info, setInfo] = useState<ApiInfoModel | null>(null);
  const [error, setError] = useState<string | null>(null);

  const refresh = () => {
    get_info()
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

  if (error) {
    return <div>{error}</div>;
  }

  if (!info) {
    return <p>no data</p>;
  }

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
}
