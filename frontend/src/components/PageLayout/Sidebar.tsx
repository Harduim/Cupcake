import { EuiPageSidebar, EuiSideNav, htmlIdGenerator } from '@elastic/eui'
import { useContext, useState } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'
import GlobalContext from '../../context/GlobalContext'

const SideBar = ({ title }: { title: string }) => {
  const { pathname } = useLocation()
  const { isLoading } = useContext(GlobalContext)
  const navigate = useNavigate()
  const [isSideNavOpenOnMobile, setisSideNavOpenOnMobile] = useState(false)
  const toggleOpenOnMobile = () => {
    setisSideNavOpenOnMobile(!isSideNavOpenOnMobile)
  }

  if (isLoading) return null
  const sideNav = [
    {
      name: 'Cupcake',
      id: htmlIdGenerator('Cupcake')(),
      items: [
        {
          name: 'Página Inicial',
          id: htmlIdGenerator('Página Inicial')(),
          onClick: () => navigate('/home'),
          isSelected: pathname === '/home',
        },
        {
          name: 'Resultados',
          id: htmlIdGenerator('Resultados')(),
          isSelected: pathname === '/results',
          onClick: () => navigate('/results'),
        },
        {
          name: 'Regras',
          id: htmlIdGenerator('Regras')(),
          onClick: () => navigate('/rules'),
          isSelected: pathname === '/rules',
        },
        {
          name: 'Administração',
          id: htmlIdGenerator('Administração')(),
          disabled: true,
          isSelected: pathname === '/admin',
          onClick: () => navigate('/admin'),
        },
      ],
    },
  ]

  return (
    <EuiPageSidebar paddingSize='m' sticky>
      <EuiSideNav
        aria-label='sidebar-nav'
        mobileTitle={title}
        toggleOpenOnMobile={toggleOpenOnMobile}
        isOpenOnMobile={isSideNavOpenOnMobile}
        items={sideNav}
      />
    </EuiPageSidebar>
  )
}

export default SideBar
