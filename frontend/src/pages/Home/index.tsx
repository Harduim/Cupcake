import {
  EuiAvatar,
  EuiButton,
  EuiCard,
  EuiCheckableCard,
  EuiFieldNumber,
  EuiFlexGrid,
  EuiFlexGroup,
  EuiFlexItem,
  EuiIcon,
  EuiPageSection,
  EuiPanel,
  EuiSelect,
  EuiSpacer,
  EuiText,
  EuiTitle,
} from '@elastic/eui'
import { useContext, useState } from 'react'
import { Bet, Match, NationalTeam } from '../../clients'
import PageLayout from '../../components/PageLayout'
import GlobalContext from '../../context/GlobalContext'
import { HOURS_BEFORE_MATCH_IN_MILLISECONDS } from '../../utils/constants'
import {
  dateToBrDateTimeString,
  timeStringToBrDateString,
  timeStringToBrDateTimeString,
} from '../../utils/datetime'

const { PUBLIC_URL } = process.env

const updateBet = (bets: Bet) => {
  console.log(bets)
}

interface IBetProps {
  match: Match
  teamMap: Map<string, NationalTeam>
  bets: Bet[]
}

const BetForm = ({ match, teamMap, bets }: IBetProps) => {
  const [_bet, setBet] = useState<Bet>()

  const handleChange = (prop: string, value: string | number) => {
    const newBet = { ..._bet, [prop]: value }
    // setBet(newBet)
  }
  const matchClose = new Date(Date.parse(match.date) - HOURS_BEFORE_MATCH_IN_MILLISECONDS)
  const nationalTeamA = teamMap.get(match.nationalTeamAId)
  const nationalTeamB = teamMap.get(match.nationalTeamBId)
  const isDisabled =
    !match.nationalTeamAId || !match.nationalTeamBId || matchClose.getTime() < new Date().getTime()

  const iconNotDefined = <EuiAvatar size='xl' name='N' />
  const iconTeamA = <EuiIcon size='xxl' type={`${PUBLIC_URL}/flags/${nationalTeamA?.name}.svg`} />
  const iconTeamB = <EuiIcon size='xxl' type={`${PUBLIC_URL}/flags/${nationalTeamB?.name}.svg`} />

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
            onClick={() => {}}
            display={undefined}
          />
        </EuiFlexItem>
        <EuiFlexItem>
          <EuiCard
            icon={!nationalTeamA ? iconNotDefined : iconTeamB}
            title={nationalTeamB?.name.replace('_', ' ') || 'Não Definido'}
            isDisabled={isDisabled}
            onClick={() => {}}
            display={'success'}
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
          value={match.golA || ''}
          disabled={isDisabled}
          onChange={(e: any) => {
            handleChange('golA', parseInt(e.target.value))
          }}
        />
        <EuiFieldNumber
          placeholder=' '
          value={match.golB || ''}
          disabled={isDisabled}
          onChange={(e: any) => {
            handleChange('golB', parseInt(e.target.value))
          }}
        />
      </EuiFlexGroup>

      <EuiSpacer size='s' />
      <EuiSpacer size='m' />
      <EuiButton
        color='primary'
        disabled={isDisabled}
        onClick={() => {
          // updateBet(_match)
        }}
        fill
      >
        Salvar
      </EuiButton>
      <EuiSpacer size='m' />
      <EuiPanel color='success' paddingSize='s'>
        <EuiText size='s'>Palpite Registrado</EuiText>
      </EuiPanel>
      <EuiPanel color='warning' paddingSize='s'>
        <EuiText size='s'>Você só tem mais N dias</EuiText>
      </EuiPanel>
    </EuiPanel>
  )
}

const Home = () => {
  const { matches, brackets, bets, teamMap, isLoading } = useContext(GlobalContext)

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
              <EuiFlexGrid columns={3} gutterSize='m'>
                {matches
                  .filter(m => m.bracketId === b.id)
                  .sort((a, b) => Date.parse(a.date) - Date.parse(b.date))
                  .map(m => (
                    <EuiFlexItem key={m.id}>
                      <BetForm match={m} teamMap={teamMap} bets={bets} />
                    </EuiFlexItem>
                  ))}
              </EuiFlexGrid>
            </EuiPageSection>
          )
        })}
    </PageLayout>
  )
}

export default Home
