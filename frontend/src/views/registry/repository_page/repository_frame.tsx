import ButtonRemove from '@src/components/button_remove';
import { join_url, route } from '@src/routes';
import { RepositoryModel } from '@src/services/registry_service';
import React from 'react';
import { Link } from 'react-router-dom';

type RepositoryFrameProps = {
  repository: RepositoryModel;
  on_repository_remove(): void;
  on_tag_remove(tag: string): void;
};

export default function RepositoryFrame({
  repository,
  on_repository_remove,
  on_tag_remove,
}: RepositoryFrameProps) {
  const tags_view = repository.tags
    .sort((a, b) => {
      return a > b ? 1 : -1;
    })
    .map((tag, idx) => {
      return (
        <span key={idx} className='mx-2'>
          [{tag}]
        </span>
      );
    });

  //   const on_tag_remove = (tag: string) => {
  //     console.log(tag);
  //   };

  //   const on_repository_remove = () => {
  //     console.log('remove repo');
  //   };

  return (
    <div>
      <ul>
        <li>repository: {repository.name}</li>
        <li>tags: {tags_view}</li>
      </ul>

      <ButtonRemove on_remove={on_repository_remove} />

      <TagsTable repository={repository} on_remove={on_tag_remove} />
    </div>
  );
}

type TagsTableProps = {
  repository: RepositoryModel;
  on_remove(tag: string): void;
};
function TagsTable({ repository, on_remove }: TagsTableProps) {
  const on_remove_clicked = (e: any, tag: string) => {
    e.preventDefault();
    on_remove(tag);
  };

  const tags_view = repository.tags
    .sort((a, b) => {
      return a > b ? 1 : -1;
    })
    .map((tag, idx) => {
      return (
        <tr key={idx}>
          <td>
            <Link
              to={join_url(
                route.registry_repository_tag,
                repository.id + '/' + tag,
              )}
            >
              {tag}
            </Link>
          </td>
          <td>
            <a href='#' onClick={(e) => on_remove_clicked(e, tag)}>
              remove
            </a>
          </td>
        </tr>
      );
    });
  return (
    <table className='table table-sm'>
      <tbody>{tags_view}</tbody>
    </table>
  );
}
