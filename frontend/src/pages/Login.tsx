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

const Login = () => {
  useEffect(() => {
    document.title = `Cupcake | Login`
  }, [])

  return (
    <EuiErrorBoundary>
      <EuiPageTemplate panelled={true} bottomBorder={'extended'} grow offset={1}>
        <EuiPageTemplate.Section grow bottomBorder='extended'>
          {' Login form'}
        </EuiPageTemplate.Section>
      </EuiPageTemplate>
    </EuiErrorBoundary>
  )
}

export default Login
