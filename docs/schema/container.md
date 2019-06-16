# Response Schema for Container informations

## Container

| Key         | Type    | Description                                                           |
| ----------- | ------- | --------------------------------------------------------------------- |
| containerId | string  | Container indentifier                                                 |
| name        | string  | Container string                                                      |
| image       | string  | Container image                                                       |
| status      | string  | Container status                                                      |
| running     | boolean | True when up, false otherwise                                         |
| cpu         | object  | Container cpu usage, see [container cpu stat](#container-cpu-stat)    |
| memory      | object  | Container memory usage, see [container mem stat](#container-mem-stat) |

### Container CPU Stat

| Key     | Type   | Description                             |
| ------- | ------ | --------------------------------------- |
| percent | string | Container CPU usage in %                |
| times   | object | A [cpu time](./cpu.md#cpu-times) object |

### Container Mem Stat

| Key                     | Type   | Description                                |
| ----------------------- | ------ | ------------------------------------------ |
| containerId             | string | Container identifier                       |
| cache                   | string | Cache used                                 |
| rss                     | string | Number of rss                              |
| rssHuge                 | string | Number of rss huge                         |
| mappedFile              | string | Number of mapped file                      |
| pgpgin                  | string | Number of page-ins                         |
| pgpgout                 | string | Number of page-outs                        |
| pgfault                 | string | Number of page faults                      |
| pgmajfault              | string | Number of major page faults                |
| inactiveAnon            | string | Number of inactive anonymous memory        |
| activeAnon              | string | Number of active anonymous memory          |
| inactiveFile            | string | Number of inactive cache memory file       |
| activeFile              | string | Number of active cache memory file         |
| unevictable             | string | Amount of memory that cannot be reclaimed  |
| hierarchicalMemoryLimit | string | Maximum RAM available in the host          |
| totalCache              | string | Total of cache                             |
| totalRss                | string | Total of rss                               |
| totalRssHuge            | string | Total of rss huge                          |
| totalMappedFile         | string | Total of mapped file                       |
| totalPgpgin             | string | Total of page-ins                          |
| totalPgpgout            | string | Total of page-outs                         |
| totalPgfault            | string | Total of page faults                       |
| totalPgmajfault         | string | Total of major page faults                 |
| totalInactiveAnon       | string | Total of inactive anonymous memory         |
| totalActiveAnon         | string | Total of active anonymous memory           |
| totalInactiveFile       | string | Total of inactive cache memory file        |
| totalActiveFile         | string | Total of active cache memory file          |
| totalUnevictable        | string | Total of unevictable memory                |
| memUsage                | string | Amount of memory used                      |
| memMaxUsage             | string | Amount of maximun memory usable            |
| memoryLimit             | string | Amount of physical memory that can be used |
| memoryFailcnt           | string | Amount of overload memory                  |

More info on [docker docs](https://docs.docker.com/v17.09/engine/admin/runmetrics/#metrics-from-cgroups-memory-cpu-block-io)
