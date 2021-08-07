import useTileBuild from '@/composables/useTileBuild'
import TileStatus from '@/enums/tileStatus'
import TileType from '@/enums/tileType'
import TileBuild from '@/types/tileBuild'
import TileConfig from '@/types/tileConfig'
import TileState from '@/types/tileState'
import {RootState} from '@/store'
import {computed} from 'vue'
import {useStore} from 'vuex'

export default function useTileCommons(config: TileConfig) {
  const store = useStore<RootState>()

  const type = computed((): TileType => {
    return config.type
  })

  const stateKey = computed((): string => {
    return config.stateKey
  })

  const theme = computed((): string => {
    return store.getters.theme.toString().toLowerCase()
  })

  const state = computed((): TileState | undefined => {
    if (!Object.keys(store.state.tilesState).includes(stateKey.value)) {
      return
    }

    return store.state.tilesState[stateKey.value]
  })

  const label = computed((): string | undefined => {
    if (config.label) {
      if (config.label === '-') {
        return
      }

      return config.label
    }

    if (state.value === undefined) {
      return
    }

    return state.value.label
  })

  const build = computed((): TileBuild | undefined => {
    if (state.value === undefined) {
      return
    }

    return state.value.build
  })

  const status = computed((): TileStatus | undefined => {
    if (state.value === undefined) {
      return
    }

    return state.value.status
  })

  const url = computed((): string | undefined => {
    if (build.value === undefined) {
      return
    }

    return build.value.url
  })

  const previousStatus = computed((): string | undefined => {
    if (build.value === undefined) {
      return
    }

    return build.value.previousStatus
  })

  const isQueued = computed((): boolean => {
    return status.value === TileStatus.Queued
  })

  const isRunning = computed((): boolean => {
    return status.value === TileStatus.Running
  })

  const isSucceeded = computed((): boolean => {
    if (isQueued.value || isRunning.value) {
      return previousStatus.value === TileStatus.Success
    }

    return status.value === TileStatus.Success
  })

  const isFailed = computed((): boolean => {
    if (isQueued.value || isRunning.value) {
      return previousStatus.value === TileStatus.Failed
    }

    return status.value === TileStatus.Failed
  })

  const isWarning = computed((): boolean => {
    if (isQueued.value || isRunning.value) {
      return previousStatus.value === TileStatus.Warning
    }

    return status.value === TileStatus.Warning
  })

  return {
    type,
    url,
    theme,
    state,
    label,
    build,
    status,
    previousStatus,
    isQueued,
    isRunning,
    isSucceeded,
    isFailed,
    isWarning,
    ...useTileBuild(type, status, build),
  }
}
