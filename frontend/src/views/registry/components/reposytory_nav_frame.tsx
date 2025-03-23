import { join_url, route } from '@src/routes';
import { RepositoryModel } from '@src/services/registry_service';
import React from 'react';
import { Link } from 'react-router-dom';

type RepositoryNavFrameProps = {
  repository: RepositoryModel;
};

export default function RepositoryNavFrame({ repository }: RepositoryNavFrameProps) {
  const tags_view = repository.tags
    .sort((a, b) => {
      return a > b ? 1 : -1;
    })
    .map((tag, idx) => {
      return (
        <span key={idx} className='mx-2'>
          <Link to={join_url(route.registry_repository_tag, repository.id + '/' + tag)}>{tag}</Link>
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
        <li>
          repository: <Link to={join_url(route.registry_repository, repository.id)}>{repository.name}</Link>
        </li>
        <li>tags: {tags_view}</li>
      </ul>
    </div>
  );
}
