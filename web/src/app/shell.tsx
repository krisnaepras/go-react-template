import type { ReactNode } from 'react'
import { Link } from '@tanstack/react-router'
import { BookOpenText, CheckSquare2, LayoutDashboard, Layers3 } from 'lucide-react'

export function AppShell({ children }: { children: ReactNode }) {
  return <div className="app-shell">
    <aside className="sidebar">
      <div className="brand"><Layers3 size={23}/><span>Office Starter</span></div>
      <div className="nav-label">WORKSPACE</div>
      <nav aria-label="Navigasi utama">
        <Link to="/" activeProps={{ className: 'nav-link active' }} inactiveProps={{ className: 'nav-link' }}><LayoutDashboard size={18}/> Dashboard</Link>
        <Link to="/tasks" activeProps={{ className: 'nav-link active' }} inactiveProps={{ className: 'nav-link' }}><CheckSquare2 size={18}/> Demo tasks</Link>
      </nav>
      <div className="sidebar-footer"><BookOpenText size={18}/><span>Go Native · React · PostgreSQL</span></div>
    </aside>
    <main className="main-content"><div className="topline"><span className="eyebrow">REUSABLE PROJECT TEMPLATE</span><span className="version">v0.1.0</span></div>{children}</main>
  </div>
}
