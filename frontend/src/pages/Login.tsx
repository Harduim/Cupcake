import { EuiButton, EuiErrorBoundary, EuiPageTemplate } from '@elastic/eui'
import { useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import { login } from '../services/auth'

const { REACT_APP_API_URL } = process.env
const REDIRECT_URL = `${REACT_APP_API_URL}/auth/sso`

const Login = () => {
  useEffect(() => {
    document.title = `Cupcake | Login`
  }, [])

  const { search } = useLocation()
  const jwt_string = new URLSearchParams(search).get('tkn')
  if (jwt_string) {
    login(jwt_string)
    window.location.href = '/home'
  }
  const handleLoginClick = (e: any) => {
    e.preventDefault()
    window.location.href = REDIRECT_URL
  }

  return (
    <EuiErrorBoundary>
      <EuiPageTemplate panelled={true} bottomBorder={'extended'} grow>
        <EuiPageTemplate.EmptyPrompt
          title={<span>Faça login con sua conta Microsoft</span>}
          footer={
            <EuiButton color='primary' fill onClick={handleLoginClick}>
              Login Corporativo
            </EuiButton>
          }
        />
      </EuiPageTemplate>
    </EuiErrorBoundary>
  )
}

export default Login
