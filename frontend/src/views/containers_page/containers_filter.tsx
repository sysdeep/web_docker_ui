import { container_status_map, ContainerStatus } from "@src/data/container_status";

type Props = {
  initial_filter: ContainersFilterModel;
  onChange: (model: ContainersFilterModel) => void;
};

export type ContainersFilterModel = {
  status: string;
  search: string;
};

const statuses = [
  "all",
  container_status_map[ContainerStatus.running],
  container_status_map[ContainerStatus.paused],
  container_status_map[ContainerStatus.exited],
];

export default function ContainersFilter({ initial_filter, onChange }: Props) {
  // on status radio selected
  const onStatus = (idx: number) => {
    const st = statuses[idx];
    onChange({ ...initial_filter, status: st });
  };

  const statusValue = initial_filter.status;
  const statusesForm = statuses.map((st, idx) => {
    return <StatusRadio label={st} key={idx} idx={idx} curValue={statusValue} onChange={onStatus} />;
  });

  // on search
  const on_search = (term: string) => {
    onChange({ ...initial_filter, search: term });
  };

  return (
    <div>
      <div className='mb-3 row'>
        <legend className='col-form-label col-sm-2 pt-0'>State</legend>
        <div className='col-sm-10'>{statusesForm}</div>
      </div>

      <div className='mb-3 row'>
        <label htmlFor='nameSearch' className='col-sm-2 col-form-label'>
          Search
        </label>
        <div className='col-sm-10'>
          <input type='text' className='form-control' id='nameSearch' onChange={(e) => on_search(e.target.value)} />
        </div>
      </div>
    </div>
  );
}

//-----------------------------------------------------------------------------
type StatusRadioProps = {
  label: string;
  idx: number;
  curValue: string;
  onChange: (idx: number) => void;
};

function StatusRadio({ label, idx, curValue, onChange }: StatusRadioProps) {
  return (
    <div className='form-check form-check-inline'>
      <input
        className='form-check-input'
        type='radio'
        name='inlineRadioOptions'
        id={`inlineRadio_${idx}`}
        checked={label === curValue}
        onChange={() => onChange(idx)}
      />
      <label className='form-check-label' htmlFor={`inlineRadio_${idx}`}>
        {label}
      </label>
    </div>
  );
}
