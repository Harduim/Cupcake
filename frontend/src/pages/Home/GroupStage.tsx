import {
  EuiButton,
  EuiCallOut,
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
import api from '../../services/api'
const { PUBLIC_URL } = process.env
const MAX_SELECTIONS = 16

const handleSubmit = async (teams: string[]) => {
  await api.put('groups', teams)
}

const GroupStage = () => {
  const { teams, isLoading } = useContext(GlobalContext)

  const [selected, setSelected] = useState<string[]>([])

  const handleSelect = (newSelection: string) => {
    if (!selected.includes(newSelection) && selected.length < MAX_SELECTIONS) {
      return setSelected([...selected, newSelection])
    }
    setSelected(selected.filter(s => s !== newSelection))
  }
  const isSendEnabled = selected.length === MAX_SELECTIONS

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
      <EuiSpacer size='m' />
      <EuiButton
        color='primary'
        onClick={() => handleSubmit(selected)}
        fill
        isDisabled={!isSendEnabled}
      >
        Salvar
      </EuiButton>
      {!isSendEnabled && (
        <>
          <EuiSpacer size='m' />
          <EuiCallOut
            title='Escolha 16 seleções para habilitar salvamento'
            color='warning'
            iconType='help'
          />
        </>
      )}
    </EuiPanel>
  )
}

export default GroupStage
