import {
  EuiAvatar,
  EuiButton,
  EuiCallOut,
  EuiCard,
  EuiFieldNumber,
  EuiFlexGroup,
  EuiFlexItem,
  EuiIcon,
  EuiPanel,
  EuiSpacer,
  EuiTitle,
} from '@elastic/eui'
import { useState } from 'react'
import uuid from 'react-uuid'
import { Bet, Match, NationalTeam } from '../../clients'
import api, { queryClient } from '../../services/api'
import { HOURS_BEFORE_MATCH_IN_MILLISECONDS } from '../../utils/constants'
import { dateToBrDateTimeString, timeStringToBrDateTimeString } from '../../utils/datetime'

const { PUBLIC_URL } = process.env

const updateBet = async (bet: Bet, update: boolean) => {
  console.log(update ? 'update' : 'create', bet)
  try {
    if (update) {
      await api.put('bets', bet)
    } else {
      await api.post('bets', bet)
    }
  } catch (error) {
    console.error(error)
    return
  }
  queryClient.invalidateQueries(['bets'])
}

interface IBetProps {
  match: Match
  teamMap: Map<string, NationalTeam>
  bets?: Map<string, Bet>
}

const BetForm = ({ match, teamMap, bets }: IBetProps) => {
  const bet = bets?.get(match.id)
  let defaultBet: Bet = bet || { id: uuid(), matchId: match.id }
  const [_bet, setBet] = useState<Bet>(defaultBet)

  const handleChange = (prop: string, value: string | number | undefined) => {
    if (value === undefined || isDisabled) return
    const newBet = { ..._bet, [prop]: value }
    setBet(newBet)
  }
  const matchClose = new Date(Date.parse(match.date) - HOURS_BEFORE_MATCH_IN_MILLISECONDS)
  const nationalTeamA = teamMap.get(match.nationalTeamAId)
  const nationalTeamB = teamMap.get(match.nationalTeamBId)
  const isDisabled =
    !match.nationalTeamAId || !match.nationalTeamBId || matchClose.getTime() < new Date().getTime()

  const iconNotDefined = <EuiAvatar size='xl' name='N' />
  const iconTeamA = <EuiIcon size='xxl' type={`${PUBLIC_URL}/flags/${nationalTeamA?.name}.svg`} />
  const iconTeamB = <EuiIcon size='xxl' type={`${PUBLIC_URL}/flags/${nationalTeamB?.name}.svg`} />

  const getDisplay = (teamId?: string, winnerId?: string) => {
    let display: 'success' | 'accent' | 'subdued'
    if (!winnerId || !teamId) {
      display = 'subdued'
    } else if (_bet.winnerId === teamId) {
      display = 'success'
    } else {
      display = 'accent'
    }
    return display
  }

  const displayTeamA = getDisplay(match.nationalTeamAId, _bet?.winnerId)
  const displayTeamB = getDisplay(match.nationalTeamBId, _bet?.winnerId)

  const isSendDisabled =
    _bet.golA === undefined || _bet.golB === undefined || _bet.winnerId === undefined

  return (
    <EuiPanel>
      <EuiTitle size='xs'>
        <h2>{match.name}</h2>
      </EuiTitle>
      <EuiSpacer size='s' />
      <p>
        <b>Horário do jogo:</b> {timeStringToBrDateTimeString(match.date)}
      </p>
      <p>
        <b>Encerramento:</b> &nbsp;&nbsp; {dateToBrDateTimeString(matchClose)}
      </p>
      <EuiSpacer size='m' />
      <EuiFlexGroup justifyContent='spaceAround'>
        <p>
          <b>Escolha o Vencedor</b>
        </p>
      </EuiFlexGroup>
      <EuiSpacer size='s' />
      <EuiFlexGroup justifyContent='spaceAround'>
        <EuiFlexItem>
          <EuiCard
            icon={!nationalTeamA ? iconNotDefined : iconTeamA}
            title={nationalTeamA?.name.replaceAll('_', ' ') || 'Não Definido'}
            onClick={() => handleChange('winnerId', nationalTeamA?.id)}
            display={displayTeamA}
          />
        </EuiFlexItem>
        <EuiFlexItem>
          <EuiCard
            icon={!nationalTeamA ? iconNotDefined : iconTeamB}
            title={nationalTeamB?.name.replaceAll('_', ' ') || 'Não Definido'}
            onClick={() => handleChange('winnerId', nationalTeamB?.id)}
            display={displayTeamB}
          />
        </EuiFlexItem>
      </EuiFlexGroup>
      <EuiSpacer size='s' />
      <EuiFlexGroup justifyContent='spaceAround'>
        <p>
          <b>Defina o Placar</b>
        </p>
      </EuiFlexGroup>
      <EuiSpacer size='s' />
      <EuiFlexGroup>
        <EuiFieldNumber
          placeholder=' '
          value={_bet.golA === undefined ? '' : _bet.golA}
          disabled={isDisabled}
          onChange={(e: any) => {
            const gols = parseInt(e.target.value)
            if (gols < 0 || Number.isNaN(gols)) return
            handleChange('golA', gols)
          }}
        />
        <EuiFieldNumber
          placeholder=' '
          value={_bet.golB === undefined ? '' : _bet.golB}
          disabled={isDisabled}
          onChange={(e: any) => {
            const gols = parseInt(e.target.value)
            if (gols < 0 || Number.isNaN(gols)) return
            handleChange('golB', gols)
          }}
        />
      </EuiFlexGroup>
      <EuiSpacer size='s' />
      <EuiSpacer size='m' />
      <EuiButton
        color='primary'
        disabled={isSendDisabled || isDisabled}
        onClick={() => {
          updateBet(_bet, _bet.id === bet?.id)
        }}
        fill
      >
        Salvar
      </EuiButton>
      {isSendDisabled && !isDisabled && (
        <>
          <EuiSpacer size='m' />
          <EuiCallOut
            title='Escolha o vencedor e defina o placar para habilitar salvamento'
            color='warning'
            iconType='help'
          />
        </>
      )}
      <EuiSpacer size='m' />
    </EuiPanel>
  )
}

export default BetForm
