import React, { useEffect, useState } from 'react';

import ImagesPage from '../views/images_page/images_page';

const Application: React.FC = () => {
  const [counter, setCounter] = useState(0);
  const [darkTheme, setDarkTheme] = useState(true);

  /**
   * On component mount
   */
  useEffect(() => {
    console.log('application mounted');
  }, []);

  /**
   * On Dark theme change
   */
  // useEffect(() => {
  //   if (darkTheme) {
  //     localStorage.setItem('dark-mode', '1');
  //     document.body.classList.add('dark-mode');
  //   } else {
  //     localStorage.setItem('dark-mode', '0');
  //     document.body.classList.remove('dark-mode');
  //   }
  // }, [darkTheme]);

  /**
   * Toggle Theme
   */

  return (
    <div id='erwt'>
      <ImagesPage />
    </div>
  );
};

export default Application;
