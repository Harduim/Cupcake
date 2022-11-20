import {
  EuiCard,
  EuiFlexGrid,
  EuiFlexItem,
  EuiIcon,
  EuiPanel,
  EuiSpacer,
  EuiTitle,
} from '@elastic/eui'
import { useContext, useState } from 'react'
import GlobalContext from '../../context/GlobalContext'
const { PUBLIC_URL } = process.env
const MAX_SELECTIONS = 16

const GroupStage = () => {
  const { teams, isLoading } = useContext(GlobalContext)

  const [selected, setSelected] = useState<string[]>([])

  const handleSelect = (newSelection: string) => {
    if (!selected.includes(newSelection) && selected.length < MAX_SELECTIONS) {
      return setSelected([...selected, newSelection])
    }
    setSelected(selected.filter(s => s !== newSelection))
  }

  if (isLoading) return null
  return (
    <EuiPanel>
      <EuiTitle size='xs'>
        <h2>
          Escolha as seleções que vão se classificar - {selected.length} de {MAX_SELECTIONS}{' '}
          escolhidas
        </h2>
      </EuiTitle>
      <EuiSpacer size='m' />
      <EuiFlexGrid columns={4} gutterSize='m'>
        {teams.map(t => (
          <EuiFlexItem key={t.id}>
            <EuiCard
              icon={<EuiIcon size='xxl' type={`${PUBLIC_URL}/flags/${t?.name}.svg`} />}
              title={t?.name.replaceAll('_', ' ') || 'Não Definido'}
              isDisabled={false}
              onClick={() => handleSelect(t.id)}
              display={selected.includes(t.id) ? 'success' : 'subdued'}
            />
          </EuiFlexItem>
        ))}
      </EuiFlexGrid>
      <EuiSpacer size='m' />
      <p>
        {selected.length} de {MAX_SELECTIONS} Seleções Escolhidas
      </p>
    </EuiPanel>
  )
}

export default GroupStage
