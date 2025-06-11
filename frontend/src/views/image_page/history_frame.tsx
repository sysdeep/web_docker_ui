import { ApiFullImageModel } from "../../services/images_service";
import { format_size } from "../../utils/humanize";

interface HistoryFrameProps {
  image: ApiFullImageModel;
}

export default function HistoryFrame({ image }: HistoryFrameProps) {
  // const tags_list = (tags: string[]) => {
  //   return tags.map((tag, idx) => {
  //     return <span key={idx}>{tag}</span>;
  //   });
  // };

  const rows_view = image.history.map((history, idx) => {
    return (
      <tr key={idx}>
        <td>{idx}</td>
        <td>
          {/* TODO: много текста, можно обрезать и показывать полностью в тултипе - see portainer */}
          <code>{history.created_by}</code>
        </td>
        {/* <td>{tags_list(history.tags)}</td> */}
        <td>{format_size(history.size)}</td>
        <td>{history.created}</td>
      </tr>
    );
  });

  return (
    <div className='card my-2'>
      <div className='card-body'>
        <h2>History</h2>
        <table className='table table-striped table-sm'>
          <thead>
            <tr>
              <th>Order</th>
              <th>Layer</th>
              {/* TODO: вроде и не надо... */}
              {/* <th>Tags</th> */}
              <th>Size</th>
              <th>Created</th>
            </tr>
          </thead>
          <tbody>{rows_view}</tbody>
        </table>
      </div>
    </div>
  );
}
