import { EuiButton, EuiEmptyPrompt, EuiLink, EuiTitle } from '@elastic/eui'

const PageEmpty = () => {
  return (
    <EuiEmptyPrompt
      iconType='logoSecurity'
      title={<h1>Página não encontrada</h1>}
      body={<p>Add a new case or change your filter settings.</p>}
      actions={
        <EuiButton color='primary' fill>
          Add a case
        </EuiButton>
      }
      footer={
        <>
          <EuiTitle size='xxs'>
            <h3>Want to learn more?</h3>
          </EuiTitle>
          <EuiLink href='#' target='_blank'>
            Read the docs
          </EuiLink>
        </>
      }
    />
  )
}

export default PageEmpty
