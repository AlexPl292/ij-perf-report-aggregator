<template>
  <DashboardPage
    db-name="perfint"
    table="idea"
    persistent-id="idea_incremental_build_dashboard"
    initial-machine="Linux EC2 C6id.8xlarge (32 vCPU Xeon, 64 GB)"
    :charts="charts"
  >
    <section>
      <GroupProjectsChart
        v-for="chart in charts"
        :key="chart.definition.label"
        :label="chart.definition.label"
        :measure="chart.definition.measure"
        :projects="chart.projects"
      />
    </section>
  </DashboardPage>
  >
</template>

<script setup lang="ts">
import { ChartDefinition, combineCharts } from "../charts/DashboardCharts"
import GroupProjectsChart from "../charts/GroupProjectsChart.vue"
import DashboardPage from "../common/DashboardPage.vue"

const chartsDeclaration: ChartDefinition[] = [
  {
    labels: ["Java Design Patterns Rebuild/Build time", "Mappings Dir Size in Bytes"],
    measures: ["build_compilation_duration", "mappingsDirSizeInBytes"],
    projects: [
      "incremental-build-java/build_incremental",
      "incremental-build-java/rebuild_initial",
      "incremental-build-java-with-dep-graph/build_incremental",
      "incremental-build-java-with-dep-graph/rebuild_initial",
    ],
  },
  {
    labels: ["IntelliJ Rebuild/Build time", "Mappings Dir Size in Bytes"],
    measures: ["build_compilation_duration", "mappingsDirSizeInBytes"],
    projects: [
      "incremental-build-intellij/build_incremental",
      "incremental-build-intellij/rebuild_initial",
      "incremental-build-intellij-with-dep-graph/build_incremental",
      "incremental-build-intellij-with-dep-graph/rebuild_initial",
    ],
  },
  {
    labels: ["Youtrack JPS and Gradle Rebuild/Build time", "Mappings Dir Size in Bytes"],
    measures: ["build_compilation_duration", "mappingsDirSizeInBytes"],
    projects: [
      "incremental-build-youtrack-jps/build_incremental",
      "incremental-build-youtrack-jps/rebuild_initial",
      "incremental-build-youtrack-jps-with-dep-graph/build_incremental",
      "incremental-build-youtrack-jps-with-dep-graph/rebuild_initial",
      "incremental-build-youtrack-gradle/build_incremental",
      "incremental-build-youtrack-gradle/rebuild_initial",
    ],
  },
]

const charts = combineCharts(chartsDeclaration)
</script>
