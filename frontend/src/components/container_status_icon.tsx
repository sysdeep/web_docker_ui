import React from 'react';

const COLORS_MAP: { [key: string]: string } = {
  exited: 'text-secondary',
  running: 'text-success',
  paused: 'text-warning',
};

const ICONS_MAP: { [key: string]: string } = {
  exited: 'bi-stop-fill',
  running: 'bi-play-fill',
  paused: 'bi-pause-fill',
};

const UNKNOWN_ICON = 'bi-question-diamond-fill';

type ContainerStatusIconProps = {
  status: string;
};

export default function ContainerStatusIcon({ status }: ContainerStatusIconProps) {
  let icon_class = ICONS_MAP[status];
  if (icon_class === undefined) icon_class = UNKNOWN_ICON;

  let color_class = COLORS_MAP[status];
  if (color_class === undefined) color_class = '';

  return <i className={`bi ${icon_class} ${color_class}`} title={status}></i>;
}
