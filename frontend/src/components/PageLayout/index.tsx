import { EuiButton, EuiErrorBoundary, EuiPageTemplate, EuiText } from '@elastic/eui'
import { useEffect } from 'react'
import PageLoading from './PageLoading'
import SideBar from './Sidebar'

const LogoutButton = () => {
  return (
    <EuiButton color={'primary'} onClick={() => {}}>
      Sair{' '}
    </EuiButton>
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
        <EuiPageTemplate.Header
          rightSideItems={[<LogoutButton key='LogoutButton' />]}
          pageTitle={title}
        />
        <EuiPageTemplate.Section grow bottomBorder='extended'>
          {isLoading ? <PageLoading /> : <EuiErrorBoundary>{children}</EuiErrorBoundary>}
        </EuiPageTemplate.Section>
      </EuiPageTemplate>
    </EuiErrorBoundary>
  )
}

export default PageLayout
