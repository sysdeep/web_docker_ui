import React from 'react';
import { Outlet } from 'react-router-dom';
import TopNavBar from './top_nav_bar';

export default function Application() {
  return (
    <>
      <TopNavBar />
      <div className='container'>
        <div className='mt-2'>
          <Outlet />
        </div>
      </div>
    </>
  );
}

// https://reactrouter.com/en/main/start/tutorial
