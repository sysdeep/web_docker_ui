import React from 'react';

type ActionsBarProps = {
  on_garbage(): void;
  on_restart(): void;
};
export default function ActionsBar({
  on_garbage,
  on_restart,
}: ActionsBarProps) {
  const on_garbage_click = (e: any) => {
    e.preventDefault();
    on_garbage();
  };

  const on_restart_click = (e: any) => {
    e.preventDefault();
    on_restart();
  };

  return (
    <div className=''>
      <button className='btn btn-secondary' onClick={on_garbage_click}>
        Garbage
      </button>
      <button className='btn btn-secondary' onClick={on_restart_click}>
        Restart
      </button>
    </div>
  );
}
