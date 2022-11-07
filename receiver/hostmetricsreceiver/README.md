# Host Metrics Receiver

| Status                   |                   |
| ------------------------ | ----------------- |
| Stability                | [beta]            |
| Supported pipeline types | metrics           |
| Distributions            | [core], [contrib] |

The Host Metrics receiver generates metrics about the host system scraped
from various sources. This is intended to be used when the collector is
deployed as an agent.

## Getting Started

The collection interval and the categories of metrics to be scraped can be
configured:

```yaml
hostmetrics:
  collection_interval: <duration> # default = 1m
  scrapers:
    <scraper1>:
    <scraper2>:
    ...
```

The available scrapers are:

| Scraper      | Supported OSs                | Description                                            |
| ------------ | ---------------------------- | ------------------------------------------------------ |
| [cpu]        | All except Mac<sup>[1]</sup> | CPU utilization metrics                                |
| [disk]       | All except Mac<sup>[1]</sup> | Disk I/O metrics                                       |
| [load]       | All                          | CPU load metrics                                       |
| [filesystem] | All                          | File System utilization metrics                        |
| [memory]     | All                          | Memory utilization metrics                             |
| [network]    | All                          | Network interface I/O metrics & TCP connection metrics |
| [paging]     | All                          | Paging/Swap space utilization and I/O metrics          |
| [processes]  | Linux                        | Process count metrics                                  |
| [process]    | Linux & Windows              | Per process CPU, Memory, and Disk I/O metrics          |

[cpu]: ./internal/scraper/cpuscraper/documentation.md
[disk]: ./internal/scraper/diskscraper/documentation.md
[filesystem]: ./internal/scraper/filesystemscraper/documentation.md
[load]: ./internal/scraper/loadscraper/documentation.md
[memory]: ./internal/scraper/memoryscraper/documentation.md
[network]: ./internal/scraper/networkscraper/documentation.md
[paging]: ./internal/scraper/pagingscraper/documentation.md
[processes]: ./internal/scraper/processesscraper/documentation.md
[process]: ./internal/scraper/processscraper/documentation.md

### Notes

<sup>[1]</sup> Not supported on Mac when compiled without cgo which is the default.

Several scrapers support additional configuration:

### Disk

```yaml
disk:
  <include|exclude>:
    devices: [ <device name>, ... ]
    match_type: <strict|regexp>
```

### File System

```yaml
filesystem:
  <include_devices|exclude_devices>:
    devices: [ <device name>, ... ]
    match_type: <strict|regexp>
  <include_fs_types|exclude_fs_types>:
    fs_types: [ <filesystem type>, ... ]
    match_type: <strict|regexp>
  <include_mount_points|exclude_mount_points>:
    mount_points: [ <mount point>, ... ]
    match_type: <strict|regexp>
```

### Load

`cpu_average` specifies whether to divide the average load by the reported number of logical CPUs (default: `false`).

```yaml
load:
  cpu_average: <false|true>
```

### Network

```yaml
network:
  <include|exclude>:
    interfaces: [ <interface name>, ... ]
    match_type: <strict|regexp>
```

### Process

```yaml
process:
  <include|exclude>:
    names: [ <process name>, ... ]
    match_type: <strict|regexp>
  mute_process_name_error: <true|false>
  scrape_process_delay: <time>
```

## Advanced Configuration

### Filtering

If you are only interested in a subset of metrics from a particular source,
it is recommended you use this receiver with the
[Filter Processor](../../processor/filterprocessor).

### Different Frequencies

If you would like to scrape some metrics at a different frequency than others,
you can configure multiple `hostmetrics` receivers with different
`collection_interval` values. For example:

```yaml
receivers:
  hostmetrics:
    collection_interval: 30s
    scrapers:
      cpu:
      memory:

  hostmetrics/disk:
    collection_interval: 1m
    scrapers:
      disk:
      filesystem:

service:
  pipelines:
    metrics:
      receivers: [hostmetrics, hostmetrics/disk]
```

## Resource attributes

Currently, the hostmetrics receiver does not set any Resource attributes on the exported metrics. However, if you want to set Resource attributes, you can provide them via environment variables via the [resourcedetection](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/resourcedetectionprocessor#environment-variable) processor. For example, you can add the following resource attributes to adhere to [Resource Semantic Conventions](https://opentelemetry.io/docs/reference/specification/resource/semantic_conventions/):

```
export OTEL_RESOURCE_ATTRIBUTES="service.name=<the name of your service>,service.namespace=<the namespace of your service>,service.instance.id=<uuid of the instance>"
```

## Deprecations

### Transition to process memory metric names aligned with OpenTelemetry specification

The Host Metrics receiver has been emitting the following process memory metrics:

- [process.memory.physical_usage] for the amount of physical memory used by the process,
- [process.memory.virtual_usage] for the amount of virtual memory used by the process.

This is in conflict with the OpenTelemetry specification,
which defines [process.memory.usage] and [process.memory.virtual] as the names for these metrics.

To align the emitted metric names with the OpenTelemetry specification,
the following process will be followed to phase out the old metrics:

- Until and including `v0.63.0`, only the old metrics `process.memory.physical_usage` and `process.memory.virtual_usage` are emitted.
  You can use the [Metrics Transform processor][metricstransformprocessor_docs] to rename them.
- Between `v0.64.0` and `v0.66.0`, the new metrics are introduced as optional (disabled by default) and the old metrics are marked as deprecated.
  Only the old metrics are emitted by default.
- Between `v0.67.0` and `v0.69.0`, the new metrics are enabled and the old metrics are disabled by default.
- In `v0.70.0` and up, the old metrics are removed.

To change the enabled state for the specific metrics, use the standard configuration options that are available for all metrics.

Here's an example configuration to disable the old metrics and enable the new metrics:

```yaml
receivers:
  hostmetrics:
    scrapers:
      process:
        metrics:
          process.memory.physical_usage:
            enabled: false
          process.memory.virtual_usage:
            enabled: false
          process.memory.usage:
            enabled: true
          process.memory.virtual:
            enabled: true
```

[beta]: https://github.com/open-telemetry/opentelemetry-collector#beta
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
[core]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol
[process.memory.physical_usage]: https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/v0.63.0/receiver/hostmetricsreceiver/internal/scraper/processscraper/metadata.yaml#L61
[process.memory.virtual_usage]: https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/v0.63.0/receiver/hostmetricsreceiver/internal/scraper/processscraper/metadata.yaml#L70
[process.memory.usage]: https://github.com/open-telemetry/opentelemetry-specification/blob/v1.14.0/specification/metrics/semantic_conventions/process-metrics.md?plain=1#L38
[process.memory.virtual]: https://github.com/open-telemetry/opentelemetry-specification/blob/v1.14.0/specification/metrics/semantic_conventions/process-metrics.md?plain=1#L39
[metricstransformprocessor_docs]: https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/v0.63.0/processor/metricstransformprocessor/README.md