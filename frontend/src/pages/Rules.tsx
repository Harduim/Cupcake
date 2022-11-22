import { EuiText } from '@elastic/eui'
import PageLayout from '../components/PageLayout'

const Rules = () => {
  return (
    <PageLayout title='Regras' isLoading={false}>
      <EuiText grow>
        <h1>Fases</h1>
        <h2>Fase de grupos</h2>
        <ul>
          <li>Cada usuário escolhe 16 times que vão se classificar para as oitavas de final;</li>
          <li>
            Os usuários tem até a <strong>25/11/2022</strong> para fazer a aposta;
          </li>
          <li>Cada acerto vale 3 pontos.</li>
        </ul>
        <h2>Rodada Coringa</h2>
        <ul>
          <li>
            Cada usuário escolhe:
            <ul>
              <li>Seleção campeã;</li>
              <li>Seleção vice campeã;</li>
              <li>Placar da final.</li>
            </ul>
          </li>
          <li>
            Os acertos valem:
            <ul>
              <li>
                Gols seleção campeã: 21 (Pênaltis <strong>NÃO</strong> entram na conta);
              </li>
              <li>
                Gols vice campeã: 21 (Pênaltis <strong>NÃO</strong> entram na conta);
              </li>
              <li>Vencedor: 42.</li>
            </ul>
          </li>
          <li>
            Os usuários tem até a data de inicio das <strong>Oitavas de final</strong> para fazer a
            aposta.
          </li>
        </ul>
        <h2>Oitavas / Quartas / Semifinal / Disputa do 3° Lugar / Final</h2>
        <ul>
          <li>
            Para cada partida da chave, cada usuário escolhe:
            <ul>
              <li>
                Gols time A (<strong>NÃO</strong> considerar pênaltis);
              </li>
              <li>
                Gols time B (<strong>NÃO</strong> considerar pênaltis);
              </li>
              <li>Vencedor.</li>
            </ul>
          </li>
          <li>
            O usuário tem até <strong>6 horas</strong> antes do jogo para fazer a aposta;
          </li>
          <li>
            Os acertos valem:
            <ul>
              <li>
                Gols time A: 1 ponto. (Pênaltis <strong>NÃO</strong> entram na conta);{' '}
              </li>
              <li>
                Gols time B: 1 ponto. (Pênaltis <strong>NÃO</strong> entram na conta);
              </li>
              <li>
                Placar completo (Ambos Gols time A e B): 2 pontos. (Pênaltis <strong>NÃO </strong>
                entram na conta);
              </li>
              <li>Vencedor da partida: 2 pontos;</li>
              <li>Placar completo + Vencedor: 3 pontos.</li>
            </ul>
          </li>
          <li>
            Os pontos serão multiplicados conforme a chave:
            <ul>
              <li>Oitavas: 1;</li>
              <li>Quartas: 2;</li>
              <li>Semifinal: 3;</li>
              <li>Disputa do 3° Lugar: 3;</li>
              <li>Finais: 8.</li>
            </ul>
          </li>
        </ul>
        <h2>Apuração dos pontos</h2>
        <ul>
          <li>
            A apuração dos pontos, disponibilizada através da página "Resultados", será feita
            diariamente ou a cada dois dias;
          </li>
          <li>
            Os resultados das partidas e a contabilização dos pontos <strong>NÃO</strong> será feia
            em tempo real;
          </li>
          <li>Um relatório detalhado de apuração será disponibilizado no final do evento.</li>
        </ul>
      </EuiText>
    </PageLayout>
  )
}

export default Rules
