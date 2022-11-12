import React, { ReactNode } from 'react'
import { Navigate, Outlet, Route, Routes } from 'react-router-dom'
import PageEmpty from './components/PageLayout/PageEmpty'
import { GlobalProvider } from './context/GlobalContext'
import Admin from './pages/Admin'
import Home from './pages/Home'
import Login from './pages/Login'
import Results from './pages/Results'
import Rules from './pages/Rules'
import { isAuthenticated } from './services/auth'

const ProtectedRoute = ({
  redirectPath = '/login',
  children,
}: {
  redirectPath?: string
  children?: ReactNode
}) => {
  if (!isAuthenticated()) {
    return <Navigate to={redirectPath} replace />
  }

  const content = children || <Outlet />

  return <GlobalProvider>{content}</GlobalProvider>
}

const AppRouter = () => {
  return (
    <Routes>
      <Route path='login' element={<Login />} />
      <Route path='*' element={<PageEmpty />} />
      <Route path='/' element={<Home />} />
      <Route path='home' element={<Home />} />
      <Route path='/rules' element={<Rules />} />
      <Route path='/results' element={<Results />} />
      <Route element={<ProtectedRoute />}>
        <Route path='/admin' element={<Admin />} />
      </Route>
    </Routes>
  )
}

export default AppRouter
