import { ContainerTopModel } from '@src/services/containers_service';
import React from 'react';

type TopFrameProps = {
  top_model: ContainerTopModel | null;
};

export default function TopFrame({ top_model }: TopFrameProps) {
  if (top_model === null) {
    return (
      <div className='box'>
        <h2>Top</h2>
        <p>no data</p>
      </div>
    );
  }

  const headers_view = top_model.titles.map((title, idx) => {
    return <th key={idx}>{title}</th>;
  });

  const body_row = (row: string[]) => {
    return row.map((data, idx) => {
      return <td key={idx}>{data}</td>;
    });
  };

  const body_view = top_model.processes.map((procs, idx) => {
    return <tr key={idx}>{body_row(procs)}</tr>;
  });

  return (
    <div className='box'>
      <h2>Top</h2>
      <table className='table table-sm'>
        <thead>
          <tr>{headers_view}</tr>
        </thead>
        <tbody>{body_view}</tbody>
      </table>
    </div>
  );
}
