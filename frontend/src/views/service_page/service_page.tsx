import { useParams } from "react-router-dom";
import PageTitle from "../../components/page_title";
import { useEffect, useState } from "react";
import ButtonRefresh from "@src/components/button_refresh";
import IconServices from "@src/components/icon_services";
import { Service } from "@src/models/service";
import { useConfiguration } from "@src/store/configurationContext";
import { useServicesService } from "@src/services/services_service";

export default function ServicePage() {
  const { id } = useParams();
  // const navigate = useNavigate();
  const { base_url } = useConfiguration();
  const { get_service } = useServicesService(base_url);

  const [service, setService] = useState<Service | null>(null);

  const refresh = (uid: string) => {
    get_service(uid).then(setService).catch(console.log);
  };

  // useEffect(() => {
  //   console.log(service);
  // }, [service]);

  useEffect(() => {
    console.log("page service mounted");
    if (id) {
      refresh(id);
    }
  }, []);

  // const on_remove = () => {
  //   console.log("TODO: remove service: ", id);
  //   // services_service
  //   //   .remove_volume(id)
  //   //   .then(() => {
  //   //     navigate(route.volumes);
  //   //   })
  //   //   .catch((err) => {
  //   //     console.log(err);
  //   //   });
  // };

  if (!id) {
    return <div>no id!!!</div>;
  }

  const body = () => {
    if (service) {
      return (
        <div>
          {/* actions */}
          <div className='pull-right'>
            {/* TODO */}
            {/* {service.containers.length === 0 && <ButtonRemove on_remove={on_remove} />} */}
            <ButtonRefresh on_refresh={() => refresh(id)} />
          </div>

          {/* TODO */}
          {/* <DetailsFrame volume={service} /> */}
          {/* <ContainersFrame containers={service.containers} /> */}

          {/* TODO: ошибка какая то */}
          {/* <ContainerSpec container={service.spec.task_template.container_spec} /> */}
        </div>
      );
    }

    return <p>no service</p>;
  };

  return (
    <div>
      <PageTitle onRefresh={() => refresh(id)}>
        <IconServices />
        &nbsp; Service: {id}
      </PageTitle>

      {body()}
    </div>
  );
}

// function ServiceBody(){

// }
