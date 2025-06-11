import { ApiFullImageModel } from "../../services/images_service";
import { format_size } from "../../utils/humanize";
import ButtonRemove from "@src/components/button_remove";

interface DetailsFrameProps {
  image: ApiFullImageModel;
  on_remove(): void;
}

export default function DetailsFrame({ image, on_remove }: DetailsFrameProps) {
  const tags_view = image.image.repo_tags.map((tag, idx) => {
    return <li key={idx}>{tag}</li>;
  });
  return (
    <div className='card my-2'>
      <div className='card-body'>
        <h2>Details</h2>
        <div>
          <table className='table table-sm'>
            <tbody>
              <tr>
                <td>ID</td>
                <td>
                  <code>{image.image.id}</code>
                </td>
              </tr>
              <tr>
                <td>Tags</td>
                <td>
                  <ul className='table-ui'>
                    <code>{tags_view}</code>
                  </ul>
                </td>
              </tr>
              <tr>
                <td>Parent</td>
                <td>{image.image.parent}</td>
              </tr>
              <tr>
                <td>Comment</td>
                <td>{image.image.comment}</td>
              </tr>
              <tr>
                <td>Size</td>
                <td>{format_size(image.image.size)}</td>
              </tr>
              <tr>
                <td>Created</td>
                <td>{image.image.created}</td>
              </tr>
            </tbody>
          </table>

          <div className='clearfix'>
            <div className='float-end'>
              {/* TODO: pull - see portainer */}
              {/* TODO: push - see portainer */}
              <ButtonRemove on_remove={on_remove} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
