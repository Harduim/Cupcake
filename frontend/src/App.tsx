import { EuiProvider } from '@elastic/eui'
import '@elastic/eui/dist/eui_theme_light.css'
import './App.css'
import Home from './pages/Home'

const App = () => {
  return (
    <EuiProvider colorMode='light'>
      <Home />
    </EuiProvider>
  )
}

export default App
