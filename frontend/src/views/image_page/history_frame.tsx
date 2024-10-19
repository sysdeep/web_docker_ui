import { ApiFullImageModel } from '../../services/images_service';
import React from 'react';
import { format_size } from '../../utils/humanize';

interface HistoryFrameProps {
  image: ApiFullImageModel;
}

export default function HistoryFrame({ image }: HistoryFrameProps) {
  const tags_list = (tags: string[]) => {
    return tags.map((tag, idx) => {
      return <span key={idx}>{tag}</span>;
    });
  };

  const rows_view = image.history.map((history, idx) => {
    return (
      <tr key={idx}>
        <td>{history.id}</td>
        <td>{tags_list(history.tags)}</td>
        <td>{format_size(history.size)}</td>
        <td>{history.created}</td>
      </tr>
    );
  });

  return (
    <div className='box'>
      <h2>History</h2>
      <table className='table is-striped is-fullwidth'>
        <thead>
          <tr>
            <th>ID</th>
            <th>Tags</th>
            <th>Size</th>
            <th>Created</th>
          </tr>
        </thead>
        <tbody>{rows_view}</tbody>
      </table>
    </div>
  );
}
