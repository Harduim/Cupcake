import {
  EuiButton,
  EuiErrorBoundary,
  EuiPageSidebar,
  EuiPageTemplate,
  EuiSideNav,
  EuiText,
  htmlIdGenerator,
} from '@elastic/eui'
import { useEffect } from 'react'
import PageLoading from './PageLoading'

const LogoutButton = () => {
  return (
    <EuiButton color={'primary'} onClick={() => {}}>
      Sair{' '}
    </EuiButton>
  )
}

const SideBar = ({ title }: { title: string }) => {
  const sideNav = [
    {
      name: 'Cupcake',
      id: htmlIdGenerator('Cupcake')(),
      items: [
        {
          name: 'Página Inicial',
          id: htmlIdGenerator('Página Inicial')(),
          onClick: () => {},
          isSelected: false,
        },
        {
          name: 'Resultados',
          id: htmlIdGenerator('Resultados')(),
          href: '/#/navigation/side-nav',
          isSelected: false,
        },
        {
          name: 'Regras',
          id: htmlIdGenerator('Regras')(),
          onClick: () => {},
          isSelected: true,
        },
        {
          name: 'Administração',
          id: htmlIdGenerator('Administração')(),
          disabled: false,
          isSelected: false,
        },
      ],
    },
  ]

  return (
    <EuiPageSidebar paddingSize='m' sticky>
      <EuiSideNav
        aria-label='sidebar-nav'
        mobileTitle={title}
        toggleOpenOnMobile={() => {}}
        isOpenOnMobile={false}
        items={sideNav}
      />
    </EuiPageSidebar>
  )
}

interface IPageProps {
  isLoading: boolean
  title: string
  children: React.ReactNode
}

const PageLayout = ({ title, children, isLoading }: IPageProps) => {
  useEffect(() => {
    document.title = `Cupcake | ${title}`
  }, [])

  return (
    <EuiErrorBoundary>
      <EuiPageTemplate panelled={true} bottomBorder={'extended'} grow offset={1}>
        <EuiPageTemplate.Sidebar sticky>
          <SideBar title={title} />
        </EuiPageTemplate.Sidebar>
        <EuiPageTemplate.Header rightSideItems={[<LogoutButton key='LogoutButton' />]}>
          <EuiText textAlign='center'>
            <strong>{title}</strong>
          </EuiText>
        </EuiPageTemplate.Header>
        <EuiPageTemplate.Section grow bottomBorder='extended'>
          {isLoading ? <PageLoading /> : <EuiErrorBoundary>{children}</EuiErrorBoundary>}
        </EuiPageTemplate.Section>
      </EuiPageTemplate>
    </EuiErrorBoundary>
  )
}

export default PageLayout
