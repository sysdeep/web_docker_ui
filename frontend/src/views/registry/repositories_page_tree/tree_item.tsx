import { join_url, route } from '@src/routes';
import { Link } from 'react-router-dom';

export type TreeItemData = {
  id: string;
  name: string;
  childrens: TreeItemData[];
};

type TreeItemProps = {
  root: TreeItemData;
};

export default function TreeItem({ root }: TreeItemProps) {
  const childrens = root.childrens.map((child, idx) => {
    return (
      <li key={idx}>
        <TreeItem root={child} />
      </li>
    );
  });

  return (
    <div>
      <div>
        {!root.childrens.length ? (
          <Link to={join_url(route.registry_repository, root.id)}>{root.name}</Link>
        ) : (
          root.name
        )}
        {/* {root.name} */}
      </div>
      <div>
        <ul>{childrens}</ul>
      </div>
    </div>
  );
}
