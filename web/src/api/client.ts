export type Task = { id: number; title: string; completed: boolean; created_at: string }
type Envelope<T> = { data: T }
type APIError = { error?: { code?: string; message?: string } }

export async function request<T>(url: string, init?: RequestInit): Promise<T> {
  const response = await fetch(url, {
    ...init,
    credentials: 'same-origin',
    headers: { ...(init?.body ? { 'Content-Type': 'application/json' } : {}), ...init?.headers },
  })
  if (!response.ok) {
    const body = await response.json().catch(() => ({})) as APIError
    throw new Error(body.error?.message ?? `HTTP ${response.status}`)
  }
  if (response.status === 204) return undefined as T
  return response.json() as Promise<T>
}
export const tasksAPI = {
  list: async () => (await request<Envelope<Task[]>>('/api/v1/tasks')).data,
  create: async (title: string) => (await request<Envelope<Task>>('/api/v1/tasks', { method: 'POST', body: JSON.stringify({ title }) })).data,
  setCompleted: async (id: number, completed: boolean) => (await request<Envelope<Task>>(`/api/v1/tasks/${id}`, { method: 'PATCH', body: JSON.stringify({ completed }) })).data,
  remove: async (id: number) => request<void>(`/api/v1/tasks/${id}`, { method: 'DELETE' }),
}
