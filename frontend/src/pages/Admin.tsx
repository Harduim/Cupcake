import {
  EuiFlexGroup,
  EuiFlexItem,
  EuiPageSection,
  EuiPanel,
  EuiSpacer,
  EuiTitle,
} from '@elastic/eui'
import { useContext } from 'react'
import { Match } from '../clients'
import MatchForm from '../components/MatchForm'
import PageLayout from '../components/PageLayout'
import GlobalContext from '../context/GlobalContext'
import api, { queryClient } from '../services/api'

const updateMatch = async (match: Match) => {
  console.log(match)
  try {
    await api.put('matches', match)
  } catch (error) {
    console.error(error)
    return
  }
  queryClient.invalidateQueries(['matches'])
}

const Admin = () => {
  const { matches, brackets, teams, isLoading } = useContext(GlobalContext)

  if (isLoading || !matches || !brackets || !teams) {
    return (
      <PageLayout title='Administração' isLoading>
        {' '}
      </PageLayout>
    )
  }
  return (
    <PageLayout title='Administração' isLoading={isLoading}>
      {brackets
        .sort((a, b) => Date.parse(a.openDate) - Date.parse(b.openDate))
        .map(b => {
          return (
            <>
              <EuiPageSection key={b.id} color='subdued'>
                <EuiTitle>
                  <h1>{b.name}</h1>
                </EuiTitle>
                <EuiSpacer size='s' />
                <EuiFlexGroup gutterSize='m' wrap>
                  {b.Matches.sort((a, b) => Date.parse(a.date) - Date.parse(b.date)).map(m => (
                    <EuiFlexItem grow={false} key={m.id}>
                      <EuiPanel grow={false}>
                        <MatchForm isAdminForm match={m} teams={teams} onSubmit={updateMatch} />
                      </EuiPanel>
                    </EuiFlexItem>
                  ))}
                </EuiFlexGroup>
              </EuiPageSection>
              <EuiSpacer size='m' />
            </>
          )
        })}
    </PageLayout>
  )
}

export default Admin
