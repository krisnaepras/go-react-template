export function validateTitle(title: string): string | null {
  const length = Array.from(title.trim()).length
  if (length === 0 || length > 200) return 'Judul harus berisi 1–200 karakter.'
  return null
}
