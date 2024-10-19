import React from 'react';
import { ApiFullConfigModel } from '../../services/configs_service';
import ButtonRemove from '../../components/button_remove';

interface DetailsFrameProps {
  config: ApiFullConfigModel;
  on_remove(): void;
}

export default function DetailsFrame({ config, on_remove }: DetailsFrameProps) {
  return (
    <div>
      {/* <h2>Secret details</h2> */}
      <div>
        <table className='table-small'>
          <tbody>
            <tr>
              <td>ID</td>
              <td>{config.config.id}</td>
            </tr>
            <tr>
              <td>Name</td>
              <td>{config.config.name}</td>
            </tr>
            <tr>
              <td>Created</td>
              <td>{config.config.created}</td>
            </tr>
            <tr>
              <td>Updated</td>
              <td>{config.config.updated}</td>
            </tr>
          </tbody>
        </table>

        <div>
          <ButtonRemove on_remove={on_remove} />
        </div>

        <pre>{config.config.data_text}</pre>

        {/* <div>
      <!-- <a href="/volumes/actions/remove/{ . }" class="button error"> -->
      <!--   <i class="fa fa-trash-o" aria-hidden="true"></i> -->
      <!--   Remove -->
      <!-- </a> -->
    </div> */}
      </div>
    </div>
  );
}
