import { LineSeriesOption, ScatterSeriesOption } from "echarts/charts"
import { DatasetOption, ECBasicOption } from "echarts/types/dist/shared"
import { deepEqual } from "fast-equals"
import { debounceTime, distinctUntilChanged, forkJoin, map, Observable, of, switchMap } from "rxjs"
import { Ref, shallowRef } from "vue"
import { DataQueryResult } from "../DataQueryExecutor"
import { PersistentStateManager } from "../PersistentStateManager"
import { ChartConfigurator, collator, ValueUnit } from "../chart"
import {
  DataQuery,
  DataQueryConfigurator,
  DataQueryDimension,
  DataQueryExecutorConfiguration,
  DataQueryFilter,
  toMutableArray,
} from "../dataQuery"
import { LineChartOptions, ScatterChartOptions } from "../echarts"
import { durationAxisPointerFormatter, isDurationFormatterApplicable, nsToMs, numberAxisLabelFormatter } from "../formatter"
import { ServerConfigurator } from "./ServerConfigurator"
import { createComponentState, updateComponentState } from "./componentState"
import { configureQueryFilters, createFilterObservable, FilterConfigurator } from "./filter"
import { fromFetchWithRetryAndErrorHandling, refToObservable } from "./rxjs"

export type ChartType = "line" | "scatter"

export class MeasureConfigurator implements DataQueryConfigurator, ChartConfigurator {
  readonly data = shallowRef<Array<string>>([])
  private readonly _selected = shallowRef<Array<string> | string | null>(null)
  readonly state = createComponentState()

  createObservable(): Observable<unknown> {
    return refToObservable(this.selected, true)
  }

  setSelected(value: Array<string> | string | null) {
    this._selected.value = value
  }

  get selected(): Ref<Array<string> | null> {
    const ref = this._selected
    if (typeof ref.value === "string") {
      ref.value = [ref.value]
    }
    return ref as Ref<Array<string> | null>
  }

  constructor(serverConfigurator: ServerConfigurator,
              persistentStateManager: PersistentStateManager,
              filters: Array<FilterConfigurator> = [],
              readonly skipZeroValues: boolean = true,
              readonly chartType: ChartType = "line") {
    persistentStateManager.add("measure", this._selected)

    const isIj = serverConfigurator.db === "ij"

    createFilterObservable(serverConfigurator, filters)
      .pipe(
        debounceTime(100),
        distinctUntilChanged(deepEqual),
        switchMap(() => {
          const loadMeasureListUrl = getLoadMeasureListUrl(serverConfigurator, filters)
          if (loadMeasureListUrl == null) {
            return of(null)
          }

          this.state.loading = true
          if (isIj) {
            return forkJoin([
              fromFetchWithRetryAndErrorHandling<Array<string>>(`${serverConfigurator.serverUrl}/api/v1/meta/measure?db=${serverConfigurator.db}`),
              fromFetchWithRetryAndErrorHandling<Array<string>>(loadMeasureListUrl),
            ])
              .pipe(
                map(data => {
                  return data.flat(1)
                }),
              )
          }
          else {
            return fromFetchWithRetryAndErrorHandling<Array<string>>(loadMeasureListUrl)
          }
        }),
        updateComponentState(this.state),
      )
      .subscribe(data => {
        if (data == null) {
          return
        }

        if (isIj) {
          data = Array.from(new Set(data.map(it => /^c\.i\.ide\.[a-zA-Z]\.[a-zA-Z] preloading$/.test(it) ? "com.intellij.ide.misc.EvaluationSupport" : it)))
        }

        const selectedRef = this.selected
        this.data.value = data
        const selected = selectedRef.value
        if (selected != null && selected.length !== 0) {
          // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
          const filtered = selected.filter(it => data!.includes(it))
          if (filtered.length !== selected.length) {
            selectedRef.value = filtered
          }
        }
      })
  }

