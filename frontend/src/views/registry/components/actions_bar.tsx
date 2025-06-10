import IconRefresh from '@src/components/icon_refresh';
import { route } from '@src/routes';
import { Link } from 'react-router-dom';

type ActionsBarProps = {
  on_garbage(): void;
  on_restart(): void;
};
export default function ActionsBar({ on_garbage, on_restart }: ActionsBarProps) {
  const on_garbage_click = (e: any) => {
    e.preventDefault();
    on_garbage();
  };

  const on_restart_click = (e: any) => {
    e.preventDefault();
    on_restart();
  };

  return (
    <div className='clearfix'>
      <div className='float-end'>
        <Link className='btn btn-primary me-2' to={route.registry_repositories}>
          <i className='bi bi-link'></i>
          &nbsp; Go to Catalog
        </Link>

        <Link className='btn btn-primary me-2' to={route.registry_repositories_smart}>
          <i className='bi bi-link'></i>
          &nbsp; Go to Repos
        </Link>

        <Link className='btn btn-primary me-2' to={route.registry_repositories_tree}>
          <i className='bi bi-link'></i>
          &nbsp; Go to Tree
        </Link>

        <button className='btn btn-secondary me-2' onClick={on_garbage_click}>
          <i className='bi bi-trash3'></i>
          &nbsp; Garbage
        </button>
        <button className='btn btn-secondary' onClick={on_restart_click}>
          <IconRefresh />
          Restart
        </button>
      </div>
    </div>
  );
}
