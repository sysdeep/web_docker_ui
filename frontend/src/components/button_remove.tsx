import React from 'react';
import IconRemove from './icon_remove';

interface ButtonRemoveProps {
  //   disabled: boolean;
  on_remove(): void;
}
export default function ButtonRemove({
  //   disabled = false,
  on_remove,
}: ButtonRemoveProps) {
  const on_click = (e: any) => {
    e.preventDefault();
    on_remove();
  };

  return (
    <button className='btn btn-danger' onClick={on_click}>
      <IconRemove />
      &nbsp; Remove
    </button>
  );
}
