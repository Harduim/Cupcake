import { EuiButton, EuiCallOut, EuiFieldNumber, EuiSelect, EuiSpacer, EuiTitle } from '@elastic/eui'
import { useState } from 'react'
import { Match, NationalTeam } from '../clients'
import { HOURS_BEFORE_MATCH_IN_MILLISECONDS } from '../utils/constants'
import { dateToBrDateTimeString, timeStringToBrDateTimeString } from '../utils/datetime'

interface IPropType {
  match: Match
  teams: NationalTeam[]
  onSubmit: (match: Match) => void
  isAdminForm?: boolean
}

const MatchForm = ({ match, teams, onSubmit, isAdminForm }: IPropType) => {
  const [_match, setMatch] = useState<Match>(match)
  const options = teams.map(t => {
    return { text: t.name, value: t.id }
  })
  const winnerOptions = options.filter(o =>
    [_match.nationalTeamBId, _match.nationalTeamAId].includes(o.value)
  )
  const handleChange = (prop: string, value: string | number) => {
    const newMatch = { ..._match, [prop]: value }
    setMatch(newMatch)
  }

  const matchClose = new Date(Date.parse(match.date) - HOURS_BEFORE_MATCH_IN_MILLISECONDS)
  const isSendDisabled =
    !_match.nationalTeamAId ||
    !_match.nationalTeamBId ||
    _match.winnerId === '' ||
    _match.golA === undefined ||
    _match.golB === undefined ||
    _match.golA === null ||
    _match.golB === null ||
    _match.winnerId === undefined

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
        value={_match.golA === undefined || _match.golA === null ? '' : _match.golA}
        onChange={(e: any) => {
          const gols = parseInt(e.target.value)
          if (gols < 0 || Number.isNaN(gols)) return
          handleChange('golA', gols)
        }}
      />
      <EuiSpacer size='s' />
      <EuiFieldNumber
        prepend='Gols B'
        placeholder='Gols Seleção B'
        value={_match.golB === undefined || _match.golB === null ? '' : _match.golB}
        onChange={(e: any) => {
          const gols = parseInt(e.target.value)
          if (gols < 0 || Number.isNaN(gols)) return
          handleChange('golB', gols)
        }}
      />
      <EuiSpacer size='s' />
      <EuiSelect
        hasNoInitialSelection
        prepend='Vencedor'
        options={winnerOptions}
        value={_match.winnerId || ''}
        onChange={(e: any) => handleChange('winnerId', e.target.value)}
        disabled={winnerOptions.length === 0}
      />
      <EuiSpacer size='m' />
      <EuiButton
        color='primary'
        onClick={() => onSubmit(_match)}
        fill
        isDisabled={isSendDisabled && !isAdminForm}
      >
        Salvar
      </EuiButton>
      {isSendDisabled && !isAdminForm && (
        <>
          <EuiSpacer size='m' />
          <EuiCallOut
            title='Escolha o vencedor e defina o placar para habilitar salvamento'
            color='warning'
            iconType='help'
          />
        </>
      )}
    </>
  )
}

export default MatchForm
