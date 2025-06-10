// const humanize = require('humanize');
// TODO
import { filesize } from 'humanize';
import { formatDate, parseISO } from 'date-fns';

export function format_size(value: number): string {
  return filesize(value);
}

// 2024-10-21T14:27:32.75462288Z -> 2024-10-21 17:27:32
export function format_date(value: string): string {
  let date = parseISO(value);
  return formatDate(date, 'yyyy-MM-dd HH:mm:ss');
}
