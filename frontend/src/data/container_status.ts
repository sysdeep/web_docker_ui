export enum ContainerStatus {
  exited,
  running,
  paused,
  unknown,
}

export function makeContainerStatus(v: string): ContainerStatus {
  switch (v) {
    case "exited":
      return ContainerStatus.exited;
    case "running":
      return ContainerStatus.running;
    case "paused":
      return ContainerStatus.paused;

    default:
      return ContainerStatus.unknown;
  }
}

// const container_statuses_list = [ContainerStatus.exited];

export const container_status_map = {
  [ContainerStatus.exited]: "exited",
  [ContainerStatus.paused]: "paused",
  [ContainerStatus.running]: "running",
  [ContainerStatus.unknown]: "unknown",
};
