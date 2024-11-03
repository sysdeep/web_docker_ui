import React from 'react';

type Props = {};
export default function FrameLabels({}: Props) {
  // const items_view = env.map((row, idx) => {
  //   const [key, value] = row.split('=');
  //   return (
  //     <tr key={idx}>
  //       <td>{key}</td>
  //       <td>{value}</td>
  //     </tr>
  //   );
  // });
  return (
    <div className='card mt-2'>
      <div className='card-body'>
        <h3>Labels</h3>

        <table className='table'>{/* <tbody>{items_view}</tbody> */}</table>
      </div>
    </div>
  );
}
