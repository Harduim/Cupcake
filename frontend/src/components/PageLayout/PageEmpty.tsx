import { EuiButton, EuiEmptyPrompt } from '@elastic/eui'
import { useNavigate } from 'react-router-dom'

const PageEmpty = () => {
  const navigate = useNavigate()

  return (
    <EuiEmptyPrompt
      iconType='logoSecurity'
      title={<h1>Página Não Encontrada</h1>}
      body={<p>Que tal voltar para a página inicial?</p>}
      actions={
        <EuiButton color='primary' fill onClick={() => navigate('/')}>
          Página Inicial
        </EuiButton>
      }
    />
  )
}

export default PageEmpty
