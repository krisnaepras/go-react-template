import { createRootRoute, createRoute, createRouter, Outlet } from '@tanstack/react-router'
import { AppShell } from './shell'
import { DashboardPage } from '../features/dashboard/dashboard-page'
import { TasksPage } from '../features/tasks/tasks-page'

const rootRoute = createRootRoute({ component: () => <AppShell><Outlet /></AppShell> })
const homeRoute = createRoute({ getParentRoute: () => rootRoute, path: '/', component: DashboardPage })
const tasksRoute = createRoute({ getParentRoute: () => rootRoute, path: '/tasks', component: TasksPage })
const routeTree = rootRoute.addChildren([homeRoute, tasksRoute])
export const router = createRouter({ routeTree, defaultPreload: 'intent' })
declare module '@tanstack/react-router' {
  interface Register { router: typeof router }
}
