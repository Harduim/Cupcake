import {
  EuiCard,
  EuiFlexGrid,
  EuiFlexGroup,
  EuiFlexItem,
  EuiIcon,
  EuiPageSection,
  EuiPanel,
  EuiSelectable,
  EuiSelectableOption,
  EuiSpacer,
  EuiTitle,
} from '@elastic/eui'
import { useContext, useState } from 'react'
import { NationalTeam } from '../../clients'
import GlobalContext from '../../context/GlobalContext'
const { PUBLIC_URL } = process.env

const GroupFase = () => {
  const { teams, isLoading } = useContext(GlobalContext)

  const [selected, setSelected] = useState<NationalTeam[]>()

  return (
    <EuiPanel>
      <EuiTitle size='xs'>
        <h2>Escolha as seleções que vão se classificar</h2>
      </EuiTitle>
      <EuiSpacer size='m' />
      <EuiFlexGrid columns={4} gutterSize='m'>
        {teams.map(t => (
          <EuiFlexItem key={t.id}>
            <EuiCard
              icon={<EuiIcon size='xxl' type={`${PUBLIC_URL}/flags/${t?.name}.svg`} />}
              title={t?.name.replaceAll('_', ' ') || 'Não Definido'}
              isDisabled={false}
              onClick={() => {}}
              display={'subdued'}
            />
          </EuiFlexItem>
        ))}
      </EuiFlexGrid>
    </EuiPanel>
  )
}

export default GroupFase
