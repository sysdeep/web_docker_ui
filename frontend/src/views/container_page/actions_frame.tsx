import { ContainerAction } from '@src/services/containers_service';
import React from 'react';

type ActionsFrameProps = {
  status: string;
  on_action(action: string): void;
};

export default function ActionsFrame({ status, on_action }: ActionsFrameProps) {
  const start_classes = status === 'running' ? 'disabled' : '';
  const stop_classes = status === 'exited' ? 'disabled' : '';

  // const on_start_clicked = () => {
  //   console.log('start');
  // };

  return (
    <div className='box'>
      <h2>Actions</h2>
      <div>
        <button
          className={'btn btn-secondary mx-2 ' + start_classes}
          onClick={() => on_action(ContainerAction.start)}
        >
          <i className='bi bi-play-fill'></i> Start
        </button>
        <button
          className={'btn btn-secondary ' + stop_classes}
          onClick={() => on_action(ContainerAction.stop)}
        >
          <i className='bi bi-stop-fill'></i> Stop
        </button>
        <button
          className='btn btn-secondary mx-2'
          onClick={() => on_action(ContainerAction.kill)}
        >
          Kill TODO
        </button>
        <button
          className='btn btn-secondary mx-2'
          onClick={() => on_action(ContainerAction.restart)}
        >
          Restart TODO
        </button>
        <button
          className='btn btn-secondary mx-2'
          onClick={() => on_action(ContainerAction.pause)}
        >
          Pause TODO
        </button>
        <button
          className='btn btn-secondary mx-2'
          onClick={() => on_action(ContainerAction.resume)}
        >
          Resume TODO
        </button>
        <button
          className='btn btn-secondary mx-2'
          onClick={() => on_action(ContainerAction.remove)}
        >
          Remove TODO
        </button>
      </div>
    </div>
  );
}
