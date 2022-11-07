[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# hostmetricsreceiver/process

## Metrics

These are the metrics available for this scraper.

| Name | Description | Unit | Type | Attributes |
| ---- | ----------- | ---- | ---- | ---------- |
| process.context_switches | Number of times the process has been context switched. | {count} | Sum(Int) | <ul> <li>context_switch_type</li> </ul> |
| **process.cpu.time** | Total CPU seconds broken down by different states. | s | Sum(Double) | <ul> <li>state</li> </ul> |
| **process.disk.io** | Disk bytes transferred. | By | Sum(Int) | <ul> <li>direction</li> </ul> |
| **process.memory.physical_usage** | Deprecated: use `process.memory.usage` metric instead. The amount of physical memory in use. | By | Sum(Int) | <ul> </ul> |
| process.memory.usage | The amount of physical memory in use. | By | Sum(Int) | <ul> </ul> |
| process.memory.virtual | Virtual memory size. | By | Sum(Int) | <ul> </ul> |
| **process.memory.virtual_usage** | Deprecated: Use `process.memory.virtual` metric instead. Virtual memory size. | By | Sum(Int) | <ul> </ul> |
| process.open_file_descriptors | Number of file descriptors in use by the process. | {count} | Sum(Int) | <ul> </ul> |
| process.paging.faults | Number of page faults the process has made. This metric is only available on Linux. | {faults} | Sum(Int) | <ul> <li>paging_fault_type</li> </ul> |
| process.threads | Process threads count. | {threads} | Sum(Int) | <ul> </ul> |

**Highlighted metrics** are emitted by default. Other metrics are optional and not emitted by default.
Any metric can be enabled or disabled with the following scraper configuration:

```yaml
metrics:
  <metric_name>:
    enabled: <true|false>
```

## Resource attributes

| Name | Description | Type |
| ---- | ----------- | ---- |
| process.command | The command used to launch the process (i.e. the command name). On Linux based systems, can be set to the zeroth string in proc/[pid]/cmdline. On Windows, can be set to the first parameter extracted from GetCommandLineW. | Str |
| process.command_line | The full command used to launch the process as a single string representing the full command. On Windows, can be set to the result of GetCommandLineW. Do not set this if you have to assemble it just for monitoring; use process.command_args instead. | Str |
| process.executable.name | The name of the process executable. On Linux based systems, can be set to the Name in proc/[pid]/status. On Windows, can be set to the base name of GetProcessImageFileNameW. | Str |
| process.executable.path | The full path to the process executable. On Linux based systems, can be set to the target of proc/[pid]/exe. On Windows, can be set to the result of GetProcessImageFileNameW. | Str |
| process.owner | The username of the user that owns the process. | Str |
| process.parent_pid | Parent Process identifier (PPID). | Int |
| process.pid | Process identifier (PID). | Int |

## Metric attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| context_switch_type (type) | Type of context switched. | involuntary, voluntary |
| direction | Direction of flow of bytes (read or write). | read, write |
| paging_fault_type (type) | Type of memory paging fault. | major, minor |
| state | Breakdown of CPU usage by type. | system, user, wait |