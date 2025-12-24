import ImageListModel from "../../models/image_list_model";
import { format_size } from "../../utils/humanize";
import { Link } from "react-router-dom";
import { route, join_url } from "../../routes";
import IconRemove from "@src/components/icon_remove";
import { useState } from "react";
import { SortState } from "@src/data/sort_state";
import IconSort from "@src/components/icon_sort";

interface ImagesTableProps {
  on_remove(id: string): void;
  on_date(date: string): void;
  images: ImageListModel[];
}

const HEADER_SIZE = "size";
const HEADER_CREATED = "created";

export default function ImagesTable({ images, on_remove, on_date }: ImagesTableProps) {
  const [sortHeader, setSortHeader] = useState<string>(HEADER_CREATED);
  const [sortAsc, setSortAsc] = useState<boolean>(false);

  const onHeaderClick = (header: string) => {
    if (sortHeader === header) {
      return setSortAsc(!sortAsc);
    }

    setSortHeader(header);
  };

  const size_icon_state = makeSortState(HEADER_SIZE, sortHeader, sortAsc);
  const created_icon_state = makeSortState(HEADER_CREATED, sortHeader, sortAsc);

  const sort_k = sortAsc ? 1 : -1;
  const sortedRecords = images.sort((a, b) => {
    if (sortHeader === HEADER_SIZE) {
      return (a.size > b.size ? 1 : -1) * sort_k;
    }

    if (sortHeader === HEADER_CREATED) {
      return (a.created > b.created ? 1 : -1) * sort_k;
    }

    return 0;
  });

  const rows = sortedRecords.map((v: ImageListModel, idx: number) => {
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
    <table className='table table-sm table-striped table-bordered'>
      <thead>
        <tr>
          <th>Tags</th>
          <th role='button' onClick={() => onHeaderClick(HEADER_SIZE)}>
            Size
            <span style={{ float: "right" }}>
              <IconSort state={size_icon_state} />
            </span>
          </th>
          <th role='button' onClick={() => onHeaderClick(HEADER_CREATED)}>
            Created
            <span style={{ float: "right" }}>
              <IconSort state={created_icon_state} />
            </span>
          </th>
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
    const [date, ..._] = created.split(" ");
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

function makeSortState(header_name: string, cur_header: string, cur_asc: boolean): SortState {
  return cur_header === header_name ? (cur_asc ? SortState.asc : SortState.desc) : SortState.none;
}
