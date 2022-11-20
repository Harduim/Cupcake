import { EuiButton, EuiFieldNumber, EuiSelect, EuiSpacer, EuiTitle } from '@elastic/eui'
import { useState } from 'react'
import { Match, NationalTeam } from '../clients'
import { HOURS_BEFORE_MATCH_IN_MILLISECONDS } from '../utils/constants'
import { dateToBrDateTimeString, timeStringToBrDateTimeString } from '../utils/datetime'

interface IPropType {
  match: Match
  teams: NationalTeam[]
  onSubmit: (match: Match) => void
}

const MatchForm = ({ match, teams, onSubmit }: IPropType) => {
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
        onChange={(e: any) => handleChange('nationalTeamAId', e.target.value)}
      />
      <EuiSpacer size='s' />
      <EuiSelect
        hasNoInitialSelection
        prepend='Seleção B'
        options={options}
        value={_match.nationalTeamBId || ''}
        onChange={(e: any) => handleChange('nationalTeamBId', e.target.value)}
      />
      <EuiSpacer size='s' />
      <EuiFieldNumber
        prepend='Gols A'
        placeholder='Gols Seleção A'
        value={_match.golA || ''}
        onChange={(e: any) => handleChange('golA', parseInt(e.target.value))}
      />
      <EuiSpacer size='s' />
      <EuiFieldNumber
        prepend='Gols B'
        placeholder='Gols Seleção B'
        value={_match.golB || ''}
        onChange={(e: any) => handleChange('golB', parseInt(e.target.value))}
      />
      <EuiSpacer size='s' />
      <EuiSelect
        hasNoInitialSelection
        prepend='Vencedor'
        options={options}
        value={_match.winnerId || ''}
        onChange={(e: any) => handleChange('winnerId', e.target.value)}
      />
      <EuiSpacer size='m' />
      <EuiButton color='primary' onClick={() => onSubmit(_match)} fill>
        Salvar
      </EuiButton>
    </>
  )
}

export default MatchForm
