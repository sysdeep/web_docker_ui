import { ReactNode } from "react";
import IconRefresh from "./icon_refresh";

interface PageTitleProps {
  children: ReactNode;
  onRefresh?: () => void;
  isRefresh?: boolean;
}

export default function PageTitle({ children, onRefresh, isRefresh }: PageTitleProps) {
  const on_refresh = (e: any) => {
    e.preventDefault();
    if (onRefresh) {
      onRefresh();
    }
  };

  return (
    <div className='d-flex flex-row justify-content-between align-items-center'>
      <div className='p-2'>
        <h1>{children}</h1>
      </div>
      <div className='p-2'>
        {onRefresh && (
          <button className='btn btn-outline-secondary btn-sm float-end' onClick={on_refresh}>
            {isRefresh && <span className='spinner-grow spinner-grow-sm' aria-hidden='true'></span>}
            {!isRefresh && <IconRefresh />}
          </button>
        )}
      </div>
    </div>
  );
}
