export function farewell(name: string): string {
  const trimmed = name.trim();
  return trimmed === '' ? 'Goodbye!' : `Goodbye, ${trimmed}!`;
}
