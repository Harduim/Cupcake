import { createContext, ReactNode } from 'react'

type IGlobalContext = {
  isLoading: boolean
}

export type { IGlobalContext }

const globalContextDefaults: IGlobalContext = {
  isLoading: true,
}

interface IContextProps {
  children: ReactNode
}

const GlobalContext = createContext<IGlobalContext>(globalContextDefaults)

export const GlobalProvider = ({ children }: IContextProps) => {
  const isLoading = false

  const provides = {
    isLoading,
  }

  return <GlobalContext.Provider value={provides}>{children}</GlobalContext.Provider>
}

export default GlobalContext
