import PageTitle from "../../components/page_title";
import { useEffect, useState } from "react";
import TotalReport from "./total_report";
import { ApiSecretListModel, useSecretsService } from "../../services/secrets_service";
import IconSecrets from "../../components/icon_secrets";
import SecretsTable from "./secrets_table";
import { useConfiguration } from "@src/store/configurationContext";

export default function SecretsPage() {
  const { base_url } = useConfiguration();
  const { get_secrets, remove_secret } = useSecretsService(base_url);

  const [secrets, setSecrets] = useState<ApiSecretListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = () => {
    setLoading(true);
    get_secrets()
      .then((secrets: ApiSecretListModel[]) => {
        setSecrets(secrets);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    console.log("page secrets mounted");
    refresh();
  }, []);

  const on_remove = (id: string) => {
    remove_secret(id)
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
        <IconSecrets /> Secrets
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

      <SecretsTable secrets={secrets} on_remove={on_remove} />
      <TotalReport total={secrets.length} />
    </div>
  );
}
