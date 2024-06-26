import { Observable } from "rxjs"
import { ref, watch } from "vue"
import { DataQuery, DataQueryConfigurator, DataQueryExecutorConfiguration } from "../../common/dataQuery"
import { useSettingsStore } from "../settingsStore"
import { median } from "../../../shared/changeDetector/statistic"
import { FilterConfigurator } from "../../../configurators/filter"
import { refToObservable } from "../../../configurators/rxjs"

export class ScalingConfigurator implements DataQueryConfigurator, FilterConfigurator {
  private settingsStore = useSettingsStore()
  readonly value = ref(this.settingsStore.scaling)

  constructor() {
    watch(
      () => this.settingsStore.scaling,
      (newValue) => {
        this.value.value = newValue
      }
    )
  }

  createObservable(): Observable<unknown> {
    return refToObservable(this.value)
  }

  configureFilter(_: DataQuery): boolean {
    return true
  }

  configureQuery(_: DataQuery, _configuration: DataQueryExecutorConfiguration): boolean {
    return true
  }
}

export function scaleToMedian(arr: number[]): number[] {
  if (arr.length === 0) {
    return arr
  }
  const medianValue = median(arr)
  return medianValue === 0 ? arr : arr.map((value) => (value / medianValue) * 50)
}
