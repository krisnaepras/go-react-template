import { useState, type FormEvent } from 'react'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { Check, Loader2, Plus, Trash2 } from 'lucide-react'
import { tasksAPI } from '../../api/client'
import { validateTitle } from './validation'

export function TasksPage() {
  const [title, setTitle] = useState('')
  const [validationError, setValidationError] = useState<string | null>(null)
  const queryClient = useQueryClient()
  const tasks = useQuery({ queryKey: ['demo-tasks'], queryFn: tasksAPI.list })
  const refresh = () => queryClient.invalidateQueries({ queryKey: ['demo-tasks'] })
  const create = useMutation({ mutationFn: tasksAPI.create, onSuccess: () => { setTitle(''); void refresh() } })
  const toggle = useMutation({ mutationFn: ({ id, done }: { id: number; done: boolean }) => tasksAPI.setCompleted(id, done), onSuccess: () => { void refresh() } })
  const remove = useMutation({ mutationFn: tasksAPI.remove, onSuccess: () => { void refresh() } })
  function submit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    const err = validateTitle(title)
    setValidationError(err)
    if (!err) create.mutate(title.trim())
  }
  return <>
    <div className="page-heading"><div><div className="eyebrow">EXAMPLE MODULE</div><h1>Demo Tasks</h1><p>CRUD contoh untuk meniru struktur fitur baru.</p></div><span className="demo-tag">Development only</span></div>
    <section className="panel">
      <h2>Buat task baru</h2>
      <form className="task-form" onSubmit={submit}>
        <label className="sr-only" htmlFor="task-title">Judul task</label>
        <input id="task-title" value={title} maxLength={200} onChange={e => { setTitle(e.target.value); setValidationError(null) }} placeholder="Contoh: Review pengajuan pembelian" />
        <button className="primary-button" disabled={create.isPending} type="submit"><Plus size={17}/> Tambah</button>
      </form>
      {validationError && <p className="form-error" role="alert">{validationError}</p>}
      {create.isError && <p className="form-error" role="alert">{create.error.message}</p>}
    </section>
    <section className="panel task-panel">
      <div className="panel-heading"><h2>Daftar tasks</h2><span>{tasks.data?.length ?? 0} tasks</span></div>
      {tasks.isPending && <div className="empty-state"><Loader2 className="spinner"/> Memuat data...</div>}
      {tasks.isError && <div className="notice"><strong>API belum siap.</strong><p>{tasks.error.message}. Pastikan PostgreSQL, migrasi, dan server Go berjalan dengan DEMO_MODE=true.</p></div>}
      {tasks.isSuccess && tasks.data.length === 0 && <div className="empty-state">Belum ada task. Tambahkan task pertama di atas.</div>}
      {tasks.isSuccess && <ul className="task-list">{tasks.data.map(task => <li key={task.id} className="task-row">
        <button type="button" aria-label={task.completed ? 'Tandai belum selesai' : 'Tandai selesai'} aria-pressed={task.completed} className={`check-button ${task.completed ? 'checked' : ''}`} disabled={toggle.isPending} onClick={() => toggle.mutate({ id: task.id, done: !task.completed })}>{task.completed && <Check size={15}/>}</button>
        <div className="task-copy"><span className={task.completed ? 'task-done' : ''}>{task.title}</span><small>{new Date(task.created_at).toLocaleString('id-ID')}</small></div>
        <button type="button" className="icon-button" aria-label={`Hapus ${task.title}`} disabled={remove.isPending} onClick={() => remove.mutate(task.id)}><Trash2 size={17}/></button>
      </li>)}</ul>}
      {(toggle.isError || remove.isError) && <p className="form-error" role="alert">{toggle.error?.message ?? remove.error?.message}</p>}
    </section>
  </>
}
