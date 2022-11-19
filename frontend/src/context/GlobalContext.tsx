import { useQuery } from '@tanstack/react-query'
import { createContext, ReactNode } from 'react'
import { Bet, NationalTeam, User } from '../clients'
import { Bracket } from '../clients/models/Bracket'
import { Match } from '../clients/models/Match'
import api from '../services/api'

type IGlobalContext = {
  isLoading: boolean
  me: User
  matches: Match[]
  brackets: Bracket[]
  teams: NationalTeam[]
  teamMap: Map<string, NationalTeam>
  bets: Bet[]
}

export type { IGlobalContext }

const globalContextDefaults: IGlobalContext = {
  isLoading: true,
  me: {} as User,
  matches: [] as Match[],
  brackets: [] as Bracket[],
  teams: [] as NationalTeam[],
  bets: [] as Bet[],
  teamMap: new Map(),
}

interface IContextProps {
  children: ReactNode
}

const queryOptions = {
  refetchOnWindowFocus: true,
  retry: false,
  staleTime: 1000 * 60 * 5,
}

const GlobalContext = createContext<IGlobalContext>(globalContextDefaults)

export const GlobalProvider = ({ children }: IContextProps) => {
  const { isLoading: meIsLoading, data: me } = useQuery({
    queryKey: ['me'],
    queryFn: () => api.get('users/me').then(r => r.data),
    ...queryOptions,
  })
  const { isLoading: matchesIsLoading, data: matches } = useQuery({
    queryKey: ['matches'],
    queryFn: () => api.get('matches').then(r => r.data),
    enabled: !!me,
    ...queryOptions,
  })
  const { isLoading: bracketsIsLoading, data: brackets } = useQuery({
    queryKey: ['brackets'],
    queryFn: () => api.get('brackets').then(r => r.data),
    enabled: !!me,
    ...queryOptions,
  })
  const { isLoading: teamsIsLoading, data: teams } = useQuery({
    queryKey: ['teams'],
    queryFn: () => api.get('national-teams').then(r => r.data),
    enabled: !!me,
    ...queryOptions,
  })
  const { isLoading: betsIsLoading, data: bets } = useQuery({
    queryKey: ['bets'],
    queryFn: () => api.get('bets').then(r => r.data),
    enabled: !!me,
    ...queryOptions,
  })

  const teamMap = new Map()

  const isLoading = [
    meIsLoading,
    matchesIsLoading,
    bracketsIsLoading,
    teamsIsLoading,
    betsIsLoading,
    !matches,
    !brackets,
    !teams,
  ].some(l => l)

  if (!isLoading && teams.length > 0) {
    teams.forEach((t: NationalTeam) => teamMap.set(t.id, t))
  }

  const provides = {
    isLoading,
    matches,
    brackets,
    teams,
    teamMap,
    bets,
    me,
  }

  return <GlobalContext.Provider value={provides}>{children}</GlobalContext.Provider>
}

export default GlobalContext
