import React from 'react';
import FilterModel from './filter_model';

interface FilterPanelProps {
  filter: FilterModel;
  on_date(date: string): void;
}

export default function FilterPanel({ filter, on_date }: FilterPanelProps) {
  const on_date_click = (e: any, date: string) => {
    e.preventDefault();
    on_date(date);
  };

  const dates_list = filter.dates.map((date: string, idx: number) => {
    return (
      <li key={idx}>
        <a href='#' onClick={(e) => on_date_click(e, date)}>
          {date}
        </a>
      </li>
    );
  });

  return (
    <div>
      <p>Filter</p>
      <p>dates:</p>
      <ul>{dates_list}</ul>
    </div>
  );
}