  configureQuery(query: DataQuery, configuration: DataQueryExecutorConfiguration): boolean {
    const measureNames = toMutableArray(this.selected.value)
    if (measureNames.length === 0) {
      return false
    }

    configureQuery(measureNames, query, configuration, this.skipZeroValues)
    configuration.measures = measureNames
    configuration.chartConfigurator = this
    return true
  }

  configureChart(data: DataQueryResult, configuration: DataQueryExecutorConfiguration): ECBasicOption {
    return configureChart(configuration, data, this.chartType)
  }
}

function getLoadMeasureListUrl(serverConfigurator: ServerConfigurator, filters: Array<FilterConfigurator>): string | null {
  const query = new DataQuery()
  const configuration = new DataQueryExecutorConfiguration()
  if (!serverConfigurator.configureQuery(query, configuration)) {
    return null
  }

  if (!configureQueryFilters(query, filters)) {
    return null
  }

  let fieldPrefix: string
  if (serverConfigurator.db === "ij") {
    fieldPrefix = "measure"
  }
  else {
    fieldPrefix = serverConfigurator.table === "measure" ? "" : "measures"
  }

  // "group by" is equivalent of distinct (https://clickhouse.tech/docs/en/sql-reference/statements/select/distinct/#alternatives)
  query.addDimension(fieldPrefix.length === 0 ? {n: "name"} : {n: fieldPrefix, subName: "name"})
  query.order = fieldPrefix.length === 0 ? "name" :`${fieldPrefix}.name`
  query.table = serverConfigurator.table ?? "report"
  query.flat = true
  return serverConfigurator.computeQueryUrl(query)
}

export class PredefinedMeasureConfigurator implements DataQueryConfigurator, ChartConfigurator {
  constructor(private readonly measures: Array<string>,
              readonly skipZeroValues: Ref<boolean> = shallowRef(true),
              private readonly chartType: ChartType = "line",
              private readonly valueUnit: ValueUnit = "ms") {
  }

  createObservable(): Observable<unknown> {
    return refToObservable(this.skipZeroValues)
  }

  configureQuery(query: DataQuery, configuration: DataQueryExecutorConfiguration): boolean {
    configureQuery(this.measures, query, configuration, this.skipZeroValues.value)
    configuration.chartConfigurator = this
    configuration.measures = this.measures
    return true
  }

  configureChart(data: DataQueryResult, configuration: DataQueryExecutorConfiguration): ECBasicOption {
    return configureChart(configuration, data, this.chartType, this.valueUnit)
  }
}

export function measureNameToLabel(key: string): string {
  const metricPathEndDotIndex = key.indexOf(".")
  if (metricPathEndDotIndex == -1) {
    // remove _d or _i suffix
    return key.replace(/_[a-z]$/g, "")
  }
  else {
    return key
  }
}

function configureQuery(measureNames: Array<string>, query: DataQuery, configuration: DataQueryExecutorConfiguration, skipZeroValues: boolean): void {
  // stable order of series (UI) and fields in query (caching)
  measureNames.sort((a, b) => collator.compare(a, b))

  query.insertField({
    n: "t",
    sql: "toUnixTimestamp(generated_time)*1000",
  }, 0)

  // we cannot request several measures in one SQL query - for each measure separate SQl query with filter by measure name
  const isIj = query.db === "ij"
  const structureName = isIj ? "measure" : "measures"
  const valueName = isIj ? "duration" : "value"
  const field: DataQueryDimension = {n: ""}
  query.insertField(field, 1)

  const prevFilters: Array<DataQueryFilter> = []

  const addFilter = (filter: DataQueryFilter): void => {
    prevFilters.push(filter)
    query.addFilter(filter)
  }

  configuration.queryProducers.push({
    size(): number {
      return measureNames.length
    },
    mutate(index: number): void {
      const measure = measureNames[index]

      delete field.sql
      delete field.subName

      if (prevFilters.length > 0) {
        query.removeFilters(prevFilters)
        prevFilters.length = 0
      }

      let valueFieldName: string
      if (query.table === "measure") {
        field.n = "value"
        field.resultKey = measure.replaceAll(".", "_")
        addFilter({f: "name", v: measure})
        valueFieldName = "value"
      }
      else if (isIj && !measure.includes(" ")) {
        field.n = measure
        valueFieldName = measure
      }
      else {
        if (measure.endsWith(".end")) {
          field.n = structureName
          field.subName = "end"
          field.sql = `(${structureName}.start+${structureName}.${valueName})`
        }
        else {
          field.n = structureName
          field.subName = valueName
        }

        addFilter({f: `${structureName}.name`, v: measure.endsWith(".end") ? measure.substring(0, measure.length - ".end".length) : measure})
        valueFieldName = `${structureName}.${valueName}`
      }

      if (skipZeroValues) {
        addFilter({f: valueFieldName, o: "!=", v: 0})
      }
    },
    getSeriesName(index: number): string {
      return measureNames.length > 1 ? measureNameToLabel(measureNames[index]) : ""
    },
    getMeasureName(index: number): string {
      return measureNames[index]
    }
  })

  if (query.order != null) {
    throw new Error("order must be configured only by MetricLoader")
  }
  query.order = "t"
}

