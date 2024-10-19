import React from 'react';
import { ApiFullSecretModel } from '../../services/secrets_service';
import ButtonRemove from '@src/components/button_remove';

interface DetailsFrameProps {
  secret: ApiFullSecretModel;
  on_remove(): void;
}

export default function DetailsFrame({ secret, on_remove }: DetailsFrameProps) {
  return (
    <div>
      {/* <h2>Secret details</h2> */}
      <div>
        <table className='table table-small'>
          <tbody>
            <tr>
              <td>ID</td>
              <td>{secret.secret.id}</td>
            </tr>
            <tr>
              <td>Name</td>
              <td>{secret.secret.name}</td>
            </tr>
            <tr>
              <td>Created</td>
              <td>{secret.secret.created}</td>
            </tr>
            <tr>
              <td>Updated</td>
              <td>{secret.secret.updated}</td>
            </tr>
          </tbody>
        </table>

        <div>
          <ButtonRemove on_remove={on_remove} />
        </div>

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
