import {
  EuiFlexGrid,
  EuiFlexItem,
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

const GroupFase = () => {
  const { teams, isLoading } = useContext(GlobalContext)

  const defaultOpts = teams.map(t => {
    return { label: t.name.replace('_', ' '), searchableLabel: t.name }
  })
  const [selected, setSelected] = useState<EuiSelectableOption[]>(defaultOpts)

  return (
    <EuiPanel>
      <EuiSelectable
        aria-label='fase-grupos-selectable'
        options={selected}
        onChange={newOptions => setSelected(newOptions)}
        searchable
        isLoading={isLoading}
        searchProps={{
          'data-test-subj': 'dataTestSubj',
        }}
      >
        {(list, search) => (
          <>
            {search}
            {list}
          </>
        )}
      </EuiSelectable>
    </EuiPanel>
  )
}

export default GroupFase
