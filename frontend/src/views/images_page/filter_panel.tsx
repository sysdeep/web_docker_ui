import FilterModel from "./filter_model";

interface FilterPanelProps {
  filter: FilterModel;
  onChange(model: FilterModel): void;
}

export default function FilterPanel({ filter, onChange }: FilterPanelProps) {
  const on_date_click = (e: any, date: string) => {
    e.preventDefault();
    const new_dates = filter.dates.filter((d) => d !== date);
    onChange({ ...filter, dates: new_dates });
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
      <div className='row'>
        <div className='col-6'>
          <form>
            <div className='mb-3'>
              <label htmlFor='images_search_tag' className='form-label'>
                Tag
              </label>
              <input
                type='text'
                className='form-control'
                id='images_search_tag'
                value={filter.search_tag}
                onChange={(e) => onChange({ ...filter, search_tag: e.target.value })}
              />
            </div>
          </form>
        </div>
        <div className='col-6'>
          <p>dates:</p>
          <ul>{dates_list}</ul>
        </div>
      </div>
    </div>
  );
}
