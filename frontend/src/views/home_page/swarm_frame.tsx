import { ApiInfoModel } from "../../services/info_service";

type Props = {
  info: ApiInfoModel;
};

export default function SwarmFrame({ info }: Props) {
  return (
    <div className='card mb-2'>
      <div className='card-body'>
        <h2>Swarm</h2>

        <table className='table table-sm'>
          <tbody>
            <tr>
              <td>Node id</td>
              <td className='text-end'>{info.system.swarm.node_id}</td>
            </tr>
            <tr>
              <td>Node address</td>
              <td className='text-end'>{info.system.swarm.node_addr}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
}
