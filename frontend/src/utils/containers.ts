export function strip_container_name(name: string) {
  if (name.startsWith('/')) {
    return name.slice(1);
  }
  return name;
}
