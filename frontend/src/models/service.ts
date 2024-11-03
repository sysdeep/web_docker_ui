export type Service = {
  id: string;
  name: string;
  image: string;
  mode: ServiceMode;
  created_at: string;
  updated_at: string;
  spec: ServiceSpec;
};

type ReplicatedService = {
  replicas: number;
};

export type ServiceMode = {
  replicated: null | ReplicatedService;
  global: null | any; // NOTE: приходит только признак - есть или нет
};

type ServiceSpec = {
  name: string;
  task_template: TaskTemplate;
};

type TaskTemplate = {
  container_spec: ContainerSpec;
};

export type ContainerSpec = {
  image: string;
  args: string[];
  env: string[];
};
