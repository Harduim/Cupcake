import {
  EuiCard,
  EuiFlexGrid,
  EuiFlexItem,
  EuiIcon,
  EuiPanel,
  EuiSpacer,
  EuiTitle,
} from '@elastic/eui'
import { useContext, useState } from 'react'
import { Match } from '../../clients'
import MatchForm from '../../components/MatchForm'
import GlobalContext from '../../context/GlobalContext'

const BonusStage = () => {
  const { teams, isLoading } = useContext(GlobalContext)

  const coringaMatch: Match = {
    id: 'coringa-match',
    name: 'Finais (Coringa)',
    date: '2022-12-03',
    nationalTeamAId: '',
    nationalTeamBId: '',
    golA: null,
    golB: null,
    bracketId: '',
    winnerId: '',
  }

  if (isLoading || !teams) return null
  return (
    <EuiFlexGrid columns={2} gutterSize='m'>
      <EuiPanel>
        <EuiTitle size='xs'>
          <h2>Tente Acertar a Final</h2>
        </EuiTitle>
        <EuiSpacer size='m' />
        <MatchForm
          match={coringaMatch}
          teams={teams}
          onSubmit={match => {
            console.log(match)
          }}
        />
      </EuiPanel>
    </EuiFlexGrid>
  )
}

export default BonusStage
