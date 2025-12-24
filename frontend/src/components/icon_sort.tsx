// https://icons.getbootstrap.com/

import { SortState } from "@src/data/sort_state";

type Props = {
  state: SortState;
};

export default function IconSort({ state }: Props) {
  let sort_style = "bi-filter";

  switch (state) {
    case SortState.asc:
      sort_style = "bi-sort-down";
      break;

    case SortState.desc:
      sort_style = "bi-sort-up";
      break;

    default:
      break;
  }

  return <i className={`bi ${sort_style}`}></i>;
}
