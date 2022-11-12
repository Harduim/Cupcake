import { EuiProvider } from '@elastic/eui'
import '@elastic/eui/dist/eui_theme_light.css'
import { QueryClientProvider } from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { BrowserRouter } from 'react-router-dom'
import './App.css'
import AppRouter from './Routes'
import { queryClient } from './services/api'

const { REACT_APP_ENVIRONMENT } = process.env

const App = () => {
  return (
    <EuiProvider colorMode='light'>
      <QueryClientProvider client={queryClient}>
        <BrowserRouter>
          <AppRouter />
          {REACT_APP_ENVIRONMENT === 'DEV' && <ReactQueryDevtools />}
        </BrowserRouter>
      </QueryClientProvider>
    </EuiProvider>
  )
}

export default App
