import {
  EuiAvatar,
  EuiButton,
  EuiCard,
  EuiFieldNumber,
  EuiFlexGrid,
  EuiFlexGroup,
  EuiFlexItem,
  EuiIcon,
  EuiPageSection,
  EuiPanel,
  EuiSpacer,
  EuiText,
  EuiTitle,
} from '@elastic/eui'
import { useContext, useState } from 'react'
import uuid from 'react-uuid'
import { Bet, Match, NationalTeam } from '../../clients'
import PageLayout from '../../components/PageLayout'
import GlobalContext from '../../context/GlobalContext'
import api, { queryClient } from '../../services/api'
import { HOURS_BEFORE_MATCH_IN_MILLISECONDS } from '../../utils/constants'
import {
  dateToBrDateTimeString,
  timeStringToBrDateString,
  timeStringToBrDateTimeString,
} from '../../utils/datetime'

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
    if (value === undefined) return
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
            title={nationalTeamA?.name.replace('_', ' ') || 'Não Definido'}
            isDisabled={isDisabled}
            onClick={() => handleChange('winnerId', nationalTeamA?.id)}
            display={displayTeamA}
          />
        </EuiFlexItem>
        <EuiFlexItem>
          <EuiCard
            icon={!nationalTeamA ? iconNotDefined : iconTeamB}
            title={nationalTeamB?.name.replace('_', ' ') || 'Não Definido'}
            isDisabled={isDisabled}
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
        disabled={isSendDisabled}
        onClick={() => {
          updateBet(_bet, _bet.id === bet?.id)
        }}
        fill
      >
        Salvar
      </EuiButton>
      <EuiSpacer size='m' />
    </EuiPanel>
  )
}

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
