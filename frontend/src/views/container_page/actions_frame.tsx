import React from 'react';

interface ActionsFrameProps {
  id: string;
}
export default function ActionsFrame({ id }: ActionsFrameProps) {
  return (
    <div className='box'>
      <h2>Actions</h2>
      <div>
        <a
          className='btn btn-secondary mx-2'
          href='/container/{id}/action/start'
        >
          Start TODO
        </a>
        <a className='btn btn-secondary' href='/container/{id}/action/stop'>
          Stop TODO
        </a>
        <a
          className='btn btn-secondary mx-2'
          href='/container/{id}/action/kill'
        >
          Kill TODO
        </a>
        <a
          className='btn btn-secondary mx-2'
          href='/container/{id}/action/restart'
        >
          Restart TODO
        </a>
        <a
          className='btn btn-secondary mx-2'
          href='/container/{id}/action/pause'
        >
          Pause TODO
        </a>
        <a
          className='btn btn-secondary mx-2'
          href='/container/{id}/action/resume'
        >
          Resume TODO
        </a>
        <a
          className='btn btn-secondary mx-2'
          href='/container/{id}/action/remove'
        >
          Remove TODO
        </a>
      </div>
    </div>
  );
}
