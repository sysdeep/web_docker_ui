import { ApiFullVolumeModel } from "../../services/volumes_service";

interface DetailsFrameProps {
  volume: ApiFullVolumeModel;
}

export default function DetailsFrame({ volume }: DetailsFrameProps) {
  return (
    <div className='box'>
      {/* <h2>Volume details</h2> */}
      <div>
        <table className='table is-fullwidth'>
          <tbody>
            <tr>
              <td>Name</td>
              <td>{volume.volume.name}</td>
            </tr>
            <tr>
              <td>Created</td>
              <td>{volume.volume.created}</td>
            </tr>
            <tr>
              <td>Mount path</td>
              <td>{volume.volume.mount_point}</td>
            </tr>
            <tr>
              <td>Driver</td>
              <td>{volume.volume.driver}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
}