function configureChart(configuration: DataQueryExecutorConfiguration,
                        dataList: DataQueryResult,
                        chartType: ChartType,
                        valueUnit: ValueUnit = "ms"): LineChartOptions | ScatterChartOptions {
  const series = new Array<LineSeriesOption | ScatterSeriesOption>()
  let useDurationFormatter = true

  const dataset: Array<DatasetOption> = []

  for (let dataIndex = 0, n = dataList.length; dataIndex < n; dataIndex++) {
    const measureName = configuration.measureNames[dataIndex]
    const seriesName = configuration.seriesNames[dataIndex]
    const seriesData = dataList[dataIndex]
    series.push({
      // formatter is detected by measure name - that's why series id is specified (see usages of seriesId)
      id: measureName === seriesName ? seriesName : `${measureName}@${seriesName}`,
      name: seriesName,
      type: chartType,
      showSymbol: seriesData[0].length < 100,
      // 10 is a default value for scatter (undefined doesn't work to unset)
      symbolSize: chartType === "line" ? Math.min(800 / seriesData[0].length, 9) : 10,
      symbol: "circle",
      legendHoverLink: true,
      // applicable only for line chart
      sampling: "lttb",
      seriesLayoutBy: "row",
      datasetIndex: dataIndex,
      dimensions: [{name: "time", type: "time"}, {name: seriesName, type: "int"}],
    })

    if (useDurationFormatter && !isDurationFormatterApplicable(measureName)) {
      useDurationFormatter = false
    }

    dataset.push({
      source: seriesData,
      sourceHeader: false,
    })
  }

  // if (chartType == "scatter") {
  //   dataset.push({
  //     fromDatasetIndex: 0,
  //     transform: {
  //       type: "ecStat:regression",
  //       // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  //       // @ts-ignore
  //       config: {
  //         method: "polynomial",
  //         dimensions: [1],
  //       },
  //     },
  //   })
  //   series.push({
  //     silent: true,
  //     type: "line",
  //     smooth: true,
  //     datasetIndex: dataset.length - 1,
  //     // symbolSize: 0.1,
  //     // symbol: "circle",
  //     // label: {show: false, fontSize: 16},
  //     // labelLayout: {dx: -20},
  //     // encode: {label: 2, tooltip: 1},
  //   })
  // }

  const isNs = valueUnit == "ns"
  const valueInMsFormatter = useDurationFormatter ? durationAxisPointerFormatter : numberAxisLabelFormatter
  const formatter: (valueInMs: number) => string = isNs ? v => valueInMsFormatter(nsToMs(v)) : valueInMsFormatter
  return {
    dataset,
    yAxis: {
      axisLabel: {
        formatter,
      },
      axisPointer: {
        label: {
          formatter(data): string {
            return formatter(data["value"] as number)
          },
        },
      },
    },
    series: series as LineSeriesOption,
  }
}