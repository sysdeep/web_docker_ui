import React from 'react';
import IconRefresh from './icon_refresh';

type ButtonRemoveProps = {
  //   disabled: boolean;
  on_refresh(): void;
  loading?: boolean;
};
export default function ButtonRefresh({
  //   disabled = false,
  on_refresh,
  loading = false,
}: ButtonRemoveProps) {
  const on_click = (e: any) => {
    e.preventDefault();
    on_refresh();
  };

  return (
    <button className='btn btn-secondary' onClick={on_click}>
      {loading && (
        <span
          className='spinner-grow spinner-grow-sm'
          aria-hidden='true'
        ></span>
      )}
      {!loading && <IconRefresh />} Refresh
    </button>
  );
}
