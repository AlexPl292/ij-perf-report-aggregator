package server

import (
  "context"
  "database/sql"
  "github.com/jmoiron/sqlx"
  "github.com/pkg/errors"
  "github.com/valyala/quicktemplate"
  "io"
  "math"
  "net/http"
  "report-aggregator/pkg/model"
  "report-aggregator/pkg/util"
  "strconv"
  "strings"
  "time"
)

func (t *StatsServer) handleMetricsRequest(request *http.Request) ([]byte, error) {
  query, err := parseQuery(request)
  if err != nil {
    return nil, err
  }

  product, machines, eventType, err := getProductAndMachine(query)
  if err != nil {
    return nil, err
  }

  buffer := byteBufferPool.Get()
  defer byteBufferPool.Put(buffer)
  err = t.computeMetricsResponse(product, machines, eventType, buffer, request.Context())
  if err != nil {
    return nil, err
  }
  return CopyBuffer(buffer), nil
}

func (t *StatsServer) computeMetricsResponse(product string, machines []string, eventType rune, writer io.Writer, context context.Context) error {
  var metricNames []string
  if eventType == 'd' {
    metricNames = model.DurationMetricNames
  } else {
    metricNames = model.InstantMetricNames
  }

  rows, err := t.selectData(metricNames, eventType, product, machines, context)
  if err != nil {
    return err
  }

  defer util.Close(rows, t.logger)

  templateWriter := quicktemplate.AcquireWriter(writer)
  defer quicktemplate.ReleaseWriter(templateWriter)
  jsonWriter := templateWriter.N()
  jsonWriter.S("[")

  offset := 4

  columnPointers := make([]interface{}, offset+len(metricNames))
  for i := range columnPointers {
    columnPointers[i] = new(interface{})
  }

  var sb strings.Builder

  isFirst := true
  lastBuildWithoutUniqueSuffix := ""
  for rows.Next() {
    err := rows.Scan(columnPointers...)
    if err != nil {
      return err
    }

    err = rows.Scan(columnPointers...)
    if err != nil {
      return errors.WithStack(err)
    }

    generatedTime := ((*(columnPointers[0].(*interface{}))).(time.Time)).Unix()

    if isFirst {
      isFirst = false
    } else {
      jsonWriter.S(",")
    }

    // timestamp
    jsonWriter.S(`{"t":`)
    // seconds to milliseconds
    jsonWriter.D(int(generatedTime * 1000))
    jsonWriter.S(",")

    // build number with addition if not unique (build as x coordinate)
    sb.Reset()
    sb.WriteString(strconv.Itoa(int((*(columnPointers[1].(*interface{}))).(uint8))))
    sb.WriteRune('.')
    sb.WriteString(strconv.Itoa(int((*(columnPointers[2].(*interface{}))).(uint16))))
    buildC3 := int((*(columnPointers[3].(*interface{}))).(uint16))
    if buildC3 != 0 {
      sb.WriteRune('.')
      sb.WriteString(strconv.Itoa(buildC3))
    }

    buildAsString := sb.String()
    // https://www.amcharts.com/docs/v4/tutorials/handling-repeating-categories-on-category-axis/
    if lastBuildWithoutUniqueSuffix == buildAsString {
      // not unique - add time
      sb.WriteRune(' ')
      sb.WriteRune('(')
      sb.WriteString(strconv.FormatInt(generatedTime, 10))
      sb.WriteRune(')')
      buildAsString = sb.String()
    } else {
      lastBuildWithoutUniqueSuffix = buildAsString
    }

    jsonWriter.S(`"build":`)
    jsonWriter.Q(buildAsString)

    for index, name := range metricNames {
      jsonWriter.S(`,"`)
      jsonWriter.S(name)
      jsonWriter.S(`":`)

      switch untypedValue := (*(columnPointers[index+offset].(*interface{}))).(type) {
      case float64:
        jsonWriter.F(math.Round(untypedValue))
      case float32:
        jsonWriter.F(float64(untypedValue))
      case int32:
        jsonWriter.D(int(untypedValue))
      case uint8:
        jsonWriter.D(int(untypedValue))
      case uint16:
        jsonWriter.D(int(untypedValue))
      default:
        return errors.Errorf("unknown type: %v", untypedValue)
      }
    }

    jsonWriter.S("}")
  }

  jsonWriter.S("]")

  return nil
}

func (t *StatsServer) selectData(metricNames []string, eventType rune, product string, machines []string, context context.Context) (*sql.Rows, error) {
  var sb strings.Builder
  sb.WriteString("select generated_time, build_c1, build_c2, build_c3")
  for _, name := range metricNames {
    sb.WriteString(", ")
    sb.WriteString(name)
    sb.WriteRune('_')
    sb.WriteRune(eventType)
  }

  sb.WriteString(" from report where product = ? and machine in (?) order by build_c1, build_c2, build_c3, generated_time")

  query, queryArgs, err := sqlx.In(sb.String(), product, machines)
  if err != nil {
    return nil, errors.WithStack(err)
  }
  return t.db.QueryContext(context, query, queryArgs...)
}
