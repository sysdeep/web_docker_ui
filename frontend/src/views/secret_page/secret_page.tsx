import { useNavigate, useParams } from "react-router-dom";
import PageTitle from "../../components/page_title";
import { useEffect, useState } from "react";
import DetailsFrame from "./detailes_frame";
import { ApiFullSecretModel, useSecretsService } from "../../services/secrets_service";
import IconSecrets from "../../components/icon_secrets";
import { route } from "@src/routes";
import { useConfiguration } from "@src/store/configurationContext";

export default function SecretPage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const { base_url } = useConfiguration();
  const { get_secret, remove_secret } = useSecretsService(base_url);

  const [secret, setSecret] = useState<ApiFullSecretModel | null>(null);

  const refresh = (uid: string) => {
    get_secret(uid)
      .then((secret) => {
        setSecret(secret);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    if (id) {
      refresh(id);
    }
  }, []);

  const on_remove = () => {
    if (secret && id) {
      remove_secret(id)
        .then(() => {
          console.log("remove ok");
          navigate(route.secrets);
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
    if (secret) {
      return (
        <div>
          <DetailsFrame secret={secret} on_remove={on_remove} />
        </div>
      );
    }

    return <p>no secret</p>;
  };

  const page_name = secret ? secret.secret.name : id;

  return (
    <div>
      <PageTitle onRefresh={() => refresh(id)}>
        <IconSecrets />
        &nbsp; Secret: {page_name}
      </PageTitle>

      {body()}
    </div>
  );
}
