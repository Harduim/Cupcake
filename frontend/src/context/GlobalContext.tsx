import { useQuery } from '@tanstack/react-query'
import { createContext, ReactNode } from 'react'
import { User } from '../clients'
import { Match } from '../clients/models/Match'
import api from '../services/api'

type IGlobalContext = {
  isLoading: boolean
  me: User
  matches: Match[]
}

export type { IGlobalContext }

const globalContextDefaults: IGlobalContext = {
  isLoading: true,
  me: {} as User,
  matches: [] as Match[],
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

  const isLoading = meIsLoading || matchesIsLoading

  const provides = {
    isLoading,
    matches,
    me,
  }

  return <GlobalContext.Provider value={provides}>{children}</GlobalContext.Provider>
}

export default GlobalContext
