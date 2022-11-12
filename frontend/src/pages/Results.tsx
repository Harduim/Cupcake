import { EuiBasicTable, EuiPageTemplate, EuiText } from '@elastic/eui'
import PageLayout from '../components/PageLayout'

const Results = () => {
  return (
    <PageLayout title='Apuração' isLoading={false}>
      <EuiPageTemplate.Section grow bottomBorder='extended'>
        <EuiText>
          <h2>Apuração</h2>
        </EuiText>
      </EuiPageTemplate.Section>

      <EuiBasicTable
        tableCaption='Tabela de apuração'
        items={[]}
        rowHeader='rowHeader'
        columns={[]}
        // rowProps={getRowProps}
        // cellProps={getCellProps}
      />
    </PageLayout>
  )
}

export default Results
