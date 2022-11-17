import { useQuery } from '@tanstack/react-query'
import { createContext, ReactNode } from 'react'
import { User } from '../clients'
import api from '../services/api'

type IGlobalContext = {
  isLoading: boolean
  me: User
}

export type { IGlobalContext }

const globalContextDefaults: IGlobalContext = {
  isLoading: true,
  me: {} as User,
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

  const isLoading = meIsLoading

  const provides = {
    isLoading,
    me,
  }

  return <GlobalContext.Provider value={provides}>{children}</GlobalContext.Provider>
}

export default GlobalContext
