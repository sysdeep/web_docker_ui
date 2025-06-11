import { ContainerAction } from "@src/services/containers_service";

type ActionsFrameProps = {
  status: string;
  on_action(action: string): void;
};

export default function ActionsFrame({ status, on_action }: ActionsFrameProps) {
  const start_classes = status === "running" || status === "paused" ? "disabled" : "";
  const stop_classes = status === "exited" || status === "paused" ? "disabled" : "";
  const kill_classes = status === "exited" || status === "paused" ? "disabled" : "";
  const restart_classes = status === "exited" || status === "paused" ? "disabled" : "";
  const pause_classes = status === "running" ? "" : "disabled";
  const unpause_classes = status === "paused" ? "" : "disabled";
  const remove_classes = status === "exited" ? "" : "disabled";

  // const on_start_clicked = () => {
  //   console.log('start');
  // };

  return (
    <div>
      <div>
        <button className={"btn btn-secondary mx-2 " + start_classes} onClick={() => on_action(ContainerAction.start)}>
          <i className='bi bi-play-fill'></i> Start
        </button>
        <button className={"btn btn-secondary mx-2 " + stop_classes} onClick={() => on_action(ContainerAction.stop)}>
          <i className='bi bi-stop-fill'></i> Stop
        </button>
        <button className={"btn btn-secondary mx-2 " + kill_classes} onClick={() => on_action(ContainerAction.kill)}>
          <i className='bi bi-x-lg'></i> Kill
        </button>
        <button
          className={"btn btn-secondary mx-2 " + restart_classes}
          onClick={() => on_action(ContainerAction.restart)}
        >
          <i className='bi bi-arrow-repeat'></i> Restart
        </button>
        <button className={"btn btn-secondary mx-2 " + pause_classes} onClick={() => on_action(ContainerAction.pause)}>
          <i className='bi bi-box-arrow-in-down'></i> Pause
        </button>
        <button
          className={"btn btn-secondary mx-2 " + unpause_classes}
          onClick={() => on_action(ContainerAction.resume)}
        >
          <i className='bi bi-box-arrow-down'></i> Resume
        </button>
        <button
          className={"btn btn-secondary mx-2 " + remove_classes}
          onClick={() => on_action(ContainerAction.remove)}
        >
          <i className='bi bi-trash'></i> Remove
        </button>
      </div>
    </div>
  );
}
