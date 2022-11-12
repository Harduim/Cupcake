import { EuiButton, EuiErrorBoundary, EuiPageTemplate } from '@elastic/eui'
import { useEffect } from 'react'

const Login = () => {
  useEffect(() => {
    document.title = `Cupcake | Login`
  }, [])

  return (
    <EuiErrorBoundary>
      <EuiPageTemplate panelled={true} bottomBorder={'extended'} grow>
        <EuiPageTemplate.EmptyPrompt
          title={<span>Faça login con sua conta Microsoft</span>}
          footer={
            <EuiButton color='primary' fill onClick={() => {}}>
              Login Corporativo
            </EuiButton>
          }
        />
      </EuiPageTemplate>
    </EuiErrorBoundary>
  )
}

export default Login
