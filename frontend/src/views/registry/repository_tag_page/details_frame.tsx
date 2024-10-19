import { RepositoryModel, TagManifest } from '@src/services/registry_service';
import { format_size } from '@src/utils/humanize';
import React from 'react';

type DetailsFrameProps = {
  manifest: TagManifest;
  repository: RepositoryModel;
};
export default function DetailsFrame({
  manifest,
  repository,
}: DetailsFrameProps) {
  return (
    <div>
      <div>
        <table className='table table-sm'>
          <tbody>
            <tr>
              <td>repository</td>
              <td>{repository.name}</td>
            </tr>
            <tr>
              <td>tag</td>
              <td>{manifest.name}</td>
            </tr>
            <tr>
              <td>schema version</td>
              <td>{manifest.schema_version}</td>
            </tr>
            <tr>
              <td>media type</td>
              <td>{manifest.media_type}</td>
            </tr>
            <tr>
              <td>total size</td>
              <td>{format_size(manifest.total_size)}</td>
            </tr>
            <tr>
              <td>content digest</td>
              <td>{manifest.content_digest}</td>
            </tr>
          </tbody>
        </table>
      </div>
      {/* TODO: layers and Config */}
    </div>
  );
}
