import { ApiInfoModel } from "../../services/info_service";

type Props = {
  info: ApiInfoModel;
};

export default function ClientFrame({ info }: Props) {
  return (
    <div className='card mb-2'>
      <div className='card-body'>
        <h2>Client</h2>
        <table className='table table-sm'>
          <tbody>
            <tr>
              <td>DaemonHost</td>
              <td className='text-end'> {info.daemon_host}</td>
            </tr>
            <tr>
              <td>ClientVersion</td>
              <td className='text-end'>{info.client_version}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
}
