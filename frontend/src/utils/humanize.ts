const humanize = require('humanize');

export function format_size(value: number): string {
  return humanize.filesize(value);
}
