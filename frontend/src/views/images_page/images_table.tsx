import React, { useState } from 'react';
import ImageListModel from '../../models/image_list_model';
import FilterModel from './filter_model';
import { format_size } from '../../utils/humanize';
import { Link } from 'react-router-dom';
import { route, join_url } from '../../routes';
import IconRemove from '@src/components/icon_remove';

interface ImagesTableProps {
  on_remove(id: string): void;
  on_date(date: string): void;
  images: ImageListModel[];
  filter: FilterModel;
}

export default function ImagesTable({
  images,
  filter,
  on_remove,
  on_date,
}: ImagesTableProps) {
  // const data = [1, 2, 3];

  // const [images, setImages] = useState<number[]>(data);

  // const on_remove = (uid: number) => {
  //   console.log('on remove: ' + uid);
  //   const new_images = images.filter((v: number) => v !== uid);
  //   setImages(new_images);
  // };

  const rows = images
    .filter((v: ImageListModel) => {
      if (filter.dates.length === 0) {
        return true;
      }
      const [date, ..._] = v.created.split(' ');
      return filter.dates.includes(date);
    })
    .map((v: ImageListModel, idx: number) => {
      return (
        <TableRow
          uid={v.id}
          tags={v.tags}
          size={v.size}
          created={v.created}
          on_remove={on_remove}
          on_date={on_date}
          key={idx}
        ></TableRow>
      );
    });

  return (
    <table className='table table-sm table-striped'>
      <thead>
        <tr>
          <th>Tags</th>
          <th>Size</th>
          <th>Created</th>
          <th>Options</th>
        </tr>
      </thead>
      <tbody>{rows}</tbody>
    </table>
  );
}

interface Props {
  uid: string;
  tags: string[];
  size: number;
  created: string;
  on_remove(uid: string): void;
  on_date(date: string): void;
}

function TableRow({ uid, tags, size, created, on_remove, on_date }: Props) {
  const on_remove_click = (e: any) => {
    e.preventDefault();
    on_remove(uid);
  };

  const on_date_click = (e: any) => {
    e.preventDefault();
    const [date, ..._] = created.split(' ');
    on_date(date);
  };

  const tags_ui = () => {
    if (tags.length === 0) {
      return <li key={1}>no tag</li>;
    }

    return tags.map((tag, idx) => {
      return (
        <li key={idx}>
          <Link to={join_url(route.image, uid)}>{tag}</Link>
        </li>
      );
    });
  };

  return (
    <tr>
      <td>
        <ul className='table-ui'>{tags_ui()}</ul>
      </td>
      <td className='text-right'>{format_size(size)}</td>
      <td>
        <a href='#' onClick={on_date_click}>
          {created}
        </a>
      </td>
      <td>
        <a href='#' onClick={on_remove_click}>
          <IconRemove />
          &nbsp; remove
        </a>
      </td>
    </tr>
  );
}
