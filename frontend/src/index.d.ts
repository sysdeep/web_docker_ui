declare module "*.css";
declare module "*.png";
declare module "*.jpg";
declare module "*.jpeg";

/**
 * Formats the value like a 'human-readable' file size (i.e. '13 KB', '4.1 MB', '102 bytes', etc).
 *
 * For example:
 * If value is 123456789, the output would be 117.7 MB.
 */
declare module "humanize" {
  declare function filesize(
    filesize: number,
    kilo?: number, // 1024
    decimals?: number, // 0
    decPoint?: string, // '.'
    thousandsSep?: string, // ','
    suffixSep?: string, // ' '
  ): string;
}
