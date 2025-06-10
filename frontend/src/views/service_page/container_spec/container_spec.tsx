import { ContainerSpec } from '@src/models/service';
import FrameEnv from './frame_env';
import FrameContainer from './frame_container';
import FrameLabels from './frame_labels';
import FrameMounts from './frame_mounts';

type Props = {
  container: ContainerSpec;
};

export default function ContainerSpec({ container }: Props) {
  return (
    <div>
      <FrameContainer container={container} />
      <FrameEnv env={container.env} />
      <FrameLabels />
      <FrameMounts />
    </div>
  );
}
