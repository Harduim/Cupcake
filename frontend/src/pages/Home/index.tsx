import PageLayout from '../../components/PageLayout'
import FaseGrupos from './FaseGrupos'
import Rodada from './Rodada'

const Home = () => {
  return (
    <PageLayout title='Página inicial' isLoading={false}>
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
