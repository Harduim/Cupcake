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
import { useQuery } from '@tanstack/react-query'
import { useContext, useEffect, useState } from 'react'
import GlobalContext from '../../context/GlobalContext'
import api from '../../services/api'
const { PUBLIC_URL } = process.env
const MAX_SELECTIONS = 16

const handleSubmit = async (teams: (string | null)[], userId: string) => {
  await api.put('groups', {
    groups: teams.map(t => {
      return { national_team_id: t, user_id: userId }
    }),
  })
}

interface IGroup {
  user_id: string
  national_team_id: string | null
}

const queryOptions = {
  refetchOnWindowFocus: true,
  retry: false,
  staleTime: 1000 * 60 * 5,
}

const GroupStage = () => {
  const { teams, isLoading, me } = useContext(GlobalContext)

  const { isLoading: selectedIsLoading, data: selected } = useQuery({
    queryKey: ['selectedForGroups'],
    queryFn: async () => {
      const response = await api.get('groups')
      return response.data as IGroup[]
    },
    ...queryOptions,
  })

  const [_selected, _setSelected] = useState<(string | null)[]>([])

  useEffect(() => {
    if (isLoading || !selected) return
    _setSelected(selected.map(s => s.national_team_id))
  }, [selected, isLoading])

  const handleSelect = (newSelection: string) => {
    if (!_selected.includes(newSelection) && _selected.length < MAX_SELECTIONS) {
      return _setSelected([..._selected, newSelection])
    }
    _setSelected(_selected.filter(s => s !== newSelection))
  }
  const isSendEnabled = _selected.length === MAX_SELECTIONS

  if (isLoading || selectedIsLoading) return null
  return (
    <EuiPanel>
      <EuiTitle size='xs'>
        <h2>
          Escolha as seleções que vão se classificar - {_selected.length} de {MAX_SELECTIONS}{' '}
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
              display={_selected.includes(t.id) ? 'success' : 'subdued'}
            />
          </EuiFlexItem>
        ))}
      </EuiFlexGrid>
      <EuiSpacer size='m' />
      <p>
        {_selected.length} de {MAX_SELECTIONS} Seleções Escolhidas
      </p>
      <EuiSpacer size='m' />
      <EuiButton
        color='primary'
        onClick={() => handleSubmit(_selected, me.id)}
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
