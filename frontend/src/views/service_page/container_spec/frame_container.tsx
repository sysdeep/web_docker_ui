import { ContainerSpec } from '@src/models/service';
import React from 'react';

type Props = {
  container: ContainerSpec;
};
export default function FrameContainer({ container }: Props) {
  return (
    <div className='card mt-2'>
      <div className='card-body'>
        <h3>Container</h3>

        <table className='table'>
          <tbody>
            {/* TODO */}
            {/* <tr>
              <td>CMD</td>
              <td></td>
            </tr> */}
            <tr>
              <td>Args</td>
              <td>
                <ul>
                  {container.args.map((r, idx) => {
                    return (
                      <li key={idx}>
                        <code>{r}</code>
                      </li>
                    );
                  })}
                </ul>
              </td>
            </tr>
            {/* TODO */}
            {/* <tr>
              <td>User</td>
              <td></td>
            </tr>
            <tr>
              <td>Working directory</td>
              <td></td>
            </tr>
            <tr>
              <td>Stop grace period</td>
              <td></td>
            </tr> */}
            <tr>
              <td>Image</td>
              <td>{container.image}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
}
