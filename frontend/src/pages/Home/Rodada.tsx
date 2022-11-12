import { EuiPageTemplate, EuiText } from '@elastic/eui'

const Rodada = ({ title }: { title: string }) => {
  return (
    <EuiPageTemplate.Section grow bottomBorder='extended'>
      <EuiText>
        <h2>{title}</h2>
      </EuiText>
    </EuiPageTemplate.Section>
  )
}

export default Rodada
