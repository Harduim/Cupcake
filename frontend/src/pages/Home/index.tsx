import {
  EuiFlexGrid,
  EuiFlexItem,
  EuiPageSection,
  EuiPanel,
  EuiSelectable,
  EuiSpacer,
  EuiTitle,
} from '@elastic/eui'
import { useContext } from 'react'
import PageLayout from '../../components/PageLayout'
import GlobalContext from '../../context/GlobalContext'
import { timeStringToBrDateString } from '../../utils/datetime'
import BetForm from './BetForm'

const Home = () => {
  const { brackets, bets, teamMap, isLoading } = useContext(GlobalContext)

  if (isLoading) {
    return (
      <PageLayout title='Bolão Copa do Mundo 2022' isLoading>
        {' '}
      </PageLayout>
    )
  }
  return (
    <PageLayout title='Bolão Copa do Mundo 2022' isLoading={isLoading}>
      <EuiPageSection color='subdued'>
        <EuiTitle>
          <h1>Fase de Grupos</h1>
        </EuiTitle>
        <EuiSpacer size='m' />
        <EuiPanel>
          <EuiSelectable
            aria-label='Searchable example'
            options={[{ label: 'Alpha' }, { label: 'B', searchableLabel: 'Beta' }]}
            onChange={newOptions => {}}
            searchable
            searchProps={{
              'data-test-subj': 'dataTestSubj',
            }}
          >
            {(list, search) => (
              <>
                {search}
                {list}
              </>
            )}
          </EuiSelectable>
        </EuiPanel>
      </EuiPageSection>
      <EuiSpacer size='m' />
      {brackets
        .sort((a, b) => Date.parse(a.openDate) - Date.parse(b.openDate))
        .map(b => {
          const isClosed = Date.parse(b.openDate) - Date.parse(b.closeDate) > 0
          return (
            <>
              <EuiPageSection key={b.id} color='subdued'>
                <EuiTitle>
                  <h1>{b.name}</h1>
                </EuiTitle>
                <p>
                  {isClosed
                    ? 'Encerrou ' + timeStringToBrDateString(b.closeDate)
                    : 'Inicia ' + timeStringToBrDateString(b.openDate)}
                </p>
                <EuiSpacer size='m' />
                <EuiFlexGrid columns={b.Matches.length > 3 ? 3 : 2} gutterSize='m'>
                  {b.Matches.sort((a, b) => Date.parse(a.date) - Date.parse(b.date)).map(m => (
                    <EuiFlexItem key={m.id}>
                      <BetForm match={m} teamMap={teamMap} bets={bets} />
                    </EuiFlexItem>
                  ))}
                </EuiFlexGrid>
              </EuiPageSection>
              <EuiSpacer size='m' />
            </>
          )
        })}
    </PageLayout>
  )
}

export default Home
