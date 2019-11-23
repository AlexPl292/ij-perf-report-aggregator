package model

type Report struct {
  Version string `json:"version"`

  Generated string `json:"generated"`
  Project   string `json:"project"`

  Build     string `json:"build"`
  BuildDate string `json:"buildDate"`

  Os          string `json:"os"`
  ProductCode string `json:"productCode"`
  Runtime     string `json:"runtime"`

  // not used yet
  TraceEvents []TraceEvent `json:"traceEvents"`

  MainActivities           []Activity `json:"items"`
  PrepareAppInitActivities []Activity `json:"prepareAppInitActivities"`

  TotalDurationActual int `json:"totalDurationActual"`
}

type ReportInfo struct {
  Report *Report

  RawData       []byte
  GeneratedTime int64
  ExtraData     ExtraData
}

type ExtraData struct {
  LastGeneratedTime int64
  BuildTime         int64

  ProductCode string
  BuildNumber string

  Machine string

  TcBuildId          int
  TcInstallerBuildId int
  TcBuildProperties  []byte
  Changes            [][]byte
}

type TraceEvent struct {
  Name  string `json:"name"`
  Phase string `json:"ph"`
  // in microseconds
  Timestamp int `json:"ts"`

  // in old reports (v10) can be int instead of string
  //Thread   string `json:"tid"`
  Category string `json:"cat"`
}

type Activity struct {
  Name   string `json:"name"`
  Thread string `json:"thread"`

  // in milliseconds
  Start    int `json:"start"`
  End      int `json:"end"`
  Duration int `json:"duration"`
}

// computed metrics
type DurationEventMetrics struct {
  Bootstrap               int
  AppInitPreparation      int
  AppInit                 int
  PluginDescriptorLoading int

  AppComponentCreation     int
  ProjectComponentCreation int

  // project post-startup dumb-aware activities
  ProjectDumbAware int

  ModuleLoading   int
  EditorRestoring int
}

type InstantEventMetrics struct {
  // value - not duration, but start, because it is instant event and not duration event
  Splash int

  StartUpCompleted int
}
