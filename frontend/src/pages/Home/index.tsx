import { useContext } from 'react'
import PageLayout from '../../components/PageLayout'
import GlobalContext from '../../context/GlobalContext'
import FaseGrupos from './FaseGrupos'
import Rodada from './Rodada'

const Home = () => {
  const { isLoading, matches, me } = useContext(GlobalContext)
  return (
    <PageLayout title='Página Inicial' isLoading={isLoading}>
      {matches?.map((m, i) => (
        <div key={i}>
          {m.date}
          {m.nationalTeamAId}
        </div>
      ))}
      <FaseGrupos />
      <Rodada title='Oitavas' />
      <Rodada title='Rodada Coringa' />
      <Rodada title='Quartas' />
      <Rodada title='Semifinal' />
      <Rodada title='Disputa pelo 3º lugar' />
      <Rodada title='Final' />
    </PageLayout>
  )
}

export default Home
