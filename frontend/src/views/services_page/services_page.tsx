import PageTitle from "../../components/page_title";
import { useEffect, useState } from "react";
import TotalReport from "./total_report";
import IconServices from "@src/components/icon_services";
import { useServicesService } from "@src/services/services_service";
import ServicesTable from "./services_table";
import { Service } from "@src/models/service";
import { useConfiguration } from "@src/store/configurationContext";

export default function ServicesPage() {
  const { base_url } = useConfiguration();
  const { get_services } = useServicesService(base_url);

  const [services, setServices] = useState<Service[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = () => {
    setLoading(true);
    get_services()
      .then(setServices)
      .catch(console.log)
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    console.log("page services mounted");
    refresh();
  }, []);

  const on_remove = (id: string) => {
    console.log("TODO: service remove: ", id);
    // services_service
    //   .remove_volume(name)
    //   .then(() => {
    //     refresh();
    //   })
    //   .catch((err) => {
    //     console.log(err);
    //   });
  };

  return (
    <div>
      <PageTitle onRefresh={refresh} isRefresh={loading}>
        <IconServices /> Services
      </PageTitle>

      <ServicesTable services={services} on_remove={on_remove} />
      <TotalReport total={services.length} />
    </div>
  );
}
