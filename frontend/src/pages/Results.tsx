import { EuiBasicTable, EuiPageTemplate, EuiSpacer } from '@elastic/eui'
import { useQuery } from '@tanstack/react-query'
import { useContext } from 'react'
import { User } from '../clients'
import PageLayout from '../components/PageLayout'
import GlobalContext from '../context/GlobalContext'
import api from '../services/api'

const queryOptions = {
  refetchOnWindowFocus: true,
  retry: false,
  staleTime: 1000 * 60 * 5,
}

const Results = () => {
  const { isLoading: globalIsLoading, me } = useContext(GlobalContext)

  const { isLoading: usersIsLoading, data: users } = useQuery({
    queryKey: ['users'],
    queryFn: () => api.get('users').then(r => r.data),
    enabled: !!me,
    ...queryOptions,
  })

  const isLoading = globalIsLoading || usersIsLoading

  if (isLoading || !me) {
    return (
      <PageLayout title='Apuração' isLoading>
        {' '}
      </PageLayout>
    )
  }

  return (
    <PageLayout title='Resultados' isLoading={isLoading}>
      <EuiPageTemplate.Section grow bottomBorder='extended'>
        <EuiBasicTable
          tableCaption='Tabela de apuração'
          columns={[
            { field: 'name', name: 'Nome', sortable: true },
            { field: 'points', name: 'Pontuação', sortable: true },
          ]}
          items={users.sort((a: User, b: User) => b.points - a.points)}
        />
        <EuiSpacer size='xl' />
      </EuiPageTemplate.Section>
    </PageLayout>
  )
}

export default Results
