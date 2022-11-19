import {
  EuiButton,
  EuiFieldNumber,
  EuiFlexGroup,
  EuiFlexItem,
  EuiPageSection,
  EuiPanel,
  EuiSelect,
  EuiSpacer,
  EuiTitle,
} from '@elastic/eui'
import { useContext, useState } from 'react'
import { Match, NationalTeam } from '../clients'
import PageLayout from '../components/PageLayout'
import GlobalContext from '../context/GlobalContext'
import api, { queryClient } from '../services/api'
import { HOURS_BEFORE_MATCH_IN_MILLISECONDS } from '../utils/constants'
import { dateToBrDateTimeString, timeStringToBrDateTimeString } from '../utils/datetime'

const updateMatch = async (match: Match) => {
  console.log(match)
  try {
    await api.put('matches', match)
  } catch (error) {
    console.error(error)
  }
  queryClient.invalidateQueries(['matches'])
}

const MatchForm = ({ match, teams }: { match: Match; teams: NationalTeam[] }) => {
  const [_match, setMatch] = useState<Match>(match)
  const options = teams.map(t => {
    return { text: t.name, value: t.id }
  })

  const handleChange = (prop: string, value: string | number) => {
    const newMatch = { ..._match, [prop]: value }
    setMatch(newMatch)
  }
  const matchClose = new Date(Date.parse(match.date) - HOURS_BEFORE_MATCH_IN_MILLISECONDS)

  return (
    <>
      <EuiTitle size='xs'>
        <h2>{match.name}</h2>
      </EuiTitle>
      <EuiSpacer size='s' />
      <p>
        <b>Horário do jogo:</b> {timeStringToBrDateTimeString(match.date)}
      </p>
      <p>
        <b>Encerramento:</b> &nbsp;&nbsp;{dateToBrDateTimeString(matchClose)}
      </p>

      <EuiSpacer size='s' />
      <EuiSelect
        hasNoInitialSelection
        prepend='Seleção A'
        options={options}
        value={_match.nationalTeamAId || ''}
        onChange={(e: any) => {
          handleChange('nationalTeamAId', e.target.value)
        }}
      />
      <EuiSpacer size='s' />
      <EuiSelect
        hasNoInitialSelection
        prepend='Seleção B'
        options={options}
        value={_match.nationalTeamBId || ''}
        onChange={(e: any) => {
          handleChange('nationalTeamBId', e.target.value)
        }}
      />
      <EuiSpacer size='s' />
      <EuiFieldNumber
        prepend='Gols A'
        placeholder='Gols Seleção A'
        value={_match.golA || ''}
        onChange={(e: any) => {
          handleChange('golA', parseInt(e.target.value))
        }}
      />
      <EuiSpacer size='s' />
      <EuiFieldNumber
        prepend='Gols B'
        placeholder='Gols Seleção B'
        value={_match.golB || ''}
        onChange={(e: any) => {
          handleChange('golB', parseInt(e.target.value))
        }}
      />
      <EuiSpacer size='s' />
      <EuiSelect
        hasNoInitialSelection
        prepend='Vencedor'
        // TODO: Filter by team_a and team_b
        options={options}
        value={_match.winnerId || ''}
        onChange={(e: any) => {
          handleChange('winnerId', e.target.value)
        }}
      />
      <EuiSpacer size='m' />
      <EuiButton
        color='primary'
        onClick={() => {
          updateMatch(_match)
        }}
        fill
      >
        Salvar
      </EuiButton>
    </>
  )
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
            <EuiPageSection key={b.id}>
              <EuiTitle>
                <h1>{b.name}</h1>
              </EuiTitle>
              <EuiSpacer size='s' />
              <EuiFlexGroup gutterSize='m' wrap>
                {matches
                  .filter(m => m.bracketId === b.id)
                  .sort((a, b) => Date.parse(a.date) - Date.parse(b.date))
                  .map(m => (
                    <EuiFlexItem grow={false} key={m.id}>
                      <EuiPanel grow={false}>
                        <MatchForm match={m} teams={teams} />
                      </EuiPanel>
                    </EuiFlexItem>
                  ))}
              </EuiFlexGroup>
            </EuiPageSection>
          )
        })}
    </PageLayout>
  )
}

export default Admin
