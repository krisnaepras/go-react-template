import type { ReactNode } from 'react'
import { Link } from '@tanstack/react-router'
import { ArrowRight, Boxes, Database, LockKeyhole, Rocket, Workflow } from 'lucide-react'

export function DashboardPage() {
  return <>
    <header className="hero">
      <div className="eyebrow">WELCOME TO YOUR STARTER</div>
      <h1>Fondasi web kantor yang siap dikembangkan.</h1>
      <p>Go Native untuk API, React + Vite untuk UI, PostgreSQL untuk data. Simpel, modular, dan mudah di-maintain bersama AI coding agent.</p>
      <Link to="/tasks" className="primary-button">Buka demo CRUD <ArrowRight size={17}/></Link>
    </header>
    <div className="section-title"><h2>Yang sudah disiapkan</h2><span>Reusable foundations</span></div>
    <div className="feature-grid">
      <Feature icon={<Boxes size={23}/>} title="Modular monolith" desc="Pisahkan domain, service, repository, dan handler per modul."/>
      <Feature icon={<Database size={23}/>} title="PostgreSQL" desc="Migration versioned, query berparameter, pgxpool."/>
      <Feature icon={<Workflow size={23}/>} title="API contract" desc="OpenAPI untuk menjaga sinkronisasi frontend dan backend."/>
      <Feature icon={<LockKeyhole size={23}/>} title="Secure by default" desc="Demo API dimatikan di production sampai autentikasi tersedia."/>
      <Feature icon={<Rocket size={23}/>} title="CI & Docker" desc="Pemeriksaan otomatis, container build, dan dev compose."/>
    </div>
    <div className="callout"><strong>Langkah production:</strong> Tambahkan SSO/OIDC, session management, RBAC, audit log, dan backup sebelum menyimpan data kantor yang nyata.</div>
  </>
}
function Feature({ icon, title, desc }: { icon: ReactNode; title: string; desc: string }) {
  return <div className="feature-card"><div className="feature-icon">{icon}</div><h3>{title}</h3><p>{desc}</p></div>
}
