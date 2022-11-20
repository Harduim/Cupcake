import {
  EuiFlexGrid,
  EuiPanel,
  EuiSpacer,
  EuiTitle,
} from '@elastic/eui'
import { useQuery } from '@tanstack/react-query'
import { useContext } from 'react'
import { JokerMatch, Match } from '../../clients'
import MatchForm from '../../components/MatchForm'
import GlobalContext from '../../context/GlobalContext'
import api from '../../services/api'

const handleUpdate = async (coringaMatch: Match) => {
  console.log(coringaMatch)
  try {
    await api.put('joker', {
      id: coringaMatch.id,
      golA: coringaMatch.golA,
      golB: coringaMatch.golB,
      nationalTeamAId: coringaMatch.nationalTeamAId,
      nationalTeamBId: coringaMatch.nationalTeamBId,
      winnerId: coringaMatch.winnerId,
    })
  } catch (error) {
    console.error(error)
    return
  }
}

const queryOptions = {
  refetchOnWindowFocus: true,
  retry: false,
  staleTime: 1000 * 60 * 5,
}

const BonusStage = () => {
  const { teams, isLoading } = useContext(GlobalContext)

  const { isLoading: jokerIsLoading, data: joker } = useQuery({
    queryKey: ['joker'],
    queryFn: async () => {
      const response = await api.get('joker')
      return response.data as JokerMatch
    },
    ...queryOptions,
  })

  if (isLoading || jokerIsLoading || !teams || !joker) return null

  const jokerMatch: Match = {
    id: joker.id,
    name: 'Finais (Coringa)',
    date: '2022-12-03',
    nationalTeamAId: joker.nationalTeamAId,
    nationalTeamBId: joker.nationalTeamBId,
    golA: joker.golA,
    golB: joker.golB,
    bracketId: '',
    winnerId: joker.winnerId,
  }

  return (
    <EuiFlexGrid columns={2} gutterSize='m'>
      <EuiPanel>
        <EuiTitle size='xs'>
          <h2>Tente Acertar a Final</h2>
        </EuiTitle>
        <EuiSpacer size='m' />
        <MatchForm match={jokerMatch} teams={teams} onSubmit={handleUpdate} />
      </EuiPanel>
    </EuiFlexGrid>
  )
}

export default BonusStage
