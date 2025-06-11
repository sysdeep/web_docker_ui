import { createContext, useContext } from "react";

type Configuration = {
  base_url: string;
  version: string;
  use_registry: boolean;
};

const default_configuration: Configuration = {
  base_url: "/",
  version: "0.0.0",
  use_registry: false,
};

export const ConfigurationContext = createContext<Configuration>(default_configuration);

export function useConfiguration() {
  const configuration = useContext(ConfigurationContext);

  return { ...configuration };
}
