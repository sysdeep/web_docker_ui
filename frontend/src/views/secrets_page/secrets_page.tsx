import PageTitle from "../../components/page_title";
import { useEffect, useMemo, useState } from "react";
import TotalReport from "./total_report";
import { ApiSecretListModel, SecretsService } from "../../services/secrets_service";
import IconSecrets from "../../components/icon_secrets";
import SecretsTable from "./secrets_table";
import { useConfiguration } from "@src/store/configurationContext";

export default function SecretsPage() {
  const { base_url } = useConfiguration();
  const secrets_service = useMemo(() => {
    return new SecretsService(base_url);
  }, []);

  const [secrets, setSecrets] = useState<ApiSecretListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const refresh = () => {
    setLoading(true);
    secrets_service
      .get_secrets()
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
    secrets_service
      .remove_secret(id)
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
